package ports

import (
	"context"

	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
)

// PaymentPort is the driven port the order core uses to charge customers.
type PaymentPort interface {
	Charge(ctx context.Context, order *domain.Order) error
}
