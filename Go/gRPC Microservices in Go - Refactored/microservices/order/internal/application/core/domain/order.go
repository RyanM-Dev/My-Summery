package domain

import (
	"errors"
	"time"
)

// ErrNotFound is returned when an order does not exist.
var ErrNotFound = errors.New("order not found")

// Order statuses.
const (
	StatusPending = "Pending"
	StatusPaid    = "Paid"
	StatusShipped = "Shipped"
	StatusFailed  = "Failed"
)

type Order struct {
	ID         int64       `json:"id"`
	CustomerID int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

func NewOrder(customerID int64, orderItems []OrderItem) Order {
	return Order{
		CreatedAt:  time.Now().Unix(),
		Status:     StatusPending,
		CustomerID: customerID,
		OrderItems: orderItems,
	}
}

// TotalPrice is the sum of unit price × quantity over all items.
func (o *Order) TotalPrice() float32 {
	var totalPrice float32
	for _, orderItem := range o.OrderItems {
		totalPrice += orderItem.UnitPrice * float32(orderItem.Quantity)
	}
	return totalPrice
}
