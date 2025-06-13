package application

import (
	"context"
	"time"

	"github.com/janapc/event-tickets/clients/internal/domain"
)

type InputSaveClient struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OutputSaveClient struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type SaveClient struct {
	Repository domain.IClientRepository
}

func NewSaveClient(repo domain.IClientRepository) *SaveClient {
	return &SaveClient{
		Repository: repo,
	}
}

func (s *SaveClient) Execute(ctx context.Context, input InputSaveClient) (*OutputSaveClient, error) {
	client, err := domain.NewClient(domain.ClientParams{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}
	newClient, err := s.Repository.Save(ctx, client)
	if err != nil {
		return nil, err
	}
	return &OutputSaveClient{
		ID:        newClient.ID,
		Name:      newClient.Name,
		Email:     newClient.Email,
		CreatedAt: newClient.CreatedAt,
	}, nil
}
