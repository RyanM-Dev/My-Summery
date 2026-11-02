package db

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/google/uuid"
	shippingv1 "github.com/huseyinbabal/microservices-proto/gen/go/shipping/v1"
	"github.com/huseyinbabal/microservices/shipping/internal/application/core/domain"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TopicShipments carries shipping.v1 events, keyed by order ID.
const TopicShipments = "shipments"

type Shipment struct {
	gorm.Model
	OrderID int64 `gorm:"uniqueIndex"` // one shipment per order
	Status  string
}

// ProcessedEvent records every consumed event ID, so a redelivered event is
// recognized and skipped.
type ProcessedEvent struct {
	EventID     string    `gorm:"primaryKey;size:36"`
	ProcessedAt time.Time `gorm:"autoCreateTime"`
}

// OutboxEvent is an event waiting to be published to Kafka.
type OutboxEvent struct {
	ID          uint64 `gorm:"primaryKey"`
	Topic       string `gorm:"size:255"`
	Key         string `gorm:"size:255"`
	EventType   string `gorm:"size:255"`
	Payload     []byte
	Headers     []byte // JSON: trace context of the event that caused this one
	CreatedAt   time.Time
	PublishedAt *time.Time `gorm:"index"`
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceURL string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dataSourceURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("db connection error: %w", err)
	}
	if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName("shipments"))); err != nil {
		return nil, fmt.Errorf("db tracing plugin error: %w", err)
	}
	if err := db.AutoMigrate(&Shipment{}, &ProcessedEvent{}, &OutboxEvent{}); err != nil {
		return nil, fmt.Errorf("db migration error: %w", err)
	}
	return &Adapter{db: db}, nil
}

// CreateShipment does three writes in one transaction: remember the event ID,
// save the shipment, and queue ShipmentCreated. A known event ID makes it a
// no-op that returns false.
func (a *Adapter) CreateShipment(ctx context.Context, eventID string, shipment *domain.Shipment) (bool, error) {
	headers, err := traceHeaders(ctx)
	if err != nil {
		return false, err
	}

	created := false
	err = a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&ProcessedEvent{EventID: eventID})
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return nil // seen before
		}

		model := Shipment{OrderID: shipment.OrderID, Status: shipment.Status}
		if err := tx.Create(&model).Error; err != nil {
			return err
		}
		shipment.ID = int64(model.ID)
		shipment.CreatedAt = model.CreatedAt

		payload, err := proto.Marshal(&shippingv1.ShipmentCreated{
			EventId:    uuid.NewString(),
			ShipmentId: shipment.ID,
			OrderId:    shipment.OrderID,
			CreatedAt:  timestamppb.New(model.CreatedAt),
		})
		if err != nil {
			return err
		}
		created = true
		return tx.Create(&OutboxEvent{
			Topic:     TopicShipments,
			Key:       strconv.FormatInt(shipment.OrderID, 10),
			EventType: "shipping.v1.ShipmentCreated",
			Payload:   payload,
			Headers:   headers,
		}).Error
	})
	return created, err
}

func (a *Adapter) GetByOrderID(ctx context.Context, orderID int64) (domain.Shipment, error) {
	var model Shipment
	err := a.db.WithContext(ctx).Where("order_id = ?", orderID).First(&model).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return domain.Shipment{}, domain.ErrNotFound
	}
	if err != nil {
		return domain.Shipment{}, err
	}
	return domain.Shipment{
		ID:        int64(model.ID),
		OrderID:   model.OrderID,
		Status:    model.Status,
		CreatedAt: model.CreatedAt,
	}, nil
}

// PublishOutbox locks up to limit unpublished events, hands them to publish,
// and marks them published only if publish succeeds.
func (a *Adapter) PublishOutbox(ctx context.Context, limit int, publish func(context.Context, []OutboxEvent) error) (int, error) {
	var count int
	err := a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		var events []OutboxEvent
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
			Where("published_at IS NULL").Order("id").Limit(limit).
			Find(&events).Error; err != nil {
			return err
		}
		if len(events) == 0 {
			return nil
		}
		if err := publish(ctx, events); err != nil {
			return err
		}
		ids := make([]uint64, 0, len(events))
		for _, e := range events {
			ids = append(ids, e.ID)
		}
		count = len(events)
		return tx.Model(&OutboxEvent{}).Where("id IN ?", ids).Update("published_at", time.Now()).Error
	})
	return count, err
}

func traceHeaders(ctx context.Context) ([]byte, error) {
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	return json.Marshal(carrier)
}
