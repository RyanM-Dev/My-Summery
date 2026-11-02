package grpc

import (
	"testing"

	paymentv1 "github.com/huseyinbabal/microservices-proto/gen/go/payment/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestValidate(t *testing.T) {
	tests := []struct {
		name       string
		request    *paymentv1.CreateRequest
		wantFields []string
	}{
		{name: "valid", request: &paymentv1.CreateRequest{UserId: 1, OrderId: 1, TotalPrice: 10}},
		{name: "all fields invalid", request: &paymentv1.CreateRequest{}, wantFields: []string{"user_id", "order_id", "total_price"}},
		{name: "negative price", request: &paymentv1.CreateRequest{UserId: 1, OrderId: 1, TotalPrice: -2}, wantFields: []string{"total_price"}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := validate(tc.request)
			if tc.wantFields == nil {
				assert.NoError(t, err)
				return
			}

			st := status.Convert(err)
			assert.Equal(t, codes.InvalidArgument, st.Code())
			require.Len(t, st.Details(), 1)
			badRequest, ok := st.Details()[0].(*errdetails.BadRequest)
			require.True(t, ok)
			var fields []string
			for _, v := range badRequest.GetFieldViolations() {
				fields = append(fields, v.GetField())
			}
			assert.Equal(t, tc.wantFields, fields)
		})
	}
}
