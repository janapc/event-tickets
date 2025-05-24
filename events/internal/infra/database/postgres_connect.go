package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"go.opentelemetry.io/otel"

	"github.com/XSAM/otelsql"
)

func PostgresConnect() (*sql.DB, error) {
	driverName, err := otelsql.Register("postgres",
		otelsql.WithAttributes(),
		otelsql.WithTracerProvider(otel.GetTracerProvider()),
	)
	if err != nil {
		return nil, err
	}
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	database, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return database, err
}
