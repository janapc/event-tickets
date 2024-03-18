package application

import (
	"errors"
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

func (s *SaveClient) Execute(input InputSaveClient) (*OutputSaveClient, error) {
	client, err := domain.NewClient(input.Name, input.Email)
	if err != nil {
		return nil, err
	}
	if c, _ := s.Repository.GetByEmail(input.Email); c != nil {
		return nil, errors.New("client already exists")
	}
	err = s.Repository.Save(client)
	if err != nil {
		return nil, err
	}
	return &OutputSaveClient{
		ID:        client.ID,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
	}, nil
}
