package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/huseyinbabal/microservices/payment/config"
	"github.com/huseyinbabal/microservices/payment/internal/adapters/db"
	"github.com/huseyinbabal/microservices/payment/internal/adapters/grpc"
	"github.com/huseyinbabal/microservices/payment/internal/application/core/api"
	"github.com/huseyinbabal/microservices/payment/internal/telemetry"
)

const service = "payment"

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

	opts := grpc.Options{Reflection: env == "development"}
	if certDir := config.GetTLSCertDir(); certDir != "" {
		if opts.Creds, err = grpc.ServerCredentials(certDir); err != nil {
			fatal("Failed to load TLS credentials", err)
		}
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, config.GetApplicationPort(), opts)

	go func() {
		<-ctx.Done()
		slog.Info("shutting down payment service")
		grpcAdapter.Stop()
	}()

	if err := grpcAdapter.Run(); err != nil {
		fatal("gRPC server stopped", err)
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
