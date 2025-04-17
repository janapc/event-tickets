package application

import (
	"errors"
	"time"

	"github.com/janapc/event-tickets/clients/internal/domain"
)

type OutputGetClientByEmail struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type GetClientByEmail struct {
	Repository domain.IClientRepository
}

func NewGetClientByEmail(repo domain.IClientRepository) *GetClientByEmail {
	return &GetClientByEmail{
		Repository: repo,
	}
}

func (g *GetClientByEmail) Execute(email string) (*OutputGetClientByEmail, error) {
	client, err := g.Repository.GetByEmail(email)
	if err != nil {
		return nil, errors.New("client is not found")
	}
	return &OutputGetClientByEmail{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
	}, nil
}
