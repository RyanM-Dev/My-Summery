package api

import (
	"context"
	"errors"
	"testing"

	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"github.com/huseyinbabal/microservices/order/internal/ports/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func newOrder() domain.Order {
	return domain.NewOrder(123, []domain.OrderItem{
		{ProductCode: "camera", UnitPrice: 12.3, Quantity: 3},
	})
}

// saveWithID makes the mocked Save assign an ID, like the real database.
func saveWithID(_ context.Context, order *domain.Order) { order.ID = 1 }

func Test_Should_Place_Order(t *testing.T) {
	payment := mocks.NewMockPaymentPort(t)
	db := mocks.NewMockDBPort(t)
	db.EXPECT().Save(mock.Anything, mock.Anything).Run(saveWithID).Return(nil)
	payment.EXPECT().Charge(mock.Anything, mock.Anything).Return(nil)
	db.EXPECT().MarkPaid(mock.Anything, mock.MatchedBy(func(o *domain.Order) bool {
		return o.ID == 1 && o.Status == domain.StatusPaid
	})).Return(nil)

	application := NewApplication(db, payment)
	order, err := application.PlaceOrder(context.Background(), newOrder())

	assert.NoError(t, err)
	assert.Equal(t, domain.StatusPaid, order.Status)
}

func Test_Should_Return_Error_When_Db_Persistence_Fail(t *testing.T) {
	payment := mocks.NewMockPaymentPort(t)
	db := mocks.NewMockDBPort(t)
	db.EXPECT().Save(mock.Anything, mock.Anything).Return(errors.New("connection error"))
	// Charge is never expected: nothing is charged for an order that wasn't saved.

	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(context.Background(), newOrder())

	assert.EqualError(t, err, "connection error")
}

func Test_Should_Return_Error_When_Payment_Fail(t *testing.T) {
	payment := mocks.NewMockPaymentPort(t)
	db := mocks.NewMockDBPort(t)
	db.EXPECT().Save(mock.Anything, mock.Anything).Run(saveWithID).Return(nil)
	payment.EXPECT().Charge(mock.Anything, mock.Anything).
		Return(status.Error(codes.InvalidArgument, "insufficient balance"))
	db.EXPECT().UpdateStatus(mock.Anything, int64(1), domain.StatusFailed).Return(nil)

	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(context.Background(), newOrder())

	st, _ := status.FromError(err)
	assert.Equal(t, codes.InvalidArgument, st.Code())
	assert.Equal(t, "order creation failed", st.Message())
	require.Len(t, st.Details(), 1)
	violation := st.Details()[0].(*errdetails.BadRequest).GetFieldViolations()[0]
	assert.Equal(t, "payment", violation.GetField())
	assert.Equal(t, "insufficient balance", violation.GetDescription())
}

func Test_Should_Keep_Unavailable_When_Payment_Is_Down(t *testing.T) {
	payment := mocks.NewMockPaymentPort(t)
	db := mocks.NewMockDBPort(t)
	db.EXPECT().Save(mock.Anything, mock.Anything).Run(saveWithID).Return(nil)
	payment.EXPECT().Charge(mock.Anything, mock.Anything).
		Return(status.Error(codes.Unavailable, "connection refused"))
	db.EXPECT().UpdateStatus(mock.Anything, int64(1), domain.StatusFailed).Return(nil)

	application := NewApplication(db, payment)
	_, err := application.PlaceOrder(context.Background(), newOrder())

	assert.Equal(t, codes.Unavailable, status.Code(err))
}

func Test_Should_Mark_Order_Shipped(t *testing.T) {
	db := mocks.NewMockDBPort(t)
	db.EXPECT().UpdateStatus(mock.Anything, int64(7), domain.StatusShipped).Return(nil)

	application := NewApplication(db, mocks.NewMockPaymentPort(t))

	assert.NoError(t, application.MarkShipped(context.Background(), 7))
}

func TestTotalPrice(t *testing.T) {
	order := domain.NewOrder(1, []domain.OrderItem{
		{ProductCode: "a", UnitPrice: 12.5, Quantity: 2},
		{ProductCode: "b", UnitPrice: 5, Quantity: 1},
	})
	assert.Equal(t, float32(30), order.TotalPrice())
}
