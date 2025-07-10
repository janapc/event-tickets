package messaging

import (
	"encoding/json"
	"github.com/janapc/event-tickets/payments/internal/application/payment/command"
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/pkg/eventbus"
)

type OnTransactionSucceeded struct {
	Repo payment.IPaymentRepository
	Bus  *eventbus.EventBus
}

func NewOnTransactionSucceeded(repo payment.IPaymentRepository, bus *eventbus.EventBus) *OnTransactionSucceeded {
	return &OnTransactionSucceeded{
		Repo: repo,
		Bus:  bus,
	}
}

func (c *OnTransactionSucceeded) Handle(msg string) error {
	successPaymentHandler := command.NewSuccessPaymentHandler(c.Repo, c.Bus)
	var input command.SuccessPaymentCommand
	err := json.Unmarshal([]byte(msg), &input)
	if err != nil {
		return err
	}
	err = successPaymentHandler.Handle(input)
	if err != nil {
		return err
	}
	return nil
}
