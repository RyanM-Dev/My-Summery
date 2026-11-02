package ports

import (
	"context"

	"github.com/huseyinbabal/microservices/payment/internal/application/core/domain"
)

// DBPort is the driven port the payment core uses to store payments.
type DBPort interface {
	Get(ctx context.Context, id int64) (domain.Payment, error)
	GetByOrderID(ctx context.Context, orderID int64) (domain.Payment, error)
	Save(ctx context.Context, payment *domain.Payment) error
}
