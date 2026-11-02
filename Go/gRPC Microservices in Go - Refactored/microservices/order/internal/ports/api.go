package ports

import (
	"context"

	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
)

// APIPort is the driver port: what the outside world can ask the order core to do.
type APIPort interface {
	PlaceOrder(ctx context.Context, order domain.Order) (domain.Order, error)
	GetOrder(ctx context.Context, id int64) (domain.Order, error)
	// MarkShipped is called when Shipping reports a shipment (async).
	MarkShipped(ctx context.Context, id int64) error
}
