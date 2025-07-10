package main

import (
	commandPayment "github.com/janapc/event-tickets/payments/internal/application/payment/command"
	"github.com/janapc/event-tickets/payments/internal/application/transaction/messaging"
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/internal/infra/email"
	"log"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/janapc/event-tickets/payments/internal/domain"
	"github.com/janapc/event-tickets/payments/internal/infra/database"
	"github.com/janapc/event-tickets/payments/internal/infra/messaging/kafka"
	"github.com/janapc/event-tickets/payments/internal/interfaces/http"
	"github.com/janapc/event-tickets/payments/pkg/eventbus"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Panic(err)
	}
	err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	defer database.DB.Close()

	kakfaClient := kafka.NewKafkaClient([]string{os.Getenv("KAFKA_BROKERS")})
	defer kakfaClient.WaitForShutdown()

	app := fiber.New()

	paymentRepo := database.NewPaymentRepository(database.DB)
	transactionRepo := database.NewTransactionRepository(database.DB)

	bus := eventbus.NewEventBus()
	registerEvents(bus, kakfaClient)

	api := http.NewPaymentHandler(commandPayment.NewCreatePaymentHandler(paymentRepo, bus))
	api.RegisterRoutes(app)

	onPaymentCreated := messaging.NewOnPaymentCreated(transactionRepo, bus)
	go kakfaClient.Consumer(os.Getenv("PAYMENT_CREATED_TOPIC"), os.Getenv("KAFKA_GROUP_ID"), onPaymentCreated.Handle)

	onTransactionCreated := messaging.NewOnTransactionCreated(transactionRepo, bus)
	go kakfaClient.Consumer(os.Getenv("TRANSACTION_CREATED_TOPIC"), os.Getenv("KAFKA_GROUP_ID"), onTransactionCreated.Handle)

	onTransactionFailed := messaging.NewOnTransactionFailed(paymentRepo, bus)
	go kakfaClient.Consumer(os.Getenv("TRANSACTION_FAILED_TOPIC"), os.Getenv("KAFKA_GROUP_ID"), onTransactionFailed.Handle)

	onTransactionSucceeded := messaging.NewOnTransactionSucceeded(paymentRepo, bus)
	go kakfaClient.Consumer(os.Getenv("TRANSACTION_SUCCEEDED_TOPIC"), os.Getenv("KAFKA_GROUP_ID"), onTransactionSucceeded.Handle)

	app.Listen(":3000")
}

func registerEvents(bus *eventbus.EventBus, kakfaClient *kafka.Client) {
	bus.Subscribe(payment.CreatedEventName, func(e eventbus.Event) {
		event := e.(*payment.CreatedEvent)
		topic := os.Getenv("PAYMENT_CREATED_TOPIC")
		params := domain.ProducerParameters{
			Topic: topic,
			Key:   uuid.NewString(),
			Value: event,
		}
		err := kakfaClient.Producer(params)
		if err != nil {
			slog.Error("kafka producer error: ", err)
		}
	})

	bus.Subscribe(transaction.CreatedEventName, func(e eventbus.Event) {
		event := e.(*transaction.CreatedEvent)
		topic := os.Getenv("TRANSACTION_CREATED_TOPIC")
		params := domain.ProducerParameters{
			Topic: topic,
			Key:   uuid.NewString(),
			Value: event,
		}
		err := kakfaClient.Producer(params)
		if err != nil {
			slog.Error("kafka producer error: ", err)
		}
	})

	bus.Subscribe(transaction.FailedEventName, func(e eventbus.Event) {
		event := e.(*transaction.FailedEvent)
		topic := os.Getenv("TRANSACTION_FAILED_TOPIC")
		params := domain.ProducerParameters{
			Topic: topic,
			Key:   uuid.NewString(),
			Value: event,
		}
		err := kakfaClient.Producer(params)
		if err != nil {
			slog.Error("kafka producer error: ", err)
		}
	})

	bus.Subscribe(transaction.SucceededEventName, func(e eventbus.Event) {
		event := e.(*transaction.SucceededEvent)
		topic := os.Getenv("TRANSACTION_SUCCEEDED_TOPIC")
		params := domain.ProducerParameters{
			Topic: topic,
			Key:   uuid.NewString(),
			Value: event,
		}
		err := kakfaClient.Producer(params)
		if err != nil {
			slog.Error("kafka producer error: ", err)
		}
	})

	bus.Subscribe(payment.FailedEventName, func(e eventbus.Event) {
		event := e.(*payment.FailedEvent)
		bodyEmail := email.IntlPaymentFailed(event.UserLanguage, event.UserName)
		email.SendEmail(event.UserEmail, bodyEmail.Subject, bodyEmail.Message)
	})

	bus.Subscribe(payment.SucceededEventName, func(e eventbus.Event) {
		event := e.(*payment.SucceededEvent)
		topic := os.Getenv("PAYMENT_SUCCEEDED_TOPIC")
		params := domain.ProducerParameters{
			Topic: topic,
			Key:   uuid.NewString(),
			Value: event,
		}
		err := kakfaClient.Producer(params)
		if err != nil {
			slog.Error("kafka producer error: ", err)
		}
		bodyEmail := email.IntlPaymentSucceeded(event.UserLanguage, event.UserName)
		email.SendEmail(event.UserEmail, bodyEmail.Subject, bodyEmail.Message)
	})
}
