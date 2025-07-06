package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func main() {
	const (
		kafkaAddr = "localhost:9093"
		topic     = "orders"
		interval  = 5 * time.Second // интервал между сообщениями
	)

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaAddr},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})
	defer writer.Close()

	rand.Seed(time.Now().UnixNano())

	for {
		order := generateRandomOrder()
		data, err := json.Marshal(order)
		if err != nil {
			log.Println("failed to marshal order:", err)
			continue
		}

		err = writer.WriteMessages(context.Background(), kafka.Message{
			Key:   []byte(order["order_uid"].(string)),
			Value: data,
		})
		if err != nil {
			log.Println("failed to write message:", err)
		} else {
			log.Println("order sent:", order["order_uid"])
		}

		time.Sleep(interval)
	}
}

func generateRandomOrder() map[string]interface{} {
	uid := uuid.NewString()
	return map[string]interface{}{
		"order_uid":          uid,
		"track_number":       "TRACK" + strconv.Itoa(rand.Intn(100000)),
		"entry":              "WBIL",
		"locale":             "en",
		"internal_signature": "",
		"customer_id":        "customer_" + strconv.Itoa(rand.Intn(1000)),
		"delivery_service":   "dhl",
		"shardkey":           strconv.Itoa(rand.Intn(10)),
		"sm_id":              rand.Intn(100),
		"date_created":       time.Now().UTC().Format(time.RFC3339),
		"oof_shard":          strconv.Itoa(rand.Intn(10)),
		"delivery": map[string]interface{}{
			"name":    "Name " + strconv.Itoa(rand.Intn(100)),
			"phone":   "+9720000000",
			"zip":     "1234567",
			"city":    "City",
			"address": "Street " + strconv.Itoa(rand.Intn(100)),
			"region":  "Region",
			"email":   "example@mail.com",
		},
		"payment": map[string]interface{}{
			"transaction":   uid,
			"request_id":    "",
			"currency":      "USD",
			"provider":      "wbpay",
			"amount":        rand.Intn(5000),
			"payment_dt":    time.Now().Unix(),
			"bank":          "alpha",
			"delivery_cost": 500,
			"goods_total":   200,
			"custom_fee":    0,
		},
		"items": []map[string]interface{}{
			{
				"chrt_id":      rand.Int63n(10000000),
				"track_number": "TRACK" + strconv.Itoa(rand.Intn(100000)),
				"price":        rand.Intn(1000),
				"rid":          uuid.NewString(),
				"name":         "Item " + strconv.Itoa(rand.Intn(10)),
				"sale":         rand.Intn(50),
				"size":         "M",
				"total_price":  rand.Intn(1000),
				"nm_id":        rand.Int63n(1000000),
				"brand":        "BrandX",
				"status":       202,
			},
		},
	}
}
