package command

import (
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/pkg/eventbus"
)

type FailPaymentCommand struct {
	PaymentID    string `json:"payment_id"`
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserLanguage string `json:"user_language"`
}

type FailPaymentHandler struct {
	PaymentRepo payment.IPaymentRepository
	Bus         *eventbus.EventBus
}

func NewFailPaymentHandler(repo payment.IPaymentRepository, bus *eventbus.EventBus) *FailPaymentHandler {
	return &FailPaymentHandler{
		PaymentRepo: repo,
		Bus:         bus,
	}
}

func (h *FailPaymentHandler) Handle(cmd FailPaymentCommand) error {
	p, err := h.PaymentRepo.FindByID(cmd.PaymentID)
	if err != nil {
		return err
	}
	p.MarkFailed()
	h.PaymentRepo.Update(p)
	h.Bus.Publish(&payment.FailedEvent{
		UserName:     cmd.UserName,
		UserLanguage: cmd.UserLanguage,
		UserEmail:    cmd.UserEmail,
	})
	return nil
}
