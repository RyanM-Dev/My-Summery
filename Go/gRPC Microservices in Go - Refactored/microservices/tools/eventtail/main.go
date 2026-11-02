// eventtail prints the events in a Kafka topic as JSON. Protobuf payloads
// are binary, so console consumers show them as noise; this tool decodes
// them using the event-type header.
//
//	go run ./tools/eventtail -topic orders
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	// Register the event types so they can be found by name.
	_ "github.com/huseyinbabal/microservices-proto/gen/go/order/v1"
	_ "github.com/huseyinbabal/microservices-proto/gen/go/shipping/v1"
)

func main() {
	brokers := flag.String("brokers", "localhost:9092", "comma-separated bootstrap servers")
	topic := flag.String("topic", "orders", "topic to read")
	follow := flag.Bool("follow", false, "keep waiting for new events")
	flag.Parse()

	client, err := kgo.NewClient(
		kgo.SeedBrokers(strings.Split(*brokers, ",")...),
		kgo.ConsumeTopics(*topic),
		kgo.ConsumeResetOffset(kgo.NewOffset().AtStart()), // no group: read everything, commit nothing
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if !*follow {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, 3*time.Second)
		defer cancel()
	}

	for {
		fetches := client.PollFetches(ctx)
		if ctx.Err() != nil {
			return
		}
		fetches.EachRecord(func(r *kgo.Record) {
			fmt.Printf("%s p%d@%d key=%s %s\n", r.Topic, r.Partition, r.Offset, r.Key, decode(r))
		})
	}
}

func decode(r *kgo.Record) string {
	eventType, dlqError := "", ""
	for _, h := range r.Headers {
		switch h.Key {
		case "event-type":
			eventType = string(h.Value)
		case "dlq-error":
			dlqError = string(h.Value)
		}
	}
	if dlqError != "" {
		return fmt.Sprintf("DEAD LETTER (%s) value=%q", dlqError, r.Value)
	}
	mt, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(eventType))
	if err != nil {
		return fmt.Sprintf("type=%q (unknown) %d bytes", eventType, len(r.Value))
	}
	msg := mt.New().Interface()
	if err := proto.Unmarshal(r.Value, msg); err != nil {
		return fmt.Sprintf("type=%s undecodable: %v", eventType, err)
	}
	return eventType + " " + protojson.MarshalOptions{}.Format(msg)
}
