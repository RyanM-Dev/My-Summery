package api

import (
	"context"
	"errors"
	"testing"

	"github.com/huseyinbabal/microservices/payment/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/payment/internal/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Should_Charge(t *testing.T) {
	db := mocks.NewMockDBPort(t)
	db.EXPECT().GetByOrderID(mock.Anything, int64(1)).Return(domain.Payment{}, domain.ErrNotFound)
	db.EXPECT().Save(mock.Anything, mock.Anything).
		Run(func(_ context.Context, p *domain.Payment) { p.ID = 7 }).
		Return(nil)

	application := NewApplication(db)
	result, err := application.Charge(context.Background(), domain.NewPayment(123, 1, 24.6))

	assert.NoError(t, err)
	assert.Equal(t, int64(7), result.ID)
}

func Test_Should_Return_Error_When_Db_Persistence_Fail(t *testing.T) {
	db := mocks.NewMockDBPort(t)
	db.EXPECT().GetByOrderID(mock.Anything, int64(1)).Return(domain.Payment{}, domain.ErrNotFound)
	db.EXPECT().Save(mock.Anything, mock.Anything).Return(errors.New("connection error"))

	application := NewApplication(db)
	_, err := application.Charge(context.Background(), domain.NewPayment(123, 1, 24.6))

	assert.EqualError(t, err, "connection error")
}

func Test_Should_Return_Existing_Payment_For_Same_Order(t *testing.T) {
	db := mocks.NewMockDBPort(t)
	db.EXPECT().GetByOrderID(mock.Anything, int64(1)).Return(domain.Payment{ID: 7, OrderID: 1}, nil)
	// Save is never expected: a repeated request must not charge twice.

	application := NewApplication(db)
	result, err := application.Charge(context.Background(), domain.NewPayment(123, 1, 24.6))

	assert.NoError(t, err)
	assert.Equal(t, int64(7), result.ID)
}
