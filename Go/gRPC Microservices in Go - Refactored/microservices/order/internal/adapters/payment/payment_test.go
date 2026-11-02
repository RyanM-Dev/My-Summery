package payment

import (
	"context"
	"net"
	"sync/atomic"
	"testing"

	paymentv1 "github.com/huseyinbabal/microservices-proto/gen/go/payment/v1"
	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/test/bufconn"
)

// fakePayment fails the first `failures` calls with `code`, then succeeds.
type fakePayment struct {
	paymentv1.UnimplementedPaymentServiceServer
	calls    atomic.Int32
	failures int32
	code     codes.Code
}

func (f *fakePayment) Create(context.Context, *paymentv1.CreateRequest) (*paymentv1.CreateResponse, error) {
	if f.calls.Add(1) <= f.failures {
		return nil, status.Error(f.code, "simulated failure")
	}
	return &paymentv1.CreateResponse{PaymentId: 1}, nil
}

// startFake serves fake in memory and returns an adapter connected to it.
func startFake(t *testing.T, fake *fakePayment) *Adapter {
	t.Helper()
	listener := bufconn.Listen(1 << 20)
	server := grpc.NewServer()
	paymentv1.RegisterPaymentServiceServer(server, fake)
	go func() { _ = server.Serve(listener) }()
	t.Cleanup(server.Stop)

	adapter, err := newAdapter("passthrough:///bufnet", nil,
		grpc.WithContextDialer(func(ctx context.Context, _ string) (net.Conn, error) {
			return listener.DialContext(ctx)
		}),
	)
	require.NoError(t, err)
	t.Cleanup(func() { _ = adapter.Close() })
	return adapter
}

func testOrder() *domain.Order {
	return &domain.Order{ID: 1, CustomerID: 23, OrderItems: []domain.OrderItem{
		{ProductCode: "CAM123", UnitPrice: 1.23, Quantity: 3},
	}}
}

func TestCharge_RetriesUnavailable(t *testing.T) {
	fake := &fakePayment{failures: 2, code: codes.Unavailable}
	adapter := startFake(t, fake)

	err := adapter.Charge(context.Background(), testOrder())

	assert.NoError(t, err)
	assert.Equal(t, int32(3), fake.calls.Load(), "2 failures + 1 success")
}

func TestCharge_DoesNotRetryInvalidArgument(t *testing.T) {
	fake := &fakePayment{failures: 100, code: codes.InvalidArgument}
	adapter := startFake(t, fake)

	err := adapter.Charge(context.Background(), testOrder())

	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	assert.Equal(t, int32(1), fake.calls.Load())
}

func TestCharge_CircuitOpensAfterRepeatedFailures(t *testing.T) {
	fake := &fakePayment{failures: 100, code: codes.Internal}
	adapter := startFake(t, fake)

	for range 5 {
		assert.Equal(t, codes.Internal, status.Code(adapter.Charge(context.Background(), testOrder())))
	}
	err := adapter.Charge(context.Background(), testOrder())

	assert.Equal(t, codes.Unavailable, status.Code(err))
	assert.Contains(t, err.Error(), "circuit breaker is open")
	assert.Equal(t, int32(5), fake.calls.Load(), "the 6th call never reached the server")
}
