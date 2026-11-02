package events

import (
	"context"
	"errors"
	"fmt"

	shippingv1 "github.com/huseyinbabal/microservices-proto/gen/go/shipping/v1"
	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/order/internal/ports"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
)

// TopicShipments carries shipping.v1 events, keyed by order ID.
const TopicShipments = "shipments"

// ShipmentCreatedHandler marks the order Shipped when Shipping reports it.
// The status update is idempotent, so duplicates are harmless.
func ShipmentCreatedHandler(api ports.APIPort) Handler {
	return func(ctx context.Context, record *kgo.Record) error {
		var event shippingv1.ShipmentCreated
		if err := proto.Unmarshal(record.Value, &event); err != nil {
			return fmt.Errorf("%w: decode ShipmentCreated: %v", ErrPermanent, err)
		}
		err := api.MarkShipped(ctx, event.GetOrderId())
		if errors.Is(err, domain.ErrNotFound) {
			return fmt.Errorf("%w: order %d: %v", ErrPermanent, event.GetOrderId(), err)
		}
		return err
	}
}
