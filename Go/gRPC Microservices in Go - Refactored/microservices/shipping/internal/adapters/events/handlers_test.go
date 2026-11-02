package events

import (
	"context"
	"testing"

	orderv1 "github.com/huseyinbabal/microservices-proto/gen/go/order/v1"
	"github.com/huseyinbabal/microservices/shipping/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/shipping/internal/ports/mocks"
	"github.com/stretchr/testify/assert"
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

func TestOrderPaidHandler_CreatesShipment(t *testing.T) {
	api := mocks.NewMockAPIPort(t)
	api.EXPECT().CreateShipment(context.Background(), "e1", int64(42)).Return(domain.Shipment{ID: 1}, nil)

	err := OrderPaidHandler(api)(context.Background(), record(t, &orderv1.OrderPaid{EventId: "e1", OrderId: 42}))

	assert.NoError(t, err)
}

func TestOrderPaidHandler_RejectsBadEvents(t *testing.T) {
	handler := OrderPaidHandler(mocks.NewMockAPIPort(t))

	assert.ErrorIs(t, handler(context.Background(), &kgo.Record{Value: []byte{0xff, 0xff}}), ErrPermanent)
	assert.ErrorIs(t, handler(context.Background(), record(t, &orderv1.OrderPaid{OrderId: 42})), ErrPermanent)
}
