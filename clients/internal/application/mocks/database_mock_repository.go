package mocks

import (
	"database/sql"

	"github.com/janapc/event-tickets/clients/internal/domain"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseMockRepository struct {
	DB *sql.DB
}

func NewDatabaseMockRepository() (*DatabaseMockRepository, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}
	_, err = db.Exec(`CREATE TABLE clients(
		id CHAR(36) NOT NULL,
		name TEXT(150) NOT NULL,
		email TEXT(150) NOT NULL,
		created_at TIMESTAMP NOT NULL, 
		PRIMARY KEY(id)
		)`)
	if err != nil {
		return nil, err
	}
	return &DatabaseMockRepository{
		DB: db,
	}, nil
}

func (i *DatabaseMockRepository) Save(client *domain.Client) error {
	stmt, err := i.DB.Prepare("INSERT INTO clients(id, name, email, created_at) VALUES($1,$2,$3,$4)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(client.ID, client.Name, client.Email, client.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (i *DatabaseMockRepository) GetByEmail(email string) (*domain.Client, error) {
	stmt, err := i.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE email = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var client domain.Client
	err = stmt.QueryRow(email).Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
