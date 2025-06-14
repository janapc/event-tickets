package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/janapc/event-tickets/clients/internal/infra/logger"
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
		logger.Logger.WithContext(ctx).Error("Failed to open database connection")
		return fmt.Errorf("failed to open database: %w", err)
	}
	if err = DB.PingContext(ctx); err != nil {
		logger.Logger.WithContext(ctx).Error("Failed to connect to database")
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	logger.Logger.WithContext(ctx).Info("Successfully connected to SQLite database.")

	return err
}

func Close(ctx context.Context) {
	if DB != nil {
		err := DB.Close()
		if err != nil {
			logger.Logger.WithContext(ctx).Error("Failed to close database connection")
		} else {
			logger.Logger.WithContext(ctx).Info("Database connection closed.")
		}
	}
}
