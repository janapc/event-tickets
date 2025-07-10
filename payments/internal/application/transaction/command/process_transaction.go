package command

import (
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/pkg/eventbus"
	"log"
)

type ProcessTransactionCommand struct {
	TransactionID    string  `json:"transaction_id"`
	UserName         string  `json:"user_name"`
	UserEmail        string  `json:"user_email"`
	EventId          string  `json:"event_id"`
	PaymentToken     string  `json:"payment_token"`
	EventName        string  `json:"event_name"`
	EventDescription string  `json:"event_description"`
	EventImageUrl    string  `json:"event_image_url"`
	UserLanguage     string  `json:"user_language"`
	PaymentID        string  `json:"payment_id"`
	Amount           float64 `json:"amount"`
}

type ProcessTransactionHandler struct {
	TransactionRepo transaction.ITransactionRepository
	Bus             *eventbus.EventBus
}

func NewProcessTransactionHandler(repo transaction.ITransactionRepository, bus *eventbus.EventBus) *ProcessTransactionHandler {
	return &ProcessTransactionHandler{
		TransactionRepo: repo,
		Bus:             bus,
	}
}

func (h *ProcessTransactionHandler) Handle(cmd ProcessTransactionCommand) error {
	tx, err := h.TransactionRepo.FindByID(cmd.TransactionID)
	if err != nil {
		return err
	}
	success, reason := simulateGateway(cmd.PaymentToken, cmd.Amount)

	if success {
		tx.MarkSuccess()
		log.Printf("Transaction %s processed successfully", tx.ID)
		h.Bus.Publish(&transaction.SucceededEvent{
			UserName:         cmd.UserName,
			UserEmail:        cmd.UserEmail,
			UserLanguage:     cmd.UserLanguage,
			PaymentID:        cmd.PaymentID,
			EventId:          cmd.EventId,
			EventName:        cmd.EventName,
			EventDescription: cmd.EventDescription,
			EventImageUrl:    cmd.EventImageUrl,
		})

	} else {
		tx.MarkFailed(reason)
		log.Printf("Transaction %s failed: %s", tx.ID, reason)

		h.Bus.Publish(&transaction.FailedEvent{
			UserName:     cmd.UserName,
			UserEmail:    cmd.UserEmail,
			UserLanguage: cmd.UserLanguage,
			PaymentID:    cmd.PaymentID,
		})
	}
	return h.TransactionRepo.Update(tx)
}

func simulateGateway(token string, amount float64) (bool, string) {
	if token == "" || amount <= 0 {
		return false, "invalid payment data"
	}
	if len(token)%2 == 0 {
		return true, ""
	}
	return false, "transaction rejected by gateway"
}
