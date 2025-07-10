package messaging

import (
	"encoding/json"
	"github.com/janapc/event-tickets/payments/internal/application/transaction/command"
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/pkg/eventbus"
)

type OnPaymentCreated struct {
	Repo transaction.ITransactionRepository
	Bus  *eventbus.EventBus
}

func NewOnPaymentCreated(repo transaction.ITransactionRepository, bus *eventbus.EventBus) *OnPaymentCreated {
	return &OnPaymentCreated{
		Repo: repo,
		Bus:  bus,
	}
}

func (c *OnPaymentCreated) Handle(msg string) error {
	createTransaction := command.NewCreateTransactionHandler(c.Repo, c.Bus)
	var input command.CreateTransactionCommand
	err := json.Unmarshal([]byte(msg), &input)
	if err != nil {
		return err
	}
	err = createTransaction.Handle(input)
	if err != nil {
		return err
	}
	return nil
}
