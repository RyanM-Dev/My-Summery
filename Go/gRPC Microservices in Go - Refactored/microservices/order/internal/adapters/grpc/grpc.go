package grpc

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	orderv1 "github.com/huseyinbabal/microservices-proto/gen/go/order/v1"
	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Adapter) Create(ctx context.Context, request *orderv1.CreateRequest) (*orderv1.CreateResponse, error) {
	slog.InfoContext(ctx, "Creating order...", "user_id", request.GetUserId())

	if err := validateCreate(request); err != nil {
		return nil, err
	}

	orderItems := make([]domain.OrderItem, 0, len(request.GetOrderItems()))
	for _, item := range request.GetOrderItems() {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.GetProductCode(),
			UnitPrice:   item.GetUnitPrice(),
			Quantity:    item.GetQuantity(),
		})
	}

	newOrder := domain.NewOrder(request.GetUserId(), orderItems)
	result, err := a.api.PlaceOrder(ctx, newOrder)
	if err != nil {
		slog.ErrorContext(ctx, "order creation failed", "error", err)
		return nil, toStatus(err)
	}
	return &orderv1.CreateResponse{OrderId: result.ID}, nil
}

func (a *Adapter) Get(ctx context.Context, request *orderv1.GetRequest) (*orderv1.GetResponse, error) {
	if request.GetOrderId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "order id cannot be less than 1")
	}

	result, err := a.api.GetOrder(ctx, request.GetOrderId())
	if err != nil {
		return nil, toStatus(err)
	}

	orderItems := make([]*orderv1.OrderItem, 0, len(result.OrderItems))
	for _, item := range result.OrderItems {
		orderItems = append(orderItems, &orderv1.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	return &orderv1.GetResponse{
		UserId:     result.CustomerID,
		OrderItems: orderItems,
		Status:     result.Status,
	}, nil
}

// toStatus maps core errors to gRPC status errors. Errors that already carry
// a status (such as a wrapped Payment failure) pass through unchanged.
func toStatus(err error) error {
	if errors.Is(err, domain.ErrNotFound) {
		return status.Error(codes.NotFound, err.Error())
	}
	if _, ok := status.FromError(err); ok {
		return err
	}
	return status.Error(codes.Internal, "internal error")
}

// validateCreate returns InvalidArgument with one field violation per bad field.
func validateCreate(request *orderv1.CreateRequest) error {
	var violations []*errdetails.BadRequest_FieldViolation
	if request.GetUserId() < 1 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field: "user_id", Description: "user id cannot be less than 1",
		})
	}
	if len(request.GetOrderItems()) == 0 {
		violations = append(violations, &errdetails.BadRequest_FieldViolation{
			Field: "order_items", Description: "order must contain at least one item",
		})
	}
	for i, item := range request.GetOrderItems() {
		if item.GetProductCode() == "" {
			violations = append(violations, &errdetails.BadRequest_FieldViolation{
				Field: fmt.Sprintf("order_items[%d].product_code", i), Description: "product code cannot be empty",
			})
		}
		if item.GetUnitPrice() <= 0 {
			violations = append(violations, &errdetails.BadRequest_FieldViolation{
				Field: fmt.Sprintf("order_items[%d].unit_price", i), Description: "unit price must be greater than 0",
			})
		}
		if item.GetQuantity() < 1 {
			violations = append(violations, &errdetails.BadRequest_FieldViolation{
				Field: fmt.Sprintf("order_items[%d].quantity", i), Description: "quantity cannot be less than 1",
			})
		}
	}
	if len(violations) == 0 {
		return nil
	}

	st, err := status.New(codes.InvalidArgument, "invalid order request").
		WithDetails(&errdetails.BadRequest{FieldViolations: violations})
	if err != nil {
		return status.Error(codes.InvalidArgument, "invalid order request")
	}
	return st.Err()
}
