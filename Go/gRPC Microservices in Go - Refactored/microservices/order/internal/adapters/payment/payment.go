package payment

import (
	"context"
	"time"

	retry "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	paymentv1 "github.com/huseyinbabal/microservices-proto/gen/go/payment/v1"
	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// chargeTimeout is the whole budget for one Charge, retries included.
const chargeTimeout = 10 * time.Second

type Adapter struct {
	conn    *grpc.ClientConn
	payment paymentv1.PaymentServiceClient
}

// NewAdapter creates the Payment client. creds may be nil for plaintext.
// The connection is opened lazily on the first call and reused for every
// Charge; close it with Close on shutdown.
func NewAdapter(paymentServiceURL string, creds credentials.TransportCredentials) (*Adapter, error) {
	return newAdapter(paymentServiceURL, creds)
}

// newAdapter lets tests add dial options, such as an in-memory dialer.
func newAdapter(paymentServiceURL string, creds credentials.TransportCredentials, extra ...grpc.DialOption) (*Adapter, error) {
	if creds == nil {
		creds = insecure.NewCredentials()
	}

	opts := append([]grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		// With a headless Kubernetes Service, DNS returns every Payment pod;
		// round_robin spreads calls across them instead of pinning one pod.
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin": {}}]}`),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		// The first interceptor is the outermost: the breaker sees one
		// result per Charge, after the retries have run.
		grpc.WithChainUnaryInterceptor(
			circuitBreakerInterceptor(newCircuitBreaker()),
			retry.UnaryClientInterceptor(
				retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
				retry.WithMax(5),
				retry.WithBackoff(retry.BackoffLinearWithJitter(time.Second, 0.2)),
				retry.WithPerRetryTimeout(3*time.Second),
			),
		),
	}, extra...)

	conn, err := grpc.NewClient(paymentServiceURL, opts...)
	if err != nil {
		return nil, err
	}
	return &Adapter{conn: conn, payment: paymentv1.NewPaymentServiceClient(conn)}, nil
}

func (a *Adapter) Charge(ctx context.Context, order *domain.Order) error {
	ctx, cancel := context.WithTimeout(ctx, chargeTimeout)
	defer cancel()

	_, err := a.payment.Create(ctx, &paymentv1.CreateRequest{
		UserId:     order.CustomerID,
		OrderId:    order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}

func (a *Adapter) Close() error {
	return a.conn.Close()
}
