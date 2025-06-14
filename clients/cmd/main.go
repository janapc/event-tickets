package main

import (
	"context"
	"os"

	"github.com/janapc/event-tickets/clients/internal/application"
	"github.com/janapc/event-tickets/clients/internal/infra/api"
	"github.com/janapc/event-tickets/clients/internal/infra/database"
	_ "github.com/janapc/event-tickets/clients/internal/infra/docs"
	"github.com/janapc/event-tickets/clients/internal/infra/logger"
	"github.com/janapc/event-tickets/clients/internal/infra/messaging/kafka"
	"github.com/janapc/event-tickets/clients/internal/infra/telemetry"
	"github.com/joho/godotenv"
)

func init() {
	logger.Init()
	if os.Getenv("ENV") != "PROD" {
		if err := godotenv.Load(); err != nil {
			logger.Logger.Panic(err)
		}
	}
	ctx := context.Background()
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

// @title Client API
// @version 1.0
// @description Api to manager clients

// @host localhost:3004
// @BasePath /
func main() {
	ctx := context.Background()
	defer database.Close(ctx)
	defer func() {
		if err := telemetry.Shutdown(ctx); err != nil {
			logger.Logger.WithContext(ctx).Errorf("Failed to shutdown telemetry error: %v", err)
		}
	}()
	port := os.Getenv("PORT")
	repository := database.NewClientRepository(database.DB)
	api := api.NewServer(repository)

	kakfaClient := kafka.NewKafkaClient([]string{os.Getenv("KAFKA_BROKERS")})
	defer kakfaClient.WaitForShutdown()

	processMessage := application.NewProcessMessage(repository, kakfaClient, os.Getenv("CLIENT_CREATED_TOPIC"), os.Getenv("SEND_TICKET_TOPIC"))
	go kakfaClient.Consumer(os.Getenv("SUCCESS_PAYMENT_TOPIC"), os.Getenv("KAFKA_GROUP_ID"), processMessage.Execute)
	api.Init(port)
}
