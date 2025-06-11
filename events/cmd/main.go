package main

import (
	"context"
	"os"

	"github.com/janapc/event-tickets/events/internal/infra/api"
	"github.com/janapc/event-tickets/events/internal/infra/database"
	"github.com/janapc/event-tickets/events/internal/infra/logger"
	"github.com/janapc/event-tickets/events/internal/infra/telemetry"
	"github.com/joho/godotenv"

	_ "github.com/janapc/event-tickets/events/internal/infra/docs"
)

func init() {
	logger.Init()
	ctx := context.Background()
	if os.Getenv("ENV") != "PROD" {
		if err := godotenv.Load(); err != nil {
			logger.Logger.Panic(err)
		}
	}
	env := os.Getenv("ENV")
	if env == "PROD" {
		err := telemetry.Init(ctx)
		if err != nil {
			logger.Logger.Panic(err)
		}
	}
	err := database.Init(ctx)
	if err != nil {
		logger.Logger.Panic(err)
	}
}

// @title Events API
// @version 1.0
// @description api to manager events

// @host localhost:3001/
// @BasePath events
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	defer database.Close(context.Background())
	defer func() {
		if err := telemetry.Shutdown(context.Background()); err != nil {
			logger.Logger.WithError(err).Error("Error shutting down telemetry")
		}
	}()
	repository := database.NewEventRepository(database.DB)
	server := api.NewApi(repository)
	port := os.Getenv("PORT")
	server.Init(port)
}
