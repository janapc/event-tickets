package http

import (
	"os"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/janapc/event-tickets/payments/internal/application/payment/command"
	"github.com/janapc/event-tickets/payments/internal/interfaces/http/middleware"
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
	if os.Getenv("ENV") != "PROD" {
		app.Use(fiberLogger.New())
	}
	if os.Getenv("ENV") == "PROD" {
		app.Use(otelfiber.Middleware())
		app.Use(middleware.OtelMetricMiddleware())
	}
	payments := app.Group("/payments")
	payments.Post("/", h.createPayment)
}

func (h *PaymentHandler) createPayment(c *fiber.Ctx) error {
	ctx := c.UserContext()
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
	err := h.CreatePayment.Handle(ctx, command.CreatePaymentCommand{
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
