package main

import (
	"context"
	"os"

	"github.com/janapc/event-tickets/events/internal/infra/api"
	"github.com/janapc/event-tickets/events/internal/infra/database"
	"github.com/janapc/event-tickets/events/internal/infra/telemetry"
	"github.com/joho/godotenv"

	_ "github.com/janapc/event-tickets/events/internal/infra/docs"
)

// @title Events API
// @version 1.0
// @description api to manager events

// @host localhost:3001/
// @BasePath events
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	ctx := context.Background()
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
	shutdown := telemetry.InitOpenTelemetry(ctx)
	defer shutdown()

	db, err := database.PostgresConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewEventRepository(db)
	server := api.NewApi(repository)
	port := os.Getenv("PORT")
	server.Init(port)
}
