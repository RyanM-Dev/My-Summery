package grpc

import (
	"context"
	"log/slog"

	paymentv1 "github.com/huseyinbabal/microservices-proto/gen/go/payment/v1"
	"github.com/huseyinbabal/microservices/payment/internal/application/core/domain"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Adapter) Create(ctx context.Context, request *paymentv1.CreateRequest) (*paymentv1.CreateResponse, error) {
	slog.InfoContext(ctx, "Creating payment...", "order_id", request.GetOrderId())

	if err := validate(request); err != nil {
		return nil, err
	}

	newPayment := domain.NewPayment(request.GetUserId(), request.GetOrderId(), request.GetTotalPrice())
	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		slog.ErrorContext(ctx, "charge failed", "error", err)
		return nil, status.Errorf(codes.Internal, "failed to charge: %v", err)
	}
	return &paymentv1.CreateResponse{PaymentId: result.ID}, nil
}

// validate returns InvalidArgument with one field violation per bad field.
func validate(request *paymentv1.CreateRequest) error {
	var violations []*errdetails.BadRequest_FieldViolation
	if request.GetUserId() < 1 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field: "user_id", Description: "user id cannot be less than 1",
		})
	}
	if request.GetOrderId() < 1 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field: "order_id", Description: "order id cannot be less than 1",
		})
	}
	if request.GetTotalPrice() <= 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field: "total_price", Description: "total price must be greater than 0",
		})
	}
	if len(violations) == 0 {
		return nil
	}

	st, err := status.New(codes.InvalidArgument, "invalid payment request").
		WithDetails(&errdetails.BadRequest{FieldViolations: violations})
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid payment request")
	}
	return st.Err()
}
