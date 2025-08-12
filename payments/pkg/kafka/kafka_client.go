package kafka

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/trace"
)

type Client struct {
	Brokers  []string
	readers  []*kafka.Reader
	Handlers []HandlerConfig
	writer   *kafka.Writer
}

type NamedEvent interface {
	Name() string
}

type TypedHandler[T any] func(ctx context.Context, msg T) error

type HandlerConfig struct {
	Topic   string
	GroupID string
	Handle  func(ctx context.Context, b []byte)
}

func NewKafkaClient(brokers []string) *Client {
	return &Client{
		Brokers: brokers,
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (k *Client) Producer(message []byte, ctx context.Context, topic string, key string) error {
	var traceID string
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		traceID = span.SpanContext().TraceID().String()
	}

	sendCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := k.writer.WriteMessages(sendCtx, kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: message,
		Time:  time.Now(),
		Headers: []kafka.Header{
			{Key: "traceId", Value: []byte(traceID)},
		},
	})
	if err != nil {
		return err
	}
	log.Printf("Kafka message %s sent to topic %s", key, topic)
	return nil
}

func (k *Client) Consumer(handlers []HandlerConfig) {
	for _, entry := range handlers {
		go func() {
			reader := kafka.NewReader(kafka.ReaderConfig{
				Brokers:        k.Brokers,
				GroupID:        entry.GroupID,
				Topic:          entry.Topic,
				MinBytes:       1,
				MaxBytes:       10e6,
				MaxWait:        100 * time.Millisecond,
				CommitInterval: 50 * time.Millisecond,
			})
			k.readers = append(k.readers, reader)
			log.Printf("Kafka consumer started for topic: %s", entry.Topic)

			for {
				ctx := context.Background()

				m, err := reader.ReadMessage(ctx)
				log.Printf("Kafka consumer reading message from topic: %s %s", entry.Topic, string(m.Key))
				if err != nil {
					log.Printf("Error reading message from %s: %v", entry.Topic, err)
					continue
				}
				entry.Handle(ctx, m.Value)
			}
		}()
	}
}

func RegisterTypedHandler[T any](topic, groupID string, typedHandler TypedHandler[T]) HandlerConfig {
	return HandlerConfig{
		Topic:   topic,
		GroupID: groupID,
		Handle: func(ctx context.Context, b []byte) {
			var msg T
			if err := json.Unmarshal(b, &msg); err != nil {
				log.Printf("Unmarshal error on topic %s: %v", topic, err)
				return
			}
			if err := typedHandler(ctx, msg); err != nil {
				log.Printf("Handler error on topic %s: %v", topic, err)
			}
		},
	}
}

func (k *Client) WaitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Printf("Shutting down Kafka readers...")
	for _, r := range k.readers {
		err := r.Close()
		if err != nil {
			log.Fatalf("Kafka Error closing reader %v", err)
		}
	}
}
