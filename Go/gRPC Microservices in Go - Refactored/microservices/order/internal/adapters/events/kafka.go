// Package events is the Kafka side of the service: it publishes outbox events
// (driven adapter) and consumes events from other services (driver adapter).
package events

import (
	"time"

	"github.com/twmb/franz-go/pkg/kgo"
	"github.com/twmb/franz-go/plugin/kotel"
)

// NewProducer returns a client for publishing. franz-go defaults: acks=all
// and idempotent writes, so a retried send is never stored twice.
func NewProducer(brokers []string, clientID string, tracer *kotel.Tracer) (*kgo.Client, error) {
	return kgo.NewClient(
		kgo.SeedBrokers(brokers...),
		kgo.ClientID(clientID), // shows up in broker logs, metrics, and quotas
		// Give up on a record after 30 s (delivery.timeout.ms). Retries are
		// otherwise unlimited, and the outbox relay holds row locks while it waits.
		kgo.RecordDeliveryTimeout(30*time.Second),
		kgo.WithHooks(kotel.NewKotel(kotel.WithTracer(tracer)).Hooks()...),
	)
}

// NewTracer creates the tracer that continues traces through record headers.
// Create it after the global OpenTelemetry provider is set.
func NewTracer() *kotel.Tracer {
	return kotel.NewTracer()
}
