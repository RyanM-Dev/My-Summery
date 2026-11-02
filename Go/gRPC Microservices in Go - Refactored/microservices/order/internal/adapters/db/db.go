package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Order and OrderItem are database models; domain types never leave the core.
type Order struct {
	gorm.Model
	CustomerID int64
	Status     string
	OrderItems []OrderItem
}

type OrderItem struct {
	gorm.Model
	ProductCode string
	UnitPrice   float32
	Quantity    int32
	OrderID     uint // back reference to Order
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceURL string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dataSourceURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}
	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("orders"))); err != nil {
		return nil, fmt.Errorf("db tracing plugin error: %w", err)
	}
	if err := db.AutoMigrate(&Order{}, &OrderItem{}, &OutboxEvent{}); err != nil {
		return nil, fmt.Errorf("db migration error: %w", err)
	}
	return &Adapter{db: db}, nil
}

func (a *Adapter) Get(ctx context.Context, id int64) (domain.Order, error) {
	var entity Order
	res := a.db.WithContext(ctx).Preload("OrderItems").First(&entity, id)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return domain.Order{}, domain.ErrNotFound
	}
	if res.Error != nil {
		return domain.Order{}, res.Error
	}

	orderItems := make([]domain.OrderItem, 0, len(entity.OrderItems))
	for _, item := range entity.OrderItems {
		orderItems = append(orderItems, domain.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	return domain.Order{
		ID:         int64(entity.ID),
		CustomerID: entity.CustomerID,
		Status:     entity.Status,
		OrderItems: orderItems,
		CreatedAt:  entity.CreatedAt.Unix(),
	}, nil
}

func (a *Adapter) Save(ctx context.Context, order *domain.Order) error {
	orderItems := make([]OrderItem, 0, len(order.OrderItems))
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	model := Order{
		CustomerID: order.CustomerID,
		Status:     order.Status,
		OrderItems: orderItems,
	}
	res := a.db.WithContext(ctx).Create(&model)
	if res.Error != nil {
		return res.Error
	}
	order.ID = int64(model.ID)
	return nil
}

func (a *Adapter) UpdateStatus(ctx context.Context, id int64, status string) error {
	res := a.db.WithContext(ctx).Model(&Order{}).Where("id = ?", id).Update("status", status)
	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return domain.ErrNotFound
	}
	return nil
}
