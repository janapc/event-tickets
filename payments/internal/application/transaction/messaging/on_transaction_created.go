package messaging

import (
	"encoding/json"
	"github.com/janapc/event-tickets/payments/internal/application/transaction/command"
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/pkg/eventbus"
)

type OnTransactionCreated struct {
	Repo transaction.ITransactionRepository
	Bus  *eventbus.EventBus
}

func NewOnTransactionCreated(repo transaction.ITransactionRepository, bus *eventbus.EventBus) *OnTransactionCreated {
	return &OnTransactionCreated{
		Repo: repo,
		Bus:  bus,
	}
}

func (c *OnTransactionCreated) Handle(msg string) error {
	processTransactionHandler := command.NewProcessTransactionHandler(c.Repo, c.Bus)
	var input command.ProcessTransactionCommand
	err := json.Unmarshal([]byte(msg), &input)
	if err != nil {
		return err
	}
	err = processTransactionHandler.Handle(input)
	if err != nil {
		return err
	}
	return nil
}
