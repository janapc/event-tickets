package main

import (
	"context"

	"os"

	"github.com/janapc/event-tickets/payments/internal/adapter/eventbus"
	commandPayment "github.com/janapc/event-tickets/payments/internal/application/payment/command"
	commandTransaction "github.com/janapc/event-tickets/payments/internal/application/transaction/command"
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/internal/infra/email"
	"github.com/janapc/event-tickets/payments/internal/infra/logger"
	"github.com/janapc/event-tickets/payments/internal/infra/telemetry"
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
	logger.Init()
	if os.Getenv("ENV") != "PROD" {
		if err := godotenv.Load(); err != nil {
			logger.Logger.Panic(err)
		}
	}
	ctx := context.Background()
	env := os.Getenv("ENV")
	if env == "PROD" {
		err := telemetry.Init(ctx)
		if err != nil {
			logger.Logger.Panic(err)
		}
	}
	err := database.Init(ctx)
	if err != nil {
		logger.Logger.Panic(err)
	}
}

func main() {
	ctx := context.Background()
	defer database.Close(ctx)
	defer telemetry.Shutdown(ctx)

	kakfaClient := kafka.NewKafkaClient([]string{os.Getenv("KAFKA_BROKERS")})
	defer kakfaClient.WaitForShutdown()
	paymentRepo := database.NewPaymentRepository(database.DB)
	transactionRepo := database.NewTransactionRepository(database.DB)
	bus := eventbus.NewBusAdapter(pkgEventBus.NewEventBus())

	go kakfaClient.Consumer([]kafka.HandlerConfig{
		kafka.RegisterTypedHandler[commandPayment.FailPaymentCommand](
			os.Getenv("TRANSACTION_FAILED_TOPIC"),
			"group-transaction-failed",
			commandPayment.NewFailPaymentHandler(paymentRepo, bus).Handle,
		),
		kafka.RegisterTypedHandler[commandPayment.SuccessPaymentCommand](
			os.Getenv("TRANSACTION_SUCCEEDED_TOPIC"),
			"group-transaction-succeeded",
			commandPayment.NewSuccessPaymentHandler(paymentRepo, bus).Handle,
		),
		kafka.RegisterTypedHandler[commandTransaction.CreateTransactionCommand](
			os.Getenv("PAYMENT_CREATED_TOPIC"),
			"group-payment-created",
			commandTransaction.NewCreateTransactionHandler(transactionRepo, bus).Handle,
		),
		kafka.RegisterTypedHandler[commandTransaction.ProcessTransactionCommand](
			os.Getenv("TRANSACTION_CREATED_TOPIC"),
			"group-transaction-created",
			commandTransaction.NewProcessTransactionHandler(transactionRepo, bus).Handle,
		),
	})

	app := fiber.New()

	registerEvents(bus, kakfaClient)

	api := http.NewPaymentHandler(commandPayment.NewCreatePaymentHandler(paymentRepo, bus))
	api.RegisterRoutes(app)

	port := os.Getenv("PORT")
	if err := app.Listen(port); err != nil {
		logger.Logger.WithContext(context.Background()).Errorf("failed to start server: %v", err)
		return
	}
}

func registerEvents(bus domain.IEventBus, kakfaClient *kafka.Client) {
	bus.Subscribe(payment.CreatedEventName, func(e domain.Event) {
		event := e.(*payment.CreatedEvent)
		topic := os.Getenv("PAYMENT_CREATED_TOPIC")
		message, err := event.ToMessage()
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("error converting event to message: %v", err)
			return
		}
		err = kakfaClient.Producer(message, event.Context, topic, uuid.NewString())
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("kafka producer error: %v/n to eventName: %s", err, event.Name())
		}
	})

	bus.Subscribe(transaction.CreatedEventName, func(e domain.Event) {
		event := e.(*transaction.CreatedEvent)
		topic := os.Getenv("TRANSACTION_CREATED_TOPIC")
		message, err := event.ToMessage()
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("error converting event to message: %v", err)
			return
		}
		err = kakfaClient.Producer(message, event.Context, topic, uuid.NewString())
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("kafka producer error: %v/n to eventName: %s", err, event.Name())
		}
	})

	bus.Subscribe(transaction.FailedEventName, func(e domain.Event) {
		event := e.(*transaction.FailedEvent)
		topic := os.Getenv("TRANSACTION_FAILED_TOPIC")
		message, err := event.ToMessage()
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("error converting event to message: %v", err)
			return
		}
		err = kakfaClient.Producer(message, event.Context, topic, uuid.NewString())
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("kafka producer error: %v/n to eventName: %s", err, event.Name())
		}
	})

	bus.Subscribe(transaction.SucceededEventName, func(e domain.Event) {
		event := e.(*transaction.SucceededEvent)
		topic := os.Getenv("TRANSACTION_SUCCEEDED_TOPIC")
		message, err := event.ToMessage()
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("error converting event to message: %v", err)
			return
		}
		err = kakfaClient.Producer(message, event.Context, topic, uuid.NewString())
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("kafka producer error: %v/n to eventName: %s", err, event.Name())
		}
	})

	bus.Subscribe(payment.FailedEventName, func(e domain.Event) {
		event := e.(*payment.FailedEvent)
		bodyEmail := email.IntlPaymentFailed(event.Payload.UserLanguage, event.Payload.UserName)
		err := email.SendEmail(event.Context, event.Payload.UserEmail, bodyEmail.Subject, bodyEmail.Message)
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("email send error: %v/n to eventName: %s", err, event.Name())
		}
	})

	bus.Subscribe(payment.SucceededEventName, func(e domain.Event) {
		event := e.(*payment.SucceededEvent)
		topic := os.Getenv("PAYMENT_SUCCEEDED_TOPIC")
		message, err := event.ToMessage()
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("error converting event to message: %v", err)
			return
		}
		err = kakfaClient.Producer(message, event.Context, topic, uuid.NewString())
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("kafka producer error: %v/n to eventName: %s", err, event.Name())
		}
		bodyEmail := email.IntlPaymentSucceeded(event.Payload.UserLanguage, event.Payload.UserName)
		err = email.SendEmail(event.Context, event.Payload.UserEmail, bodyEmail.Subject, bodyEmail.Message)
		if err != nil {
			logger.Logger.WithContext(event.Context).Errorf("email send error: %v/n to eventName: %s", err, event.Name())
		}
	})
}
