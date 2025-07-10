package kafka

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/janapc/event-tickets/payments/internal/domain"
	"github.com/segmentio/kafka-go"
)

type Client struct {
	Brokers []string
	readers []*kafka.Reader
}

func NewKafkaClient(brokers []string) *Client {
	return &Client{
		Brokers: brokers,
	}
}

func (k *Client) Producer(params domain.ProducerParameters) error {
	payload, err := json.Marshal(params.Value)
	if err != nil {
		return err
	}
	writer := kafka.Writer{
		Addr:                   kafka.TCP(k.Brokers...),
		Topic:                  params.Topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
	defer writer.Close()
	ctx := context.Background()
	err = writer.WriteMessages(ctx, kafka.Message{
		Key:   []byte(params.Key),
		Value: payload,
	})
	if err != nil {
		return err
	}
	log.Debugf("Kafka message %s sent to topic %s", params.Key, params.Topic)
	return nil
}

func (k *Client) Consumer(ctx context.Context, topic, groupID string, handler func(ctx context.Context, key string, value []byte) error) error {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  k.Brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 1, //dev
		MaxBytes: 10e6,
		MaxWait:  100 * time.Millisecond,
	})
	k.readers = append(k.readers, reader)

	go func() {
		log.Debugf("Starting Kafka consumer topic %s groupID %s", topic, groupID)
		for {
			msg, err := reader.ReadMessage(ctx)
			log.Debugf("Kafka consumer topic %s groupID %s", topic, groupID)

			if err != nil {
				log.Errorf("Kafka Error reading from topic %s error %v", topic, err)
				continue
			}

			if err := handler(ctx, string(msg.Key), msg.Value); err != nil {
				log.Debugf("KafkaError processing message topic %s error %v", topic, err)
			}
			log.Debugf("Kafka Message processed topic %s key %s", topic, string(msg.Key))
		}
	}()
	return nil
}

func (k *Client) WaitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Debugf("Shutting down Kafka readers...")
	for _, r := range k.readers {
		err := r.Close()
		if err != nil {
			log.Errorf("Kafka Error closing reader %v", err)
		}
	}
}
