package kafka

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/janapc/event-tickets/clients/internal/infra/telemetry"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type KafkaClient struct {
	Brokers []string
	readers []*kafka.Reader
}

func NewKafkaClient(brokers []string) *KafkaClient {
	return &KafkaClient{
		Brokers: brokers,
	}
}

func (k *KafkaClient) Producer(topic string, key, value []byte, ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}
	writer := kafka.Writer{
		Addr:                   kafka.TCP(k.Brokers...),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()
	ctx, span := telemetry.Tracer.Start(ctx, "produce-message",
		trace.WithAttributes(
			attribute.String("messaging.system", "kafka"),
			attribute.String("messaging.destination", topic),
			attribute.String("messaging.destination_kind", "topic"),
			attribute.String("messaging.operation", "send"),
			attribute.String("messaging.kafka.message_key", string(key)),
			attribute.Int("messaging.kafka.message_size", len(value)),
		),
	)
	defer span.End()
	err := writer.WriteMessages(ctx, kafka.Message{
		Key:   key,
		Value: value,
	})
	if err != nil {
		slog.Error("Error writing message to topic", "topic", topic, "error", err)
		return err
	}
	slog.Info("Message written to topic", "topic", topic, "key", string(key))
	return nil
}

func (k *KafkaClient) Consumer(topic, groupID string, handler func(context.Context, string) error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  k.Brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 1e3,
		MaxBytes: 10e6,
	})
	k.readers = append(k.readers, reader)

	go func() {
		slog.Info("Starting Kafka consumer", "topic", topic, "groupID", groupID)
		for {
			ctx := context.Background()
			msg, err := reader.ReadMessage(ctx)

			if err != nil {
				slog.Error("Error reading from", "topic ", topic, "error", err)
				continue
			}
			ctx, span := telemetry.Tracer.Start(ctx, "kafka-consume",
				trace.WithAttributes(
					attribute.String("messaging.system", "kafka"),
					attribute.String("messaging.destination", msg.Topic),
					attribute.String("messaging.destination_kind", "topic"),
					attribute.String("messaging.operation", "receive"),
					attribute.String("messaging.kafka.message_key", string(msg.Key)),
					attribute.Int("messaging.kafka.partition", msg.Partition),
					attribute.Int64("messaging.kafka.offset", msg.Offset),
					attribute.Int("messaging.kafka.message_size", len(msg.Value)),
				),
			)
			slog.Info("Received message", "topic", topic, "key", string(msg.Key))
			if err := handler(ctx, string(msg.Value)); err != nil {
				slog.Info("Error processing message", "topic", topic, "error", err)
			}
			span.End()
		}
	}()
}

func (k *KafkaClient) WaitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down Kafka readers...")
	for _, r := range k.readers {
		r.Close()
	}
}
