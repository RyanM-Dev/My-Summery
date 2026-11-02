package api

import (
	"context"
	"errors"

	"github.com/huseyinbabal/microservices/payment/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/payment/internal/ports"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

// Charge records the payment. A real implementation would call a payment
// provider here before saving.
//
// Charge is idempotent per order: the order ID works as an idempotency key,
// so a retried request returns the first payment instead of charging twice.
func (a *Application) Charge(ctx context.Context, payment domain.Payment) (domain.Payment, error) {
	existing, err := a.db.GetByOrderID(ctx, payment.OrderID)
	if err == nil {
		return existing, nil
	}
	if !errors.Is(err, domain.ErrNotFound) {
		return domain.Payment{}, err
	}

	if err := a.db.Save(ctx, &payment); err != nil {
		return domain.Payment{}, err
	}
	return payment, nil
}
