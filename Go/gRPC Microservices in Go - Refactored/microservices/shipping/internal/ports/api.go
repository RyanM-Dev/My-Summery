package ports

import (
	"context"

	"github.com/huseyinbabal/microservices/shipping/internal/application/core/domain"
)

// APIPort is the driver port. CreateShipment is driven by OrderPaid events
// (async); GetShipment by the gRPC API (sync).
type APIPort interface {
	CreateShipment(ctx context.Context, eventID string, orderID int64) (domain.Shipment, error)
	GetShipment(ctx context.Context, orderID int64) (domain.Shipment, error)
}
