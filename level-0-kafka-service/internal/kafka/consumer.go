package kafka

import (
	"context"
	"encoding/json"
	"log/slog"

	"github.com/Cladkoewka/wb-technoschool/level-0-kafka-service/internal/domain"
	"github.com/Cladkoewka/wb-technoschool/level-0-kafka-service/internal/service"
	"github.com/segmentio/kafka-go"
)

type Consumer struct {
	reader  *kafka.Reader
	service *service.OrderService
}

func NewConsumer(broker, topic, groupID string, svc *service.OrderService) *Consumer {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID,
	})
	return &Consumer{
		reader:  r,
		service: svc,
	}
}

func (c *Consumer) Run(ctx context.Context) {
	for {
		m, err := c.reader.ReadMessage(ctx)
		if err != nil {
			slog.Error("failed to read message from kafka", "err", err)
			continue
		}

		var order domain.Order
		if err := json.Unmarshal(m.Value, &order); err != nil {
			slog.Warn("invalid order json", "msg", string(m.Value), "err", err)
			continue
		}

		if err := c.service.ProcessOrder(ctx, &order); err != nil {
			slog.Error("failed to save order", "order_uid", order.OrderUID, "err", err)
			continue
		}

		slog.Info("order processed", "order_uid", order.OrderUID)
	}
}
