package main

import (
	"context"

	"log/slog"
	"os"

	"github.com/janapc/event-tickets/payments/internal/adapter/eventbus"
	commandPayment "github.com/janapc/event-tickets/payments/internal/application/payment/command"
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/internal/infra/email"
	"github.com/janapc/event-tickets/payments/internal/interfaces/messaging"
	"github.com/janapc/event-tickets/payments/pkg/kafka"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/janapc/event-tickets/payments/internal/domain"
	"github.com/janapc/event-tickets/payments/internal/infra/database"
	"github.com/janapc/event-tickets/payments/internal/interfaces/http"
	pkgEventBus "github.com/janapc/event-tickets/payments/pkg/eventbus"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
	err := database.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	defer database.DB.Close()
	ctx := context.Background()

	kakfaClient := kafka.NewKafkaClient([]string{os.Getenv("KAFKA_BROKERS")})
	defer kakfaClient.WaitForShutdown()

	app := fiber.New()

	paymentRepo := database.NewPaymentRepository(database.DB)
	transactionRepo := database.NewTransactionRepository(database.DB)

	bus := eventbus.NewBusAdapter(pkgEventBus.NewEventBus())

	registerEvents(bus, kakfaClient)

	api := http.NewPaymentHandler(commandPayment.NewCreatePaymentHandler(paymentRepo, bus))
	api.RegisterRoutes(app)

	messaging.RegisterAllConsumers(ctx, kakfaClient, transactionRepo, paymentRepo, bus)

	port := os.Getenv("PORT")
	app.Listen(port)
}

func registerEvents(bus domain.IEventBus, kakfaClient *kafka.Client) {
	bus.Subscribe(payment.CreatedEventName, func(e domain.Event) {
		event := e.(*payment.CreatedEvent)
		topic := os.Getenv("PAYMENT_CREATED_TOPIC")
		params := domain.ProducerParameters{
			Topic: topic,
			Key:   uuid.NewString(),
			Value: event,
		}
		err := kakfaClient.Producer(params)
		if err != nil {
			slog.Error("kafka producer error: ", "err", err)
		}
	})

	bus.Subscribe(transaction.CreatedEventName, func(e domain.Event) {
		event := e.(*transaction.CreatedEvent)
		topic := os.Getenv("TRANSACTION_CREATED_TOPIC")
		params := domain.ProducerParameters{
			Topic: topic,
			Key:   uuid.NewString(),
			Value: event,
		}
		err := kakfaClient.Producer(params)
		if err != nil {
			slog.Error("kafka producer error: ", "err", err)
		}
	})

	bus.Subscribe(transaction.FailedEventName, func(e domain.Event) {
		event := e.(*transaction.FailedEvent)
		topic := os.Getenv("TRANSACTION_FAILED_TOPIC")
		params := domain.ProducerParameters{
			Topic: topic,
			Key:   uuid.NewString(),
			Value: event,
		}
		err := kakfaClient.Producer(params)
		if err != nil {
			slog.Error("kafka producer error: ", "err", err)
		}
	})

	bus.Subscribe(transaction.SucceededEventName, func(e domain.Event) {
		event := e.(*transaction.SucceededEvent)
		topic := os.Getenv("TRANSACTION_SUCCEEDED_TOPIC")
		params := domain.ProducerParameters{
			Topic: topic,
			Key:   uuid.NewString(),
			Value: event,
		}
		err := kakfaClient.Producer(params)
		if err != nil {
			slog.Error("kafka producer error: ", "err", err)
		}
	})

	bus.Subscribe(payment.FailedEventName, func(e domain.Event) {
		event := e.(*payment.FailedEvent)
		bodyEmail := email.IntlPaymentFailed(event.UserLanguage, event.UserName)
		email.SendEmail(event.UserEmail, bodyEmail.Subject, bodyEmail.Message)
	})

	bus.Subscribe(payment.SucceededEventName, func(e domain.Event) {
		event := e.(*payment.SucceededEvent)
		topic := os.Getenv("PAYMENT_SUCCEEDED_TOPIC")
		params := domain.ProducerParameters{
			Topic: topic,
			Key:   uuid.NewString(),
			Value: event,
		}
		err := kakfaClient.Producer(params)
		if err != nil {
			slog.Error("kafka producer error: ", "err", err)
		}
		bodyEmail := email.IntlPaymentSucceeded(event.UserLanguage, event.UserName)
		email.SendEmail(event.UserEmail, bodyEmail.Subject, bodyEmail.Message)
	})
}
