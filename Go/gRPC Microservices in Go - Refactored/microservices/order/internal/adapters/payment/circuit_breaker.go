package payment

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/sony/gobreaker/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// newCircuitBreaker opens after at least 5 calls with a 60% failure rate,
// lets 3 test calls through when half-open, and waits 4 s before half-open.
func newCircuitBreaker() *gobreaker.CircuitBreaker[any] {
	return gobreaker.NewCircuitBreaker[any](gobreaker.Settings{
		Name:        "payment",
		MaxRequests: 3,
		Interval:    time.Minute,
		Timeout:     4 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 5 && failureRatio >= 0.6
		},
		// Client mistakes say nothing about Payment's health, so they
		// count as successes.
		IsSuccessful: func(err error) bool {
			switch status.Code(err) {
			case codes.OK, codes.InvalidArgument, codes.NotFound, codes.AlreadyExists,
				codes.PermissionDenied, codes.Unauthenticated, codes.FailedPrecondition:
				return true
			default:
				return false
			}
		},
		OnStateChange: func(name string, from, to gobreaker.State) {
			slog.Warn("circuit breaker state changed", "name", name, "from", from.String(), "to", to.String())
		},
	})
}

// circuitBreakerInterceptor runs every unary call through cb. While the
// circuit is open, calls fail at once with Unavailable instead of reaching
// the network.
func circuitBreakerInterceptor(cb *gobreaker.CircuitBreaker[any]) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any,
		cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {

		_, err := cb.Execute(func() (any, error) {
			return nil, invoker(ctx, method, req, reply, cc, opts...)
		})
		if errors.Is(err, gobreaker.ErrOpenState) || errors.Is(err, gobreaker.ErrTooManyRequests) {
			return status.Errorf(codes.Unavailable, "payment: %v", err)
		}
		return err
	}
}
