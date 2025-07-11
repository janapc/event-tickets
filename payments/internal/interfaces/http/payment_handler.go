package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/janapc/event-tickets/payments/internal/application/payment/command"
)

type PaymentHandler struct {
	CreatePayment *command.CreatePaymentHandler
}

func NewPaymentHandler(createPayment *command.CreatePaymentHandler) *PaymentHandler {
	return &PaymentHandler{
		CreatePayment: createPayment,
	}
}

func (h *PaymentHandler) RegisterRoutes(app *fiber.App) {
	payments := app.Group("/payments")
	payments.Post("/", h.createPayment)
}

func (h *PaymentHandler) createPayment(c *fiber.Ctx) error {
	var input struct {
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
	if err := c.BodyParser(&input); err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}
	err := h.CreatePayment.Handle(command.CreatePaymentCommand{
		UserName:         input.UserName,
		UserEmail:        input.UserEmail,
		EventId:          input.EventId,
		EventAmount:      input.EventAmount,
		PaymentToken:     input.PaymentToken,
		EventName:        input.EventName,
		EventDescription: input.EventDescription,
		EventImageUrl:    input.EventImageUrl,
		UserLanguage:     input.UserLanguage,
	})
	if err != nil {
		log.Error(err.Error())
		return c.Status(fiber.StatusBadGateway).JSON(
			fiber.Map{
				"error": "bad gateway",
			})
	}
	if err := c.SendStatus(fiber.StatusCreated); err != nil {
		log.Error(err.Error())
	}
	return nil
}
