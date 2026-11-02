package api

import (
	"context"
	"log/slog"

	"github.com/huseyinbabal/microservices/shipping/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/shipping/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

// CreateShipment ships a paid order. Repeated deliveries of the same event
// return the existing shipment instead of shipping twice.
func (a *Application) CreateShipment(ctx context.Context, eventID string, orderID int64) (domain.Shipment, error) {
	shipment := domain.NewShipment(orderID)
	created, err := a.db.CreateShipment(ctx, eventID, &shipment)
	if err != nil {
		return domain.Shipment{}, err
	}
	if !created {
		slog.InfoContext(ctx, "duplicate event skipped", "event_id", eventID, "order_id", orderID)
		return a.db.GetByOrderID(ctx, orderID)
	}
	slog.InfoContext(ctx, "shipment created", "shipment_id", shipment.ID, "order_id", orderID)
	return shipment, nil
}

func (a *Application) GetShipment(ctx context.Context, orderID int64) (domain.Shipment, error) {
	return a.db.GetByOrderID(ctx, orderID)
}
