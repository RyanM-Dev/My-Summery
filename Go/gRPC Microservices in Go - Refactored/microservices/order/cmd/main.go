package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/huseyinbabal/microservices/order/config"
	"github.com/huseyinbabal/microservices/order/internal/adapters/db"
	"github.com/huseyinbabal/microservices/order/internal/adapters/events"
	"github.com/huseyinbabal/microservices/order/internal/adapters/grpc"
	"github.com/huseyinbabal/microservices/order/internal/adapters/payment"
	"github.com/huseyinbabal/microservices/order/internal/application/core/api"
	"github.com/huseyinbabal/microservices/order/internal/telemetry"
	"google.golang.org/grpc/credentials"
)

const service = "order"

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

	// Sync: Order → Payment over gRPC.
	var paymentCreds credentials.TransportCredentials
	if certDir := config.GetTLSCertDir(); certDir != "" {
		if paymentCreds, err = payment.ClientCredentials(certDir); err != nil {
			fatal("Failed to load TLS credentials", err)
		}
	}
	paymentAdapter, err := payment.NewAdapter(config.GetPaymentServiceURL(), paymentCreds)
	if err != nil {
		fatal("Failed to initialize payment stub", err)
	}

	application := api.NewApplication(dbAdapter, paymentAdapter)

	// Async: outbox → Kafka "orders", and Kafka "shipments" → Order.
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
		Topic:    events.TopicShipments,
		DLQTopic: "order.dlq",
	}, events.ShipmentCreatedHandler(application), producer, tracer)
	if err != nil {
		fatal("Failed to create Kafka consumer", err)
	}

	var workers sync.WaitGroup
	workers.Go(func() { relay.Run(ctx) })
	workers.Go(func() { consumer.Run(ctx) })

	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort(), env == "development")
	go func() {
		<-ctx.Done()
		slog.Info("shutting down order service")
		grpcAdapter.Stop()
	}()

	if err := grpcAdapter.Run(); err != nil {
		fatal("gRPC server stopped", err)
	}

	workers.Wait()
	consumer.Close()
	producer.Close()
	if err := paymentAdapter.Close(); err != nil {
		slog.Error("failed to close payment connection", "error", err)
	}
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
