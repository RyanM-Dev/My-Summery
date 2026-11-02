package ports

import (
	"context"

	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
)

// DBPort is the driven port the order core uses to store orders.
type DBPort interface {
	Get(ctx context.Context, id int64) (domain.Order, error)
	Save(ctx context.Context, order *domain.Order) error
	UpdateStatus(ctx context.Context, id int64, status string) error
	// MarkPaid sets the status to Paid and records an OrderPaid event in the
	// outbox, in one transaction.
	MarkPaid(ctx context.Context, order *domain.Order) error
}
