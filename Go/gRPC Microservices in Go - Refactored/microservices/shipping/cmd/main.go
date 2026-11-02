package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/huseyinbabal/microservices/shipping/config"
	"github.com/huseyinbabal/microservices/shipping/internal/adapters/db"
	"github.com/huseyinbabal/microservices/shipping/internal/adapters/events"
	"github.com/huseyinbabal/microservices/shipping/internal/adapters/grpc"
	"github.com/huseyinbabal/microservices/shipping/internal/application/core/api"
	"github.com/huseyinbabal/microservices/shipping/internal/telemetry"
)

const service = "shipping"

func main() {
	telemetry.SetupLogging(os.Stdout)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	env := config.GetEnv()
	shutdownTracing, err := telemetry.SetupTracing(ctx, service, env)
	if err != nil {
		fatal("Failed to set up tracing", err)
	}

	dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	if err != nil {
		fatal("Failed to connect to database", err)
	}
	application := api.NewApplication(dbAdapter)

	// Async: Kafka "orders" → Shipping, and outbox → Kafka "shipments".
	brokers := config.GetKafkaBrokers()
	tracer := events.NewTracer()
	producer, err := events.NewProducer(brokers, service+"-producer", tracer)
	if err != nil {
		fatal("Failed to create Kafka producer", err)
	}
	relay := events.NewRelay(dbAdapter, producer)
	consumer, err := events.NewConsumer(events.ConsumerConfig{
		Brokers:  brokers,
		Group:    service,
		Topic:    events.TopicOrders,
		DLQTopic: "shipping.dlq",
	}, events.OrderPaidHandler(application), producer, tracer)
	if err != nil {
		fatal("Failed to create Kafka consumer", err)
	}

	var workers sync.WaitGroup
	workers.Go(func() { relay.Run(ctx) })
	workers.Go(func() { consumer.Run(ctx) })

	// Sync: ShippingService.Get over gRPC.
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort(), env == "development")
	go func() {
		<-ctx.Done()
		slog.Info("shutting down shipping service")
		grpcAdapter.Stop()
	}()

	if err := grpcAdapter.Run(); err != nil {
		fatal("gRPC server stopped", err)
	}

	workers.Wait()
	consumer.Close()
	producer.Close()
	flushCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := shutdownTracing(flushCtx); err != nil {
		slog.Error("failed to flush spans", "error", err)
	}
}

func fatal(msg string, err error) {
	slog.Error(msg, "error", err)
	os.Exit(1)
}
