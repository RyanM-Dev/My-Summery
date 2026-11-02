package grpc

import (
	"fmt"
	"log/slog"
	"net"

	shippingv1 "github.com/huseyinbabal/microservices-proto/gen/go/shipping/v1"
	"github.com/huseyinbabal/microservices/shipping/internal/ports"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type Adapter struct {
	api    ports.APIPort
	port   int
	server *grpc.Server
	health *health.Server
	shippingv1.UnimplementedShippingServiceServer
}

// NewAdapter builds the gRPC server. With reflection on, grpcurl and buf curl
// can discover the API; turn it on in development only.
func NewAdapter(api ports.APIPort, port int, reflectionEnabled bool) *Adapter {
	a := &Adapter{api: api, port: port}

	a.server = grpc.NewServer(grpc.StatsHandler(otelgrpc.NewServerHandler()))
	shippingv1.RegisterShippingServiceServer(a.server, a)

	a.health = health.NewServer()
	healthpb.RegisterHealthServer(a.server, a.health)

	if reflectionEnabled {
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
	a.health.SetServingStatus(shippingv1.ShippingService_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	slog.Info("starting shipping service", "port", a.port)
	return a.server.Serve(listener)
}

// Stop reports NOT_SERVING to health checks, then finishes in-flight calls.
func (a *Adapter) Stop() {
	a.health.Shutdown()
	a.server.GracefulStop()
}
