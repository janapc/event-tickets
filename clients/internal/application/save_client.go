package application

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/janapc/event-tickets/clients/internal/domain/events"
	"github.com/janapc/event-tickets/clients/internal/infra/logger"
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
	Bus        domain.Bus
}

func NewSaveClient(repo domain.IClientRepository, bus domain.Bus) *SaveClient {
	return &SaveClient{
		Repository: repo,
		Bus:        bus,
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
	logger.Logger.WithContext(ctx).Infof("Client saved successfully client_id %s", newClient.ID)
	event := events.NewClientCreatedEvent(
		uuid.New().String(),
		newClient.Email,
		ctx,
	)
	s.Bus.Dispatch(event)
	return &OutputSaveClient{
		ID:        newClient.ID,
		Name:      newClient.Name,
		Email:     newClient.Email,
		CreatedAt: newClient.CreatedAt,
	}, nil
}
