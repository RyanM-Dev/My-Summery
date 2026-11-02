package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/huseyinbabal/microservices/payment/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Payment is the database model; domain.Payment never leaves the core.
type Payment struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderID    int64 `gorm:"uniqueIndex"`
	TotalPrice float32
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceURL string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dataSourceURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}
	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("payments"))); err != nil {
		return nil, fmt.Errorf("db tracing plugin error: %w", err)
	}
	if err := db.AutoMigrate(&Payment{}); err != nil {
		return nil, fmt.Errorf("db migration error: %w", err)
	}
	return &Adapter{db: db}, nil
}

func (a *Adapter) Get(ctx context.Context, id int64) (domain.Payment, error) {
	var entity Payment
	err := a.db.WithContext(ctx).First(&entity, id).Error
	return toDomain(entity, err)
}

func (a *Adapter) GetByOrderID(ctx context.Context, orderID int64) (domain.Payment, error) {
	var entity Payment
	err := a.db.WithContext(ctx).Where("order_id = ?", orderID).First(&entity).Error
	return toDomain(entity, err)
}

func toDomain(entity Payment, err error) (domain.Payment, error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Payment{}, domain.ErrNotFound
	}
	if err != nil {
		return domain.Payment{}, err
	}
	return domain.Payment{
		ID:         int64(entity.ID),
		CustomerID: entity.CustomerID,
		Status:     entity.Status,
		OrderID:    entity.OrderID,
		TotalPrice: entity.TotalPrice,
		CreatedAt:  entity.CreatedAt.Unix(),
	}, nil
}

func (a *Adapter) Save(ctx context.Context, payment *domain.Payment) error {
	model := Payment{
		CustomerID: payment.CustomerID,
		Status:     payment.Status,
		OrderID:    payment.OrderID,
		TotalPrice: payment.TotalPrice,
	}
	res := a.db.WithContext(ctx).Create(&model)
	if res.Error != nil {
		return res.Error
	}
	payment.ID = int64(model.ID)
	return nil
}
