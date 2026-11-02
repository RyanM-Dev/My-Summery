package db

import (
	"context"
	"testing"

	shippingv1 "github.com/huseyinbabal/microservices-proto/gen/go/shipping/v1"
	"github.com/huseyinbabal/microservices/shipping/internal/application/core/domain"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"google.golang.org/protobuf/proto"
)

type ShippingDatabaseTestSuite struct {
	suite.Suite
	container     *mysql.MySQLContainer
	DataSourceURL string
}

func (s *ShippingDatabaseTestSuite) SetupSuite() {
	ctx := context.Background()
	container, err := mysql.Run(ctx, "mysql:8.4",
		mysql.WithDatabase("shipments"),
		mysql.WithUsername("root"),
		mysql.WithPassword("s3cr3t"),
	)
	s.Require().NoError(err, "failed to start MySQL")
	s.container = container
	dsn, err := container.ConnectionString(ctx, "charset=utf8mb4", "parseTime=True", "loc=Local")
	s.Require().NoError(err)
	s.DataSourceURL = dsn
}

func (s *ShippingDatabaseTestSuite) TearDownSuite() {
	s.NoError(testcontainers.TerminateContainer(s.container))
}

func (s *ShippingDatabaseTestSuite) Test_Same_Event_Creates_One_Shipment_And_One_Event() {
	adapter, err := NewAdapter(s.DataSourceURL)
	s.Require().NoError(err)
	ctx := context.Background()

	for i := range 3 {
		shipment := domain.NewShipment(42)
		created, err := adapter.CreateShipment(ctx, "4f1c2b6e-0000-4000-8000-000000000042", &shipment)
		s.Require().NoError(err)
		s.Equal(i == 0, created, "only the first delivery creates a shipment")
	}

	got, err := adapter.GetByOrderID(ctx, 42)
	s.Require().NoError(err)
	s.Equal(domain.StatusCreated, got.Status)

	var published []OutboxEvent
	n, err := adapter.PublishOutbox(ctx, 100, func(_ context.Context, events []OutboxEvent) error {
		published = append(published, events...)
		return nil
	})
	s.Require().NoError(err)
	s.Equal(1, n, "one ShipmentCreated for three deliveries")
	s.Equal("shipments", published[0].Topic)
	s.Equal("42", published[0].Key)

	var event shippingv1.ShipmentCreated
	s.Require().NoError(proto.Unmarshal(published[0].Payload, &event))
	s.Equal(got.ID, event.GetShipmentId())
	s.Equal(int64(42), event.GetOrderId())
}

func (s *ShippingDatabaseTestSuite) Test_Should_Return_NotFound() {
	adapter, err := NewAdapter(s.DataSourceURL)
	s.Require().NoError(err)

	_, err = adapter.GetByOrderID(context.Background(), 999_999)
	s.ErrorIs(err, domain.ErrNotFound)
}

func TestShippingDatabaseTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("integration test: needs Docker")
	}
	suite.Run(t, new(ShippingDatabaseTestSuite))
}
