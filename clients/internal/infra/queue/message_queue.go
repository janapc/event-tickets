package queue

import (
	"context"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/janapc/event-tickets/clients/internal/application"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Queue struct {
	Channel        *amqp.Channel
	ProcessMessage *application.ProcessMessage
	Logger         *slog.Logger
}

type Message struct {
	Pattern string `json:"pattern"`
	Data    application.InputProcessMessage
}

func NewQueue(channel *amqp.Channel, processMessage *application.ProcessMessage, logger *slog.Logger) *Queue {
	return &Queue{
		Channel:        channel,
		ProcessMessage: processMessage,
		Logger:         logger,
	}
}

func (q *Queue) Consumer(queueName string, workerPoolSize int) {
	deliveries, err := q.Channel.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		q.Logger.Error("error get messages of queue", "error", err)
	}
	q.Logger.Info("initialize queue message capture")
	for i := 0; i < workerPoolSize; i++ {
		go q.Worker(deliveries)
	}
}

func (q *Queue) Worker(messages <-chan amqp.Delivery) {
	for delivery := range messages {
		q.Logger.Info(string(delivery.Body))
		var input Message
		_ = json.Unmarshal(delivery.Body, &input)
		q.Logger.Info("delivery message", "messageId", delivery.MessageId, "body", input)
		err := q.ProcessMessage.Execute(input.Data, q)
		if err != nil {
			q.Logger.Error("error in save client", "error", err)
		} else {
			_ = delivery.Ack(false)
		}
	}
}

func (q *Queue) Producer(queueName string, message []byte) error {
	_, err := q.Channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = q.Channel.PublishWithContext(ctx,
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		})
	if err != nil {
		return err
	}
	q.Logger.Info("publish in queue", "message", string(message), "queue", queueName)
	return nil

}
