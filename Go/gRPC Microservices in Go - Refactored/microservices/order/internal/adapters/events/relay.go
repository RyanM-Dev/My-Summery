package events

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/huseyinbabal/microservices/order/internal/adapters/db"
	"github.com/twmb/franz-go/pkg/kgo"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

// Outbox is the part of the DB adapter the relay needs.
type Outbox interface {
	PublishOutbox(ctx context.Context, limit int, publish func(context.Context, []db.OutboxEvent) error) (int, error)
}

// Relay copies outbox rows to Kafka, oldest first. Delivery is at-least-once:
// a crash after publishing but before marking rows published republishes
// them, so consumers drop duplicates by event_id.
type Relay struct {
	outbox   Outbox
	producer *kgo.Client
	interval time.Duration
}

func NewRelay(outbox Outbox, producer *kgo.Client) *Relay {
	return &Relay{outbox: outbox, producer: producer, interval: 500 * time.Millisecond}
}

// Run publishes until ctx is cancelled.
func (r *Relay) Run(ctx context.Context) {
	ticker := time.NewTicker(r.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			// Drain the backlog in batches of 100 before waiting again.
			for {
				n, err := r.outbox.PublishOutbox(ctx, 100, r.publish)
				if err != nil {
					if ctx.Err() == nil {
						slog.Error("outbox relay failed", "error", err)
					}
					break
				}
				if n < 100 {
					break
				}
			}
		}
	}
}

func (r *Relay) publish(ctx context.Context, events []db.OutboxEvent) error {
	records := make([]*kgo.Record, 0, len(events))
	for _, e := range events {
		records = append(records, &kgo.Record{
			Context: parentContext(ctx, e.Headers),
			Topic:   e.Topic,
			Key:     []byte(e.Key),
			Value:   e.Payload,
			Headers: []kgo.RecordHeader{{Key: "event-type", Value: []byte(e.EventType)}},
		})
	}
	return r.producer.ProduceSync(ctx, records...).FirstErr()
}

// parentContext restores the trace of the request that created the event, so
// the publish span joins that trace instead of starting a new one.
func parentContext(ctx context.Context, headers []byte) context.Context {
	carrier := propagation.MapCarrier{}
	if err := json.Unmarshal(headers, &carrier); err != nil {
		return ctx
	}
	return otel.GetTextMapPropagator().Extract(ctx, carrier)
}
