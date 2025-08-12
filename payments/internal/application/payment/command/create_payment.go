package command

import (
	"context"
	"errors"

	"github.com/janapc/event-tickets/payments/internal/domain"
	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/internal/infra/logger"
)

type CreatePaymentCommand struct {
	UserName         string  `json:"user_name"`
	UserEmail        string  `json:"user_email"`
	EventId          string  `json:"event_id"`
	EventAmount      float64 `json:"event_amount"`
	PaymentToken     string  `json:"payment_token"`
	EventName        string  `json:"event_name"`
	EventDescription string  `json:"event_description"`
	EventImageUrl    string  `json:"event_image_url"`
	UserLanguage     string  `json:"user_language"`
}

type CreatePaymentHandler struct {
	PaymentRepo payment.IPaymentRepository
	Bus         domain.IEventBus
}

func NewCreatePaymentHandler(repo payment.IPaymentRepository, bus domain.IEventBus) *CreatePaymentHandler {
	return &CreatePaymentHandler{
		PaymentRepo: repo,
		Bus:         bus,
	}
}

func (h *CreatePaymentHandler) Handle(ctx context.Context, cmd CreatePaymentCommand) error {
	if cmd.EventAmount < 0 {
		return errors.New("amount must be greater than 0")
	}
	newPayment := payment.NewPayment(cmd.UserEmail, payment.StatusPending, cmd.EventId, cmd.EventAmount, cmd.PaymentToken)
	err := h.PaymentRepo.Save(ctx, newPayment)
	if err != nil {
		return err
	}
	h.Bus.Publish(payment.NewCreatedEvent(payment.CreatedEventPayload{
		UserName:         cmd.UserName,
		UserEmail:        cmd.UserEmail,
		EventId:          cmd.EventId,
		EventAmount:      cmd.EventAmount,
		PaymentToken:     cmd.PaymentToken,
		EventName:        cmd.EventName,
		EventDescription: cmd.EventDescription,
		EventImageUrl:    cmd.EventImageUrl,
		UserLanguage:     cmd.UserLanguage,
		PaymentID:        newPayment.ID,
	}, ctx))
	logger.Logger.WithContext(ctx).Infof("Payment %s created successfully", newPayment.ID)
	return nil
}
