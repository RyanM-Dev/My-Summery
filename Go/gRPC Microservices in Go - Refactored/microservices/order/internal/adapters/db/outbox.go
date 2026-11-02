package db

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/google/uuid"
	orderv1 "github.com/huseyinbabal/microservices-proto/gen/go/order/v1"
	"github.com/huseyinbabal/microservices/order/internal/application/core/domain"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// TopicOrders carries order.v1 events, keyed by order ID.
const TopicOrders = "orders"

// OutboxEvent is an event waiting to be published to Kafka. It is written in
// the same transaction as the business change, so the two can't disagree.
type OutboxEvent struct {
	ID          uint64 `gorm:"primaryKey"`
	Topic       string `gorm:"size:255"`
	Key         string `gorm:"size:255"`
	EventType   string `gorm:"size:255"`
	Payload     []byte
	Headers     []byte // JSON: trace context of the request that caused the event
	CreatedAt   time.Time
	PublishedAt *time.Time `gorm:"index"`
}

// MarkPaid sets the order to Paid and queues OrderPaid in one transaction.
func (a *Adapter) MarkPaid(ctx context.Context, order *domain.Order) error {
	items := make([]*orderv1.OrderItem, 0, len(order.OrderItems))
	for _, item := range order.OrderItems {
		items = append(items, &orderv1.OrderItem{
			ProductCode: item.ProductCode,
			UnitPrice:   item.UnitPrice,
			Quantity:    item.Quantity,
		})
	}
	payload, err := proto.Marshal(&orderv1.OrderPaid{
		EventId:    uuid.NewString(),
		OrderId:    order.ID,
		UserId:     order.CustomerID,
		TotalPrice: order.TotalPrice(),
		OrderItems: items,
		PaidAt:     timestamppb.Now(),
	})
	if err != nil {
		return err
	}
	headers, err := traceHeaders(ctx)
	if err != nil {
		return err
	}

	return a.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		res := tx.Model(&Order{}).Where("id = ?", order.ID).Update("status", domain.StatusPaid)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			return domain.ErrNotFound
		}
		return tx.Create(&OutboxEvent{
			Topic:     TopicOrders,
			Key:       strconv.FormatInt(order.ID, 10),
			EventType: "order.v1.OrderPaid",
			Payload:   payload,
			Headers:   headers,
		}).Error
	})
}

// PublishOutbox locks up to limit unpublished events (SKIP LOCKED lets several
// relay replicas share the work), hands them to publish, and marks them
// published only if publish succeeds. It returns how many were published.
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
			return err // rollback: the rows stay unpublished and are retried
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

// traceHeaders captures the current trace context (W3C traceparent) so the
// relay can continue the same trace when it publishes later.
func traceHeaders(ctx context.Context) ([]byte, error) {
	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)
	return json.Marshal(carrier)
}
