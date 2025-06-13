package database

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"

	_ "github.com/lib/pq"
	"go.opentelemetry.io/otel"

	"github.com/XSAM/otelsql"
)

var DB *sql.DB

func Init(ctx context.Context) error {
	driverName, err := otelsql.Register("postgres",
		otelsql.WithAttributes(),
		otelsql.WithTracerProvider(otel.GetTracerProvider()),
	)
	if err != nil {
		return fmt.Errorf("failed to register otelsql driver: %w", err)
	}
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	DB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		slog.Error("Failed to open database connection")
		return fmt.Errorf("failed to open database: %w", err)
	}
	if err = DB.PingContext(ctx); err != nil {
		slog.Error("Failed to connect to database")
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	slog.Info("Successfully connected to SQLite database.")

	return err
}

func Close(ctx context.Context) {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			slog.Error("Failed to close database connection")
		} else {
			slog.Info("Database connection closed.")
		}
	}
}
