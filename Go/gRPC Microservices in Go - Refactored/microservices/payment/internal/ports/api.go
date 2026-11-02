package ports

import (
	"context"

	"github.com/huseyinbabal/microservices/payment/internal/application/core/domain"
)

// APIPort is the driver port: what the outside world can ask the payment core to do.
type APIPort interface {
	Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error)
}
