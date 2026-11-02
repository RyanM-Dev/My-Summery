package grpc

import (
	"context"
	"testing"

	orderv1 "github.com/huseyinbabal/microservices-proto/gen/go/order/v1"
	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/order/internal/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestValidateCreate(t *testing.T) {
	validItem := &orderv1.OrderItem{ProductCode: "CAM123", UnitPrice: 1.23, Quantity: 3}
	tests := []struct {
		name       string
		request    *orderv1.CreateRequest
		wantFields []string
	}{
		{name: "valid", request: &orderv1.CreateRequest{UserId: 23, OrderItems: []*orderv1.OrderItem{validItem}}},
		{name: "empty", request: &orderv1.CreateRequest{}, wantFields: []string{"user_id", "order_items"}},
		{
			name: "bad item",
			request: &orderv1.CreateRequest{UserId: 23, OrderItems: []*orderv1.OrderItem{
				{ProductCode: "", UnitPrice: 0, Quantity: 0},
			}},
			wantFields: []string{
				"order_items[0].product_code",
				"order_items[0].unit_price",
				"order_items[0].quantity",
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validateCreate(tc.request)
			if tc.wantFields == nil {
				assert.NoError(t, err)
				return
			}
			st := status.Convert(err)
			assert.Equal(t, codes.InvalidArgument, st.Code())
			require.Len(t, st.Details(), 1)
			var fields []string
			for _, v := range st.Details()[0].(*errdetails.BadRequest).GetFieldViolations() {
				fields = append(fields, v.GetField())
			}
			assert.Equal(t, tc.wantFields, fields)
		})
	}
}

func TestGet_NotFound(t *testing.T) {
	api := mocks.NewMockAPIPort(t)
	api.EXPECT().GetOrder(mock.Anything, int64(42)).Return(domain.Order{}, domain.ErrNotFound)

	adapter := &Adapter{api: api}
	_, err := adapter.Get(context.Background(), &orderv1.GetRequest{OrderId: 42})

	assert.Equal(t, codes.NotFound, status.Code(err))
}
