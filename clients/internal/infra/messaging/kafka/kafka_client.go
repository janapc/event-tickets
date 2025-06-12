package kafka

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
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

func (k *KafkaClient) Producer(topic string, key, value []byte) error {
	writer := kafka.Writer{
		Addr:                   kafka.TCP(k.Brokers...),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()

	err := writer.WriteMessages(context.Background(), kafka.Message{
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

func (k *KafkaClient) Consumer(topic, groupID string, handler func(string) error) {
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
			msg, err := reader.ReadMessage(context.Background())
			if err != nil {
				slog.Error("Error reading from", "topic ", topic, "error", err)
				continue
			}
			slog.Info("Received message", "topic", topic, "key", string(msg.Key))
			if err := handler(string(msg.Value)); err != nil {
				slog.Info("Error processing message", "topic", topic, "error", err)
			}
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
