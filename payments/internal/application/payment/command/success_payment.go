package command

import (
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/pkg/eventbus"
)

type SuccessPaymentCommand struct {
	PaymentID        string `json:"payment_id"`
	UserName         string `json:"user_name"`
	UserEmail        string `json:"user_email"`
	EventId          string `json:"event_id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_description"`
	EventImageUrl    string `json:"event_image_url"`
	UserLanguage     string `json:"user_language"`
}

type SuccessPaymentHandler struct {
	PaymentRepo payment.IPaymentRepository
	Bus         *eventbus.EventBus
}

func NewSuccessPaymentHandler(repo payment.IPaymentRepository, bus *eventbus.EventBus) *SuccessPaymentHandler {
	return &SuccessPaymentHandler{
		PaymentRepo: repo,
		Bus:         bus,
	}
}

func (h *SuccessPaymentHandler) Handle(cmd SuccessPaymentCommand) error {
	p, err := h.PaymentRepo.FindByID(cmd.PaymentID)
	if err != nil {
		return err
	}
	p.MarkSuccess()
	h.PaymentRepo.Update(p)
	h.Bus.Publish(&payment.SucceededEvent{
		UserName:         cmd.UserName,
		UserEmail:        cmd.UserEmail,
		UserLanguage:     cmd.UserLanguage,
		EventId:          cmd.EventId,
		EventName:        cmd.EventName,
		EventDescription: cmd.EventDescription,
		EventImageUrl:    cmd.EventImageUrl,
	})
	return nil
}
