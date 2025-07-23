package messaging

import (
	"os"

	"github.com/janapc/event-tickets/payments/internal/application/payment/command"
	"github.com/janapc/event-tickets/payments/internal/domain"
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
)

func RegisterPaymentConsumers(
	paymentRepo payment.IPaymentRepository,
	bus domain.IEventBus,
) []ConsumerRegistration {
	return []ConsumerRegistration{
		{Topic: os.Getenv("TRANSACTION_FAILED_TOPIC"), GroupID: os.Getenv("KAFKA_GROUP_ID"), Handler: WrapCommandHandler[command.FailPaymentCommand](command.NewFailPaymentHandler(paymentRepo, bus).Handle)},
		{Topic: os.Getenv("TRANSACTION_SUCCEEDED_TOPIC"), GroupID: os.Getenv("KAFKA_GROUP_ID"), Handler: WrapCommandHandler[command.SuccessPaymentCommand](command.NewSuccessPaymentHandler(paymentRepo, bus).Handle)},
	}
}
