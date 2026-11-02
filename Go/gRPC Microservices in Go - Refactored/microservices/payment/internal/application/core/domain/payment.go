package domain

import (
	"errors"
	"time"
)

// ErrNotFound is returned when a payment does not exist.
var ErrNotFound = errors.New("payment not found")

type Payment struct {
	ID         int64   `json:"id"`
	CustomerID int64   `json:"customer_id"`
	Status     string  `json:"status"`
	OrderID    int64   `json:"order_id"`
	TotalPrice float32 `json:"total_price"`
	CreatedAt  int64   `json:"created_at"`
}

func NewPayment(customerID int64, orderID int64, totalPrice float32) Payment {
	return Payment{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CustomerID: customerID,
		OrderID:    orderID,
		TotalPrice: totalPrice,
	}
}
