package grpc

import (
	"context"
	"errors"

	shippingv1 "github.com/huseyinbabal/microservices-proto/gen/go/shipping/v1"
	"github.com/huseyinbabal/microservices/shipping/internal/application/core/domain"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Adapter) Get(ctx context.Context, request *shippingv1.GetRequest) (*shippingv1.GetResponse, error) {
	if request.GetOrderId() < 1 {
		return nil, status.Error(codes.InvalidArgument, "order id cannot be less than 1")
	}
	shipment, err := a.api.GetShipment(ctx, request.GetOrderId())
	if errors.Is(err, domain.ErrNotFound) {
		return nil, status.Error(codes.NotFound, err.Error())
	}
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}
	return &shippingv1.GetResponse{
		ShipmentId: shipment.ID,
		OrderId:    shipment.OrderID,
		Status:     shipment.Status,
	}, nil
}
