// Package e2e runs the whole stack from ../docker-compose.yml and calls the
// services like a real client: the sync path (Order → Payment over gRPC) and
// the async path (Order → Kafka → Shipping → Kafka → Order). Run `make e2e`.
package e2e

import (
	"context"
	"testing"
	"time"

	orderv1 "github.com/huseyinbabal/microservices-proto/gen/go/order/v1"
	shippingv1 "github.com/huseyinbabal/microservices-proto/gen/go/shipping/v1"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/testcontainers/testcontainers-go/wait"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type CreateOrderTestSuite struct {
	suite.Suite
	stack       *compose.DockerCompose
	conn           *grpc.ClientConn
	shippingConn   *grpc.ClientConn
	orderClient    orderv1.OrderServiceClient
	shippingClient shippingv1.ShippingServiceClient
}

// SetupSuite starts MySQL, Payment, Order, and Jaeger, and waits until both
// services log that they are serving.
func (c *CreateOrderTestSuite) SetupSuite() {
	ctx := context.Background()

	stack, err := compose.NewDockerComposeWith(compose.WithStackFiles("../docker-compose.yml"))
	c.Require().NoError(err)
	c.stack = stack

	err = stack.
		WaitForService("payment", wait.ForLog("starting payment service").WithStartupTimeout(2*time.Minute)).
		WaitForService("order", wait.ForLog("starting order service").WithStartupTimeout(2*time.Minute)).
		WaitForService("shipping", wait.ForLog("starting shipping service").WithStartupTimeout(2*time.Minute)).
		Up(ctx, compose.Wait(true))
	c.Require().NoError(err, "could not run compose stack")

	c.conn, err = grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	c.Require().NoError(err)
	c.orderClient = orderv1.NewOrderServiceClient(c.conn)

	c.shippingConn, err = grpc.NewClient("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	c.Require().NoError(err)
	c.shippingClient = shippingv1.NewShippingServiceClient(c.shippingConn)
}

// TearDownSuite removes the stack, its volumes, and its network.
func (c *CreateOrderTestSuite) TearDownSuite() {
	if c.conn != nil {
		c.NoError(c.conn.Close())
	}
	if c.shippingConn != nil {
		c.NoError(c.shippingConn.Close())
	}
	if c.stack != nil {
		c.NoError(c.stack.Down(context.Background(), compose.RemoveOrphans(true), compose.RemoveVolumes(true)))
	}
}

func (c *CreateOrderTestSuite) Test_Should_Create_Order() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	createResponse, err := c.orderClient.Create(ctx, &orderv1.CreateRequest{
		UserId: 23,
		OrderItems: []*orderv1.OrderItem{
			{ProductCode: "CAM123", Quantity: 3, UnitPrice: 1.23},
		},
	})
	c.Require().NoError(err)

	getResponse, err := c.orderClient.Get(ctx, &orderv1.GetRequest{OrderId: createResponse.GetOrderId()})
	c.Require().NoError(err)
	c.Equal(int64(23), getResponse.GetUserId())
	c.Contains([]string{"Paid", "Shipped"}, getResponse.GetStatus(), "paid synchronously")
	c.Require().Len(getResponse.GetOrderItems(), 1)
	orderItem := getResponse.GetOrderItems()[0]
	c.Equal(float32(1.23), orderItem.GetUnitPrice())
	c.Equal(int32(3), orderItem.GetQuantity())
	c.Equal("CAM123", orderItem.GetProductCode())
}

// The async half: OrderPaid → Shipping creates a shipment → ShipmentCreated →
// Order becomes Shipped. Nothing waits for it, so the test polls.
func (c *CreateOrderTestSuite) Test_Should_Ship_Paid_Order_Asynchronously() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	created, err := c.orderClient.Create(ctx, &orderv1.CreateRequest{
		UserId:     7,
		OrderItems: []*orderv1.OrderItem{{ProductCode: "BAG", Quantity: 1, UnitPrice: 30}},
	})
	c.Require().NoError(err)
	orderID := created.GetOrderId()

	c.Require().Eventually(func() bool {
		order, err := c.orderClient.Get(ctx, &orderv1.GetRequest{OrderId: orderID})
		return err == nil && order.GetStatus() == "Shipped"
	}, 20*time.Second, 200*time.Millisecond, "order never became Shipped")

	shipment, err := c.shippingClient.Get(ctx, &shippingv1.GetRequest{OrderId: orderID})
	c.Require().NoError(err)
	c.Equal(orderID, shipment.GetOrderId())
	c.Equal("Created", shipment.GetStatus())
}

func (c *CreateOrderTestSuite) Test_Should_Reject_Invalid_Order() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.orderClient.Create(ctx, &orderv1.CreateRequest{UserId: 0})

	st := status.Convert(err)
	c.Equal(codes.InvalidArgument, st.Code())
	c.Require().Len(st.Details(), 1)
	c.Len(st.Details()[0].(*errdetails.BadRequest).GetFieldViolations(), 2)
}

func (c *CreateOrderTestSuite) Test_Should_Return_NotFound() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := c.orderClient.Get(ctx, &orderv1.GetRequest{OrderId: 999_999})

	c.Equal(codes.NotFound, status.Code(err))
}

func TestCreateOrderTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("end-to-end test: needs Docker")
	}
	suite.Run(t, new(CreateOrderTestSuite))
}
