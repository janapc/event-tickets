package queue

import (
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ() (*amqp.Connection, *amqp.Channel) {
	amqpServerURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(amqpServerURL)
	if err != nil {
		panic(err)
	}
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	amqpQueueName := os.Getenv("QUEUE_PAYMENT")
	_, err = channel.QueueDeclare(
		amqpQueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	return conn, channel
}
