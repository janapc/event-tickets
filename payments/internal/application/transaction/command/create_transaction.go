package command

import (
	"github.com/janapc/event-tickets/payments/internal/domain"
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
)

type CreateTransactionCommand struct {
	UserName         string  `json:"user_name"`
	UserEmail        string  `json:"user_email"`
	EventId          string  `json:"event_id"`
	EventAmount      float64 `json:"event_amount"`
	PaymentToken     string  `json:"payment_token"`
	EventName        string  `json:"event_name"`
	EventDescription string  `json:"event_description"`
	EventImageUrl    string  `json:"event_image_url"`
	UserLanguage     string  `json:"user_language"`
	PaymentID        string  `json:"payment_id"`
}

type CreateTransactionHandler struct {
	TransactionRepo transaction.ITransactionRepository
	Bus             domain.IEventBus
}

func NewCreateTransactionHandler(repo transaction.ITransactionRepository, bus domain.IEventBus) *CreateTransactionHandler {
	return &CreateTransactionHandler{
		TransactionRepo: repo,
		Bus:             bus,
	}
}

func (h *CreateTransactionHandler) Handle(cmd CreateTransactionCommand) error {
	newTransaction := transaction.NewTransaction(cmd.PaymentID)
	err := h.TransactionRepo.Save(newTransaction)
	if err != nil {
		return err
	}
	h.Bus.Publish(&transaction.CreatedEvent{
		UserName:         cmd.UserName,
		UserEmail:        cmd.UserEmail,
		UserLanguage:     cmd.UserLanguage,
		PaymentID:        cmd.PaymentID,
		PaymentToken:     cmd.PaymentToken,
		EventDescription: cmd.EventDescription,
		EventImageUrl:    cmd.EventImageUrl,
		EventName:        cmd.EventName,
		TransactionID:    newTransaction.ID,
		EventId:          cmd.EventId,
		Amount:           cmd.EventAmount,
	})
	return nil

}
