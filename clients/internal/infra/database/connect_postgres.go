package database

import (
	"database/sql"
	"fmt"
	"os"
)

func ConnectPostgres() *sql.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	return db
}
