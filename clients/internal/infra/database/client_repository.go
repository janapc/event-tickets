package database

import (
	"database/sql"

	"github.com/janapc/event-tickets/clients/internal/domain"
)

type ClientRepository struct {
	DB *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{
		DB: db,
	}
}

func (p *ClientRepository) Save(client *domain.Client) error {
	stmt, err := p.DB.Prepare("INSERT INTO clients(id, name, email, created_at) VALUES($1,$2,$3,$4)")
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

func (p *ClientRepository) GetByEmail(email string) (*domain.Client, error) {
	stmt, err := p.DB.Prepare("SELECT id, name, email, created_at FROM clients WHERE email = $1")
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
