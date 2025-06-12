package database

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() error {
	var err error
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		return err
	}
	slog.Info("Successfully connected to PostgreSQL database.")
	return nil
}

func Close() {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			slog.With(err).Error("Failed to close database connection")
		} else {
			slog.Info("Database connection closed.")
		}
	}
}
