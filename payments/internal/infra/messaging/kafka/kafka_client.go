package kafka

import (
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2/log"
	"os"
	"os/signal"
	"syscall"

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
	log.Debugf("send message %s to topic %s", params.Key, params.Topic)
	return nil
}

func (k *Client) Consumer(topic, groupID string, handler func(string) error) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  k.Brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 1e3,
		MaxBytes: 10e6,
	})
	k.readers = append(k.readers, reader)

	go func() {
		log.Debugf("Starting Kafka consumer topic %s groupID %s", topic, groupID)
		for {
			ctx := context.Background()
			msg, err := reader.ReadMessage(ctx)
			log.Debugf("Kafka consumer topic %s groupID %s", topic, groupID)

			if err != nil {
				log.Errorf("Error reading from topic %s error %v", topic, err)
				continue
			}

			if err := handler(string(msg.Value)); err != nil {
				log.Debugf("Error processing message topic %s error %v", topic, err)
			}
			log.Debugf("Message processed topic %s key %s", topic, string(msg.Key))
		}
	}()
}

func (k *Client) WaitForShutdown() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Debugf("Shutting down Kafka readers...")
	for _, r := range k.readers {
		err := r.Close()
		if err != nil {
			log.Errorf("Error closing reader %v", err)
		}
	}
}
