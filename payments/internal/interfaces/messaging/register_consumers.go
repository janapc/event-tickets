package messaging

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/janapc/event-tickets/payments/internal/domain"
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/pkg/kafka"
)

type HandlerFunc func(ctx context.Context, key string, value []byte) error

type ConsumerRegistration struct {
	Topic   string
	GroupID string
	Handler HandlerFunc
}

func RegisterAllConsumers(
	ctx context.Context,
	kafka *kafka.Client,
	transactionRepo transaction.ITransactionRepository,
	paymentRepo payment.IPaymentRepository,
	bus domain.IEventBus,
) {
	var all []ConsumerRegistration

	all = append(all, RegisterTransactionConsumers(transactionRepo, bus)...)
	all = append(all, RegisterPaymentConsumers(paymentRepo, bus)...)

	for _, c := range all {
		go func(c ConsumerRegistration) {
			kafka.Consumer(ctx, c.Topic, c.GroupID, c.Handler)
		}(c)
	}
}

func WrapCommandHandler[T any](handler func(cmd T) error) HandlerFunc {
	return func(_ context.Context, _ string, value []byte) error {
		var cmd T
		if err := json.Unmarshal(value, &cmd); err != nil {
			return fmt.Errorf("unmarshal error: %w", err)
		}
		return handler(cmd)
	}
}
