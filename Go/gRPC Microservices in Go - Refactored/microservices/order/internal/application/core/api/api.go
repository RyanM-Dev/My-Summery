package api

import (
	"context"
	"log/slog"

	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/order/internal/ports"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort) *Application {
	return &Application{db: db, payment: payment}
}

// PlaceOrder saves the order as Pending, charges the customer (sync gRPC),
// and marks the order Paid, or Failed when the charge fails. Marking it Paid
// also queues an OrderPaid event that reaches Shipping through Kafka (async).
func (a *Application) PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error) {
	if err := a.db.Save(ctx, &order); err != nil {
		return domain.Order{}, err
	}

	if paymentErr := a.payment.Charge(ctx, &order); paymentErr != nil {
		if err := a.db.UpdateStatus(ctx, order.ID, domain.StatusFailed); err != nil {
			slog.ErrorContext(ctx, "failed to mark order as failed", "order_id", order.ID, "error", err)
		}
		return domain.Order{}, paymentError(paymentErr)
	}

	order.Status = domain.StatusPaid
	if err := a.db.MarkPaid(ctx, &order); err != nil {
		return domain.Order{}, err
	}
	return order, nil
}

func (a *Application) GetOrder(ctx context.Context, id int64) (domain.Order, error) {
	return a.db.Get(ctx, id)
}

// MarkShipped is idempotent: a repeated ShipmentCreated event sets the same status.
func (a *Application) MarkShipped(ctx context.Context, id int64) error {
	return a.db.UpdateStatus(ctx, id, domain.StatusShipped)
}

// paymentError wraps a Payment failure as "order creation failed", keeping
// Payment's status code so callers can still tell bad input (InvalidArgument)
// from an outage (Unavailable). The Payment message goes into a BadRequest
// field violation on the "payment" field.
func paymentError(paymentErr error) error {
	st := status.Convert(paymentErr)

	description := st.Message()
	for _, detail := range st.Details() {
		if badRequest, ok := detail.(*errdetails.BadRequest); ok {
			for _, violation := range badRequest.GetFieldViolations() {
				description += "; " + violation.GetField() + ": " + violation.GetDescription()
			}
		}
	}

	code := st.Code()
	if code == codes.Unknown {
		code = codes.Internal
	}
	orderStatus, err := status.New(code, "order creation failed").WithDetails(&errdetails.BadRequest{
		FieldViolations: []*errdetails.BadRequest_FieldViolation{
			{Field: "payment", Description: description},
		},
	})
	if err != nil {
		return status.Error(code, "order creation failed")
	}
	return orderStatus.Err()
}
