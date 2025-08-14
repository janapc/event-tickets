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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
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
	sendCtx, cancel := context.WithTimeout(context.WithoutCancel(ctx), 10*time.Second)
	defer cancel()

	msg := kafka.Message{
		Topic:   topic,
		Key:     []byte(key),
		Value:   message,
		Time:    time.Now(),
		Headers: []kafka.Header{},
	}

	carrier := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(ctx, carrier)

	for k, v := range carrier {
		msg.Headers = append(msg.Headers, kafka.Header{
			Key:   k,
			Value: []byte(v),
		})
	}
	err := k.writer.WriteMessages(sendCtx, msg)
	if err != nil {
		return err
	}
	log.Printf("Kafka message %s sent to topic %s", key, topic)
	return nil
}

func (k *Client) Consumer(handlers []HandlerConfig) {
	tracer := otel.Tracer("kafka-consumer")
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

			log.Printf("Starting Kafka consumer for topic %s with group ID %s", entry.Topic, entry.GroupID)

			for {
				m, err := reader.ReadMessage(context.Background())
				if err != nil {
					log.Printf("Error reading message from topic %s: %v", entry.Topic, err)
					continue
				}
				carrier := propagation.MapCarrier{}
				for _, h := range m.Headers {
					carrier[h.Key] = string(h.Value)
				}
				ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)

				ctx, span := tracer.Start(ctx, "consume-kafka-message")
				span.SetAttributes(attribute.String("kafka.topic", m.Topic))
				defer span.End()
				log.Printf("Received message from topic %s: %s", m.Topic, string(m.Key))
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
				log.Printf("Error unmarshalling message for topic %s: %v", topic, err)
				return
			}
			if err := typedHandler(ctx, msg); err != nil {
				log.Printf("Error handling message for topic %s: %v", topic, err)
			}
		},
	}
}

func (k *Client) WaitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	for _, r := range k.readers {
		log.Printf("Shutting down Kafka reader for topic %s", r.Config().Topic)
		err := r.Close()
		if err != nil {
			log.Printf("Error closing Kafka reader for topic %s: %v", r.Config().Topic, err)
		} else {
			log.Printf("Kafka reader for topic %s closed successfully", r.Config().Topic)
		}
	}
}
