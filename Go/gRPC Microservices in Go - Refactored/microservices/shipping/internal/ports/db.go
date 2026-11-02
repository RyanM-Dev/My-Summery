package ports

import (
	"context"

	"github.com/huseyinbabal/microservices/shipping/internal/application/core/domain"
)

// DBPort is the driven port the shipping core uses to store shipments.
type DBPort interface {
	// CreateShipment saves the shipment and queues ShipmentCreated in one
	// transaction, unless eventID was already processed. It reports whether
	// a shipment was created.
	CreateShipment(ctx context.Context, eventID string, shipment *domain.Shipment) (bool, error)
	GetByOrderID(ctx context.Context, orderID int64) (domain.Shipment, error)
}
