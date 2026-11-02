package events

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/plugin/kotel"
)

// ErrPermanent marks failures that retrying can't fix, such as an event that
// can't be decoded. Wrap it to send a record straight to the dead-letter topic.
var ErrPermanent = errors.New("permanent failure")

// Handler processes one record. It must be idempotent: delivery is
// at-least-once, so the same record can arrive more than once.
type Handler func(ctx context.Context, record *kgo.Record) error

// ConsumerConfig names what to read and where failures go.
type ConsumerConfig struct {
	Brokers  []string
	Group    string // one group per service, e.g. "order"
	Topic    string // e.g. "shipments"
	DLQTopic string // e.g. "order.dlq"
}

// Consumer reads one topic as part of a consumer group and commits offsets
// only after every polled record was handled or dead-lettered.
type Consumer struct {
	client   *kgo.Client
	producer *kgo.Client // publishes dead letters
	tracer   *kotel.Tracer
	handle   Handler
	dlqTopic string
	attempts int
}

func NewConsumer(cfg ConsumerConfig, handle Handler, producer *kgo.Client, tracer *kotel.Tracer) (*Consumer, error) {
	client, err := kgo.NewClient(
		kgo.SeedBrokers(cfg.Brokers...),
		kgo.ClientID(cfg.Group+"-consumer"),
		kgo.ConsumerGroup(cfg.Group),
		kgo.ConsumeTopics(cfg.Topic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()), // new group: start from the oldest record
		kgo.DisableAutoCommit(),                            // commit only after processing
		kgo.BlockRebalanceOnPoll(),                         // no rebalance between poll and commit
		kgo.WithHooks(kotel.NewKotel(kotel.WithTracer(tracer)).Hooks()...),
	)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		client:   client,
		producer: producer,
		tracer:   tracer,
		handle:   handle,
		dlqTopic: cfg.DLQTopic,
		attempts: 3,
	}, nil
}

// Run polls, processes, and commits until ctx is cancelled.
func (c *Consumer) Run(ctx context.Context) {
	for {
		fetches := c.client.PollFetches(ctx)
		if fetches.IsClientClosed() || ctx.Err() != nil {
			c.client.AllowRebalance() // otherwise Close blocks while leaving the group
			return
		}
		fetches.EachError(func(topic string, partition int32, err error) {
			slog.Error("fetch failed", "topic", topic, "partition", partition, "error", err)
		})

		fetches.EachRecord(func(record *kgo.Record) {
			if err := c.processWithRetry(ctx, record); err != nil && ctx.Err() == nil {
				c.deadLetter(ctx, record, err)
			}
		})

		if ctx.Err() == nil {
			if err := c.client.CommitUncommittedOffsets(ctx); err != nil {
				slog.Error("commit failed", "error", err)
			}
		}
		c.client.AllowRebalance()
	}
}

// processWithRetry retries transient failures with a growing pause.
func (c *Consumer) processWithRetry(ctx context.Context, record *kgo.Record) error {
	var err error
	for attempt := 1; attempt <= c.attempts; attempt++ {
		err = c.process(record)
		if err == nil || errors.Is(err, ErrPermanent) {
			return err
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Duration(attempt) * 500 * time.Millisecond):
		}
	}
	return err
}

// process runs the handler inside a span that continues the producer's trace.
func (c *Consumer) process(record *kgo.Record) error {
	ctx, span := c.tracer.WithProcessSpan(record)
	defer span.End()
	if err := c.handle(ctx, record); err != nil {
		span.RecordError(err)
		slog.WarnContext(ctx, "event handling failed", "topic", record.Topic, "offset", record.Offset, "error", err)
		return err
	}
	return nil
}

// deadLetter parks a failed record with the reason, so one bad record can't
// block its partition. Replay it after fixing the cause.
func (c *Consumer) deadLetter(ctx context.Context, record *kgo.Record, cause error) {
	dlq := &kgo.Record{
		Topic: c.dlqTopic,
		Key:   record.Key,
		Value: record.Value,
		Headers: append(append([]kgo.RecordHeader{}, record.Headers...),
			kgo.RecordHeader{Key: "dlq-error", Value: []byte(cause.Error())},
			kgo.RecordHeader{Key: "dlq-source", Value: fmt.Appendf(nil, "%s/%d/%d", record.Topic, record.Partition, record.Offset)},
		),
	}
	if err := c.producer.ProduceSync(ctx, dlq).FirstErr(); err != nil {
		slog.Error("dead-letter publish failed", "error", err)
	}
}

// Close leaves the group, so its partitions move to other instances at once.
func (c *Consumer) Close() { c.client.Close() }
