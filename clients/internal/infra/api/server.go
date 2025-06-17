package api

import (
	"fmt"
	"os"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/janapc/event-tickets/clients/internal/application"
	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/janapc/event-tickets/clients/internal/infra/api/middleware"
	_ "github.com/janapc/event-tickets/clients/internal/infra/docs"
	"github.com/janapc/event-tickets/clients/internal/infra/logger"
)

type Server struct {
	Repository domain.IClientRepository
	Bus        domain.Bus
}

func NewServer(repo domain.IClientRepository, bus domain.Bus) *Server {
	return &Server{
		Repository: repo,
		Bus:        bus,
	}
}

func (s *Server) Init(port string) {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			logger.Logger.WithContext(ctx.UserContext()).Errorf("Error occurred: %v", err)
			message, statusCode := HandlerErrors(err)
			return ctx.Status(statusCode).JSON(message)
		},
	})
	if os.Getenv("ENV") != "PROD" {
		app.Use(fiberLogger.New())
	}
	if os.Getenv("ENV") == "PROD" {
		app.Use(otelfiber.Middleware())
		app.Use(middleware.OtelMetricMiddleware())
	}
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessEndpoint:  "/clients/healthcheck/live",
		ReadinessEndpoint: "/clients/healthcheck/ready",
	}))
	app.Get("/clients", s.HandlerGetClientByEmail)
	app.Post("/clients", s.HandlerSaveClient)
	app.Get("/clients/docs/*", swagger.HandlerDefault)
	addr := fmt.Sprintf(":%s", port)
	if err := app.Listen(addr); err != nil {
		logger.Logger.Errorf("Failed to start server error: %v", err)
	}
	logger.Logger.Infof("Server is running address %s", addr)
}

// GetClientByEmail godoc
// @Description get a client by email
// @Accept json
// @Produce json
// @Param email query string true "email" Format(email)
// @Success 200 {object} application.OutputGetClientByEmail
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /clients [get]
func (s *Server) HandlerGetClientByEmail(c *fiber.Ctx) error {
	ctx := c.UserContext()
	email := c.Query("email")
	getClientByEmail := application.NewGetClientByEmail(s.Repository)
	client, err := getClientByEmail.Execute(ctx, email)
	if err != nil {
		return err
	}
	return c.JSON(client)
}

// SaveClient godoc
// @Description register a new client
// @Accept json
// @Produce json
// @Param request body application.InputSaveClient true "client request"
// @Success 201 {object} application.OutputSaveClient
// @Failure 400
// @Failure 409
// @Failure 500
// @Router /clients [post]
func (s *Server) HandlerSaveClient(c *fiber.Ctx) error {
	input := new(application.InputSaveClient)
	if err := c.BodyParser(input); err != nil {
		return err
	}
	ctx := c.UserContext()
	saveClient := application.NewSaveClient(s.Repository, s.Bus)
	client, err := saveClient.Execute(ctx, *input)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(client)
}
