package api

import (
	"context"
	"testing"

	"github.com/huseyinbabal/microservices/shipping/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/shipping/internal/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Should_Create_Shipment(t *testing.T) {
	db := mocks.NewMockDBPort(t)
	db.EXPECT().CreateShipment(mock.Anything, "e1", mock.Anything).
		Run(func(_ context.Context, _ string, s *domain.Shipment) { s.ID = 5 }).
		Return(true, nil)

	shipment, err := NewApplication(db).CreateShipment(context.Background(), "e1", 42)

	assert.NoError(t, err)
	assert.Equal(t, int64(5), shipment.ID)
	assert.Equal(t, int64(42), shipment.OrderID)
	assert.Equal(t, domain.StatusCreated, shipment.Status)
}

func Test_Should_Return_Existing_Shipment_For_Duplicate_Event(t *testing.T) {
	db := mocks.NewMockDBPort(t)
	db.EXPECT().CreateShipment(mock.Anything, "e1", mock.Anything).Return(false, nil)
	db.EXPECT().GetByOrderID(mock.Anything, int64(42)).Return(domain.Shipment{ID: 5, OrderID: 42}, nil)

	shipment, err := NewApplication(db).CreateShipment(context.Background(), "e1", 42)

	assert.NoError(t, err)
	assert.Equal(t, int64(5), shipment.ID)
}
