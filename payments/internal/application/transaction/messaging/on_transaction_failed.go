package messaging

import (
	"encoding/json"
	"github.com/janapc/event-tickets/payments/internal/application/payment/command"
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/pkg/eventbus"
)

type OnTransactionFailed struct {
	Repo payment.IPaymentRepository
	Bus  *eventbus.EventBus
}

func NewOnTransactionFailed(repo payment.IPaymentRepository, bus *eventbus.EventBus) *OnTransactionFailed {
	return &OnTransactionFailed{
		Repo: repo,
		Bus:  bus,
	}
}

func (c *OnTransactionFailed) Handle(msg string) error {
	failPaymentHandler := command.NewFailPaymentHandler(c.Repo, c.Bus)
	var input command.FailPaymentCommand
	err := json.Unmarshal([]byte(msg), &input)
	if err != nil {
		return err
	}
	err = failPaymentHandler.Handle(input)
	if err != nil {
		return err
	}
	return nil
}
