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

func (p *ClientRepository) Save(client *domain.Client) (*domain.Client, error) {
	stmt, err := p.DB.Prepare("INSERT INTO clients(name, email) VALUES($1,$2) RETURNING *")
	if err != nil {
		return &domain.Client{}, err
	}
	defer stmt.Close()
	var newClient domain.Client
	err = stmt.QueryRow(client.Name, client.Email).Scan(&newClient.ID, &newClient.Name, &newClient.Email, &newClient.CreatedAt)
	if err != nil {
		return &domain.Client{}, err
	}

	return &newClient, nil

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
