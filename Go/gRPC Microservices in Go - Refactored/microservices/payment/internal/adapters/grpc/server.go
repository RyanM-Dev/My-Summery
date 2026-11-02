package grpc

import (
	"fmt"
	"log/slog"
	"net"

	paymentv1 "github.com/huseyinbabal/microservices-proto/gen/go/payment/v1"
	"github.com/huseyinbabal/microservices/payment/internal/ports"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	tls    bool
	server *grpc.Server
	health *health.Server
	paymentv1.UnimplementedPaymentServiceServer
}

// Options holds the optional server settings that come from config.
type Options struct {
	// Reflection lets grpcurl and buf curl discover the API. Development only.
	Reflection bool
	// Creds enables TLS (mTLS when it requires client certificates).
	// Nil means plaintext.
	Creds credentials.TransportCredentials
}

func NewAdapter(api ports.APIPort, port int, opts Options) *Adapter {
	a := &Adapter{api: api, port: port, tls: opts.Creds != nil}

	serverOpts := []grpc.ServerOption{
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	}
	if opts.Creds != nil {
		serverOpts = append(serverOpts, grpc.Creds(opts.Creds))
	}
	a.server = grpc.NewServer(serverOpts...)
	paymentv1.RegisterPaymentServiceServer(a.server, a)

	a.health = health.NewServer()
	healthpb.RegisterHealthServer(a.server, a.health)

	if opts.Reflection {
		reflection.Register(a.server)
	}
	return a
}

// Run blocks until Stop is called or the listener fails.
func (a *Adapter) Run() error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("failed to listen on port %d: %w", a.port, err)
	}
	a.health.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
	a.health.SetServingStatus(paymentv1.PaymentService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	slog.Info("starting payment service", "port", a.port, "tls", a.tls)
	return a.server.Serve(listener)
}

// Stop reports NOT_SERVING to health checks, then finishes in-flight calls.
func (a *Adapter) Stop() {
	a.health.Shutdown()
	a.server.GracefulStop()
}
