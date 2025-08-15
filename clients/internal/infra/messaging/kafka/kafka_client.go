package kafka

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/janapc/event-tickets/clients/internal/infra/logger"
	"github.com/segmentio/kafka-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
)

type KafkaClient struct {
	Brokers []string
	readers []*kafka.Reader
	writer  *kafka.Writer
}

func NewKafkaClient(brokers []string) *KafkaClient {
	return &KafkaClient{
		Brokers: brokers,
		writer: &kafka.Writer{
			Addr:     kafka.TCP(brokers...),
			Balancer: &kafka.LeastBytes{},
		},
	}
}

func (k *KafkaClient) Producer(message []byte, ctx context.Context, topic string, key string) error {
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
	logger.Logger.WithContext(ctx).Infof("Kafka message %s sent to topic %s", key, topic)
	return nil
}

func (k *KafkaClient) Consumer(topic, groupID string, handler func(context.Context, string) error) {
	tracer := otel.Tracer("kafka-consumer")
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        k.Brokers,
		GroupID:        groupID,
		Topic:          topic,
		MinBytes:       1,
		MaxBytes:       10e6,
		MaxWait:        100 * time.Millisecond,
		CommitInterval: 50 * time.Millisecond,
	})
	k.readers = append(k.readers, reader)

	go func() {
		logger.Logger.Infof("Starting Kafka consumer topic %s groupID %s", topic, groupID)
		for {
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				logger.Logger.Errorf("Error reading from topic %s error %v", topic, err)
				continue
			}

			carrier := propagation.MapCarrier{}
			for _, h := range msg.Headers {
				carrier[h.Key] = string(h.Value)
			}
			ctx := otel.GetTextMapPropagator().Extract(context.Background(), carrier)

			ctx, span := tracer.Start(ctx, "consume-kafka-message")
			span.SetAttributes(attribute.String("kafka.topic", msg.Topic))
			defer span.End()

			logger.Logger.WithContext(ctx).Infof("Received message topic %s key %s", topic, string(msg.Key))
			if err := handler(ctx, string(msg.Value)); err != nil {
				logger.Logger.WithContext(ctx).Errorf("Error processing message topic %s error %v", topic, err)
			}
			logger.Logger.WithContext(ctx).Infof("Message processed topic %s key %s", topic, string(msg.Key))
		}
	}()
}

func (k *KafkaClient) WaitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	for _, r := range k.readers {
		logger.Logger.Printf("Shutting down Kafka reader for topic %s", r.Config().Topic)
		err := r.Close()
		if err != nil {
			logger.Logger.Errorf("Error closing Kafka reader for topic %s: %v", r.Config().Topic, err)
		} else {
			logger.Logger.Printf("Kafka reader for topic %s closed successfully", r.Config().Topic)
		}
	}
}
