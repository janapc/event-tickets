package command

import (
	"context"

	"github.com/janapc/event-tickets/payments/internal/domain"
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/internal/infra/logger"
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
	Bus             domain.IEventBus
}

func NewProcessTransactionHandler(repo transaction.ITransactionRepository, bus domain.IEventBus) *ProcessTransactionHandler {
	return &ProcessTransactionHandler{
		TransactionRepo: repo,
		Bus:             bus,
	}
}

func (h *ProcessTransactionHandler) Handle(ctx context.Context, cmd ProcessTransactionCommand) error {
	tx, err := h.TransactionRepo.FindByID(ctx, cmd.TransactionID)
	if err != nil {
		return err
	}
	success, reason := simulateGateway(cmd.PaymentToken, cmd.Amount)

	if success {
		tx.MarkSuccess()
		logger.Logger.WithContext(ctx).Infof("Transaction %s processed successfully", tx.ID)
		h.Bus.Publish(&transaction.SucceededEvent{
			UserName:         cmd.UserName,
			UserEmail:        cmd.UserEmail,
			UserLanguage:     cmd.UserLanguage,
			PaymentID:        cmd.PaymentID,
			EventId:          cmd.EventId,
			EventName:        cmd.EventName,
			EventDescription: cmd.EventDescription,
			EventImageUrl:    cmd.EventImageUrl,
			Context:          ctx,
		})

	} else {
		tx.MarkFailed(reason)
		logger.Logger.WithContext(ctx).Infof("Transaction %s failed: %s", tx.ID, reason)
		h.Bus.Publish(&transaction.FailedEvent{
			UserName:     cmd.UserName,
			UserEmail:    cmd.UserEmail,
			UserLanguage: cmd.UserLanguage,
			PaymentID:    cmd.PaymentID,
			Context:      ctx,
		})
	}
	return h.TransactionRepo.Update(ctx, tx)
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
