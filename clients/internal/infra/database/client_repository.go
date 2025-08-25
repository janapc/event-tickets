package database

import (
	"database/sql"

	"github.com/janapc/event-tickets/clients/internal/domain"
	"golang.org/x/net/context"
)

type ClientRepository struct {
	DB *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{
		DB: db,
	}
}

func (p *ClientRepository) Save(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	stmt, err := p.DB.PrepareContext(ctx, "INSERT INTO clients(name, email) VALUES($1,$2) RETURNING *")
	if err != nil {
		return &domain.Client{}, err
	}
	defer func() {
		_ = stmt.Close()
	}()
	var newClient domain.Client
	err = stmt.QueryRowContext(ctx, client.Name, client.Email).Scan(&newClient.ID, &newClient.Name, &newClient.Email, &newClient.CreatedAt)
	if err != nil {
		return &domain.Client{}, err
	}

	return &newClient, nil

}

func (p *ClientRepository) GetByEmail(ctx context.Context, email string) (*domain.Client, error) {
	stmt, err := p.DB.PrepareContext(ctx, "SELECT id, name, email, created_at FROM clients WHERE email = $1")
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = stmt.Close()
	}()
	var client domain.Client
	err = stmt.QueryRowContext(ctx, email).Scan(&client.ID, &client.Name, &client.Email, &client.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &client, nil
}
