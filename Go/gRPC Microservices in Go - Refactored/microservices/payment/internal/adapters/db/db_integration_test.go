package db

import (
	"context"
	"testing"

	"github.com/huseyinbabal/microservices/payment/internal/application/core/domain"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
)

type PaymentDatabaseTestSuite struct {
	suite.Suite
	container     *mysql.MySQLContainer
	DataSourceURL string
}

// SetupSuite starts one MySQL container for all tests in the suite.
func (s *PaymentDatabaseTestSuite) SetupSuite() {
	ctx := context.Background()
	container, err := mysql.Run(ctx, "mysql:8.4",
		mysql.WithDatabase("payments"),
		mysql.WithUsername("root"),
		mysql.WithPassword("s3cr3t"),
	)
	s.Require().NoError(err, "failed to start MySQL")
	s.container = container

	dsn, err := container.ConnectionString(ctx, "charset=utf8mb4", "parseTime=True", "loc=Local")
	s.Require().NoError(err)
	s.DataSourceURL = dsn
}

func (s *PaymentDatabaseTestSuite) TearDownSuite() {
	s.NoError(testcontainers.TerminateContainer(s.container))
}

func (s *PaymentDatabaseTestSuite) Test_Should_Save_Payment() {
	adapter, err := NewAdapter(s.DataSourceURL)
	s.Require().NoError(err)

	payment := domain.NewPayment(1, 10, 99.5)
	s.NoError(adapter.Save(context.Background(), &payment))
	s.NotZero(payment.ID)
}

func (s *PaymentDatabaseTestSuite) Test_Should_Get_Payment() {
	adapter, err := NewAdapter(s.DataSourceURL)
	s.Require().NoError(err)

	payment := domain.NewPayment(2, 20, 12.25)
	s.Require().NoError(adapter.Save(context.Background(), &payment))

	got, err := adapter.Get(context.Background(), payment.ID)
	s.NoError(err)
	s.Equal(int64(2), got.CustomerID)
	s.Equal(int64(20), got.OrderID)
	s.Equal(float32(12.25), got.TotalPrice)
}

func (s *PaymentDatabaseTestSuite) Test_Should_Get_Payment_By_Order_ID() {
	adapter, err := NewAdapter(s.DataSourceURL)
	s.Require().NoError(err)

	payment := domain.NewPayment(3, 30, 5)
	s.Require().NoError(adapter.Save(context.Background(), &payment))

	got, err := adapter.GetByOrderID(context.Background(), 30)
	s.NoError(err)
	s.Equal(payment.ID, got.ID)
}

func (s *PaymentDatabaseTestSuite) Test_Should_Return_NotFound() {
	adapter, err := NewAdapter(s.DataSourceURL)
	s.Require().NoError(err)

	_, err = adapter.Get(context.Background(), 999_999)
	s.ErrorIs(err, domain.ErrNotFound)
}

func TestPaymentDatabaseTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("integration test: needs Docker")
	}
	suite.Run(t, new(PaymentDatabaseTestSuite))
}
