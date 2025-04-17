package main

import (
	"log/slog"
	"os"

	"github.com/janapc/event-tickets/clients/internal/application"
	"github.com/janapc/event-tickets/clients/internal/infra/api"
	"github.com/janapc/event-tickets/clients/internal/infra/database"
	_ "github.com/janapc/event-tickets/clients/internal/infra/docs"
	"github.com/janapc/event-tickets/clients/internal/infra/queue"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// @title Client API
// @version 1.0
// @description Api to manager clients

// @host localhost:3004
// @BasePath /
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")
	db := database.ConnectPostgres()
	defer db.Close()

	repository := database.NewClientRepository(db)
	api := api.NewServer(repository)

	connRabbitMQ, channelRabbitMQ := queue.ConnectRabbitMQ()
	defer connRabbitMQ.Close()
	defer channelRabbitMQ.Close()

	processMessage := application.NewProcessMessage(repository)
	queue := queue.NewQueue(channelRabbitMQ, processMessage, logger)

	queueName := os.Getenv("QUEUE_SUCCESS_PAYMENT")
	go queue.Consumer(queueName, 10)
	api.Init(port)
}
