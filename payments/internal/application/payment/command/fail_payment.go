package command

import (
	"context"

	"github.com/janapc/event-tickets/payments/internal/domain"
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/internal/infra/logger"
)

type FailPaymentCommand struct {
	PaymentID    string `json:"payment_id"`
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserLanguage string `json:"user_language"`
}

type FailPaymentHandler struct {
	PaymentRepo payment.IPaymentRepository
	Bus         domain.IEventBus
}

func NewFailPaymentHandler(repo payment.IPaymentRepository, bus domain.IEventBus) *FailPaymentHandler {
	return &FailPaymentHandler{
		PaymentRepo: repo,
		Bus:         bus,
	}
}

func (h *FailPaymentHandler) Handle(ctx context.Context, cmd FailPaymentCommand) error {
	p, err := h.PaymentRepo.FindByID(ctx, cmd.PaymentID)
	if err != nil {
		return err
	}
	p.MarkFailed()
	err = h.PaymentRepo.Update(ctx, p)
	if err != nil {
		return err
	}
	h.Bus.Publish(payment.NewFailedEvent(payment.FailedEventPayload{
		UserName:     cmd.UserName,
		UserLanguage: cmd.UserLanguage,
		UserEmail:    cmd.UserEmail,
	}, ctx))
	logger.Logger.WithContext(ctx).Infof("Payment %s marked as failed", p.ID)
	return nil
}
