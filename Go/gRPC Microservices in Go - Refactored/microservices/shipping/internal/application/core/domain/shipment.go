package domain

import (
	"errors"
	"time"
)

// ErrNotFound is returned when a shipment does not exist.
var ErrNotFound = errors.New("shipment not found")

// StatusCreated is the state of a new shipment.
const StatusCreated = "Created"

type Shipment struct {
	ID        int64     `json:"id"`
	OrderID   int64     `json:"order_id"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func NewShipment(orderID int64) Shipment {
	return Shipment{OrderID: orderID, Status: StatusCreated, CreatedAt: time.Now()}
}
