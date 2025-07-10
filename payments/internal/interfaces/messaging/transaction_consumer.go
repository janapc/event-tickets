package messaging

import (
	"github.com/janapc/event-tickets/payments/internal/application/transaction/command"
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/pkg/eventbus"
	"os"
)

func RegisterTransactionConsumers(
	transactionRepo transaction.ITransactionRepository,
	bus *eventbus.EventBus,
) []ConsumerRegistration {
	return []ConsumerRegistration{
		{Topic: os.Getenv("PAYMENT_CREATED_TOPIC"), GroupID: os.Getenv("KAFKA_GROUP_ID"), Handler: WrapCommandHandler[command.CreateTransactionCommand](command.NewCreateTransactionHandler(transactionRepo, bus).Handle)},
		{Topic: os.Getenv("TRANSACTION_CREATED_TOPIC"), GroupID: os.Getenv("KAFKA_GROUP_ID"), Handler: WrapCommandHandler[command.ProcessTransactionCommand](command.NewProcessTransactionHandler(transactionRepo, bus).Handle)},
	}
}
