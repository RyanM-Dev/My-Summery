package events

import (
	"context"
	"fmt"

	orderv1 "github.com/huseyinbabal/microservices-proto/gen/go/order/v1"
	"github.com/huseyinbabal/microservices/shipping/internal/ports"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
)

// TopicOrders carries order.v1 events, keyed by order ID.
const TopicOrders = "orders"

// OrderPaidHandler creates a shipment for every paid order.
func OrderPaidHandler(api ports.APIPort) Handler {
	return func(ctx context.Context, record *kgo.Record) error {
		var event orderv1.OrderPaid
		if err := proto.Unmarshal(record.Value, &event); err != nil {
			return fmt.Errorf("%w: decode OrderPaid: %v", ErrPermanent, err)
		}
		if event.GetEventId() == "" || event.GetOrderId() < 1 {
			return fmt.Errorf("%w: invalid OrderPaid: event_id=%q order_id=%d",
				ErrPermanent, event.GetEventId(), event.GetOrderId())
		}
		_, err := api.CreateShipment(ctx, event.GetEventId(), event.GetOrderId())
		return err
	}
}
