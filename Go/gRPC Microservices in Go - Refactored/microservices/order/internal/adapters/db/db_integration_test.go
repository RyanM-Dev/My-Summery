package db

import (
	"context"
	"errors"
	"strconv"
	"testing"

	orderv1 "github.com/huseyinbabal/microservices-proto/gen/go/order/v1"
	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"google.golang.org/protobuf/proto"
)

type OrderDatabaseTestSuite struct {
	suite.Suite
	container     *mysql.MySQLContainer
	DataSourceURL string
}

// SetupSuite starts one MySQL container for all tests in the suite.
func (o *OrderDatabaseTestSuite) SetupSuite() {
	ctx := context.Background()
	container, err := mysql.Run(ctx, "mysql:8.4",
		mysql.WithDatabase("orders"),
		mysql.WithUsername("root"),
		mysql.WithPassword("s3cr3t"),
	)
	o.Require().NoError(err, "failed to start MySQL")
	o.container = container

	dsn, err := container.ConnectionString(ctx, "charset=utf8mb4", "parseTime=True", "loc=Local")
	o.Require().NoError(err)
	o.DataSourceURL = dsn
}

// TearDownSuite destroys the container after all tests.
func (o *OrderDatabaseTestSuite) TearDownSuite() {
	o.NoError(testcontainers.TerminateContainer(o.container))
}

func (o *OrderDatabaseTestSuite) Test_Should_Save_Order() {
	adapter, err := NewAdapter(o.DataSourceURL)
	o.Require().NoError(err)

	order := domain.NewOrder(1, []domain.OrderItem{{ProductCode: "CAM", Quantity: 1, UnitPrice: 2}})
	o.NoError(adapter.Save(context.Background(), &order))
	o.NotZero(order.ID)
}

func (o *OrderDatabaseTestSuite) Test_Should_Get_Order() {
	adapter, err := NewAdapter(o.DataSourceURL)
	o.Require().NoError(err)

	order := domain.NewOrder(2, []domain.OrderItem{{ProductCode: "CAM", Quantity: 5, UnitPrice: 1.32}})
	o.Require().NoError(adapter.Save(context.Background(), &order))

	got, err := adapter.Get(context.Background(), order.ID)
	o.NoError(err)
	o.Equal(int64(2), got.CustomerID)
	o.Equal(domain.StatusPending, got.Status)
	o.Require().Len(got.OrderItems, 1)
	o.Equal("CAM", got.OrderItems[0].ProductCode)
	o.Equal(int32(5), got.OrderItems[0].Quantity)
}

func (o *OrderDatabaseTestSuite) Test_Should_Update_Status() {
	adapter, err := NewAdapter(o.DataSourceURL)
	o.Require().NoError(err)

	order := domain.NewOrder(3, []domain.OrderItem{{ProductCode: "BAG", Quantity: 1, UnitPrice: 9}})
	o.Require().NoError(adapter.Save(context.Background(), &order))
	o.Require().NoError(adapter.UpdateStatus(context.Background(), order.ID, domain.StatusPaid))

	got, err := adapter.Get(context.Background(), order.ID)
	o.NoError(err)
	o.Equal(domain.StatusPaid, got.Status)
}

func (o *OrderDatabaseTestSuite) Test_Should_MarkPaid_And_Publish_Outbox_Once() {
	adapter, err := NewAdapter(o.DataSourceURL)
	o.Require().NoError(err)
	ctx := context.Background()

	order := domain.NewOrder(4, []domain.OrderItem{{ProductCode: "CAM", Quantity: 2, UnitPrice: 10}})
	o.Require().NoError(adapter.Save(ctx, &order))
	o.Require().NoError(adapter.MarkPaid(ctx, &order))

	got, err := adapter.Get(ctx, order.ID)
	o.Require().NoError(err)
	o.Equal(domain.StatusPaid, got.Status)

	// A failed publish rolls back: the event stays queued.
	_, err = adapter.PublishOutbox(ctx, 100, func(context.Context, []OutboxEvent) error {
		return errors.New("kafka down")
	})
	o.Error(err)

	var published []OutboxEvent
	n, err := adapter.PublishOutbox(ctx, 100, func(_ context.Context, events []OutboxEvent) error {
		published = append(published, events...)
		return nil
	})
	o.Require().NoError(err)
	o.Equal(1, n)
	o.Equal("orders", published[0].Topic)
	o.Equal(strconv.FormatInt(order.ID, 10), published[0].Key)

	var event orderv1.OrderPaid
	o.Require().NoError(proto.Unmarshal(published[0].Payload, &event))
	o.Equal(order.ID, event.GetOrderId())
	o.Equal(float32(20), event.GetTotalPrice())
	o.NotEmpty(event.GetEventId())

	// Already published: nothing left to send.
	n, err = adapter.PublishOutbox(ctx, 100, func(context.Context, []OutboxEvent) error { return nil })
	o.NoError(err)
	o.Zero(n)
}

func (o *OrderDatabaseTestSuite) Test_Should_Return_NotFound() {
	adapter, err := NewAdapter(o.DataSourceURL)
	o.Require().NoError(err)

	_, err = adapter.Get(context.Background(), 999_999)
	o.ErrorIs(err, domain.ErrNotFound)
}

func TestOrderDatabaseTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("integration test: needs Docker")
	}
	suite.Run(t, new(OrderDatabaseTestSuite))
}
