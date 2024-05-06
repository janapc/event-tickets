package api

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/janapc/event-tickets/clients/internal/application"
	"github.com/janapc/event-tickets/clients/internal/domain"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

type Server struct {
	Repository domain.IClientRepository
}

func NewServer(repo domain.IClientRepository) *Server {
	return &Server{
		Repository: repo,
	}
}

func (s *Server) Init(port string) {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			message, statusCode := HandlerErrors(err)
			return ctx.Status(statusCode).JSON(message)
		},
	})
	app.Use(logger.New())
	app.Get("/clients", s.HandlerGetClientByEmail)
	app.Post("/clients", s.HandlerSaveClient)
	app.Get("/clients/docs/*", fiberSwagger.WrapHandler)
	addr := fmt.Sprintf(":%s", port)
	if err := app.Listen(addr); err != nil {
		log.Panicln(err)
	}
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
	email := c.Query("email")
	fmt.Println(email)
	getClientByEmail := application.NewGetClientByEmail(s.Repository)
	client, err := getClientByEmail.Execute(email)
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
	saveClient := application.NewSaveClient(s.Repository)
	client, err := saveClient.Execute(*input)
	if err != nil {
		return err
	}
	return c.Status(201).JSON(client)
}
