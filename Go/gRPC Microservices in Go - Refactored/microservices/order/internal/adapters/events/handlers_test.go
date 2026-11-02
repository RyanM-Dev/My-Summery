package events

import (
	"context"
	"testing"

	shippingv1 "github.com/huseyinbabal/microservices-proto/gen/go/shipping/v1"
	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/order/internal/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/twmb/franz-go/pkg/kgo"
	"google.golang.org/protobuf/proto"
)

func record(t *testing.T, event proto.Message) *kgo.Record {
	t.Helper()
	value, err := proto.Marshal(event)
	require.NoError(t, err)
	return &kgo.Record{Value: value}
}

func TestShipmentCreatedHandler_MarksShipped(t *testing.T) {
	api := mocks.NewMockAPIPort(t)
	api.EXPECT().MarkShipped(mock.Anything, int64(42)).Return(nil)

	err := ShipmentCreatedHandler(api)(context.Background(),
		record(t, &shippingv1.ShipmentCreated{EventId: "e1", OrderId: 42}))

	assert.NoError(t, err)
}

func TestShipmentCreatedHandler_UndecodableIsPermanent(t *testing.T) {
	err := ShipmentCreatedHandler(mocks.NewMockAPIPort(t))(context.Background(),
		&kgo.Record{Value: []byte{0xff, 0xff}})

	assert.ErrorIs(t, err, ErrPermanent)
}

func TestShipmentCreatedHandler_UnknownOrderIsPermanent(t *testing.T) {
	api := mocks.NewMockAPIPort(t)
	api.EXPECT().MarkShipped(mock.Anything, int64(9)).Return(domain.ErrNotFound)

	err := ShipmentCreatedHandler(api)(context.Background(),
		record(t, &shippingv1.ShipmentCreated{EventId: "e2", OrderId: 9}))

	assert.ErrorIs(t, err, ErrPermanent)
}
