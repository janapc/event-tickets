package main

import (
	"os"

	"github.com/janapc/event-tickets/events/internal/infra/api"
	"github.com/janapc/event-tickets/events/internal/infra/database"
	"github.com/joho/godotenv"

	_ "github.com/janapc/event-tickets/events/internal/infra/docs"
)

// @title Events API
// @version 1.0
// @description api to manager events

// @host localhost:3000/
// @BasePath events
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
	port := os.Getenv("PORT")

	db, err := database.PostgresConnect()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	repository := database.NewPostgresRepository(db)
	api := api.NewApi(repository)

	api.Init(port)
}
