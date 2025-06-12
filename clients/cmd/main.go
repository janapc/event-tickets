package main

import (
	"log/slog"
	"os"

	"github.com/janapc/event-tickets/clients/internal/application"
	"github.com/janapc/event-tickets/clients/internal/infra/api"
	"github.com/janapc/event-tickets/clients/internal/infra/database"
	_ "github.com/janapc/event-tickets/clients/internal/infra/docs"
	"github.com/janapc/event-tickets/clients/internal/infra/messaging/kafka"
	"github.com/joho/godotenv"
)

func init() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}

	err = database.Init()
	if err != nil {
		panic(err)
	}
}

// @title Client API
// @version 1.0
// @description Api to manager clients

// @host localhost:3004
// @BasePath /
func main() {
	defer database.Close()
	port := os.Getenv("PORT")
	repository := database.NewClientRepository(database.DB)
	api := api.NewServer(repository)

	kakfaClient := kafka.NewKafkaClient([]string{os.Getenv("KAFKA_BROKER")})
	defer kakfaClient.WaitForShutdown()

	processMessage := application.NewProcessMessage(repository, kakfaClient, os.Getenv("CLIENT_CREATED_TOPIC"), os.Getenv("SEND_TICKET_TOPIC"))
	go kakfaClient.Consumer(os.Getenv("SUCCESS_PAYMENT_TOPIC"), "my-group", processMessage.Execute)
	api.Init(port)
}
