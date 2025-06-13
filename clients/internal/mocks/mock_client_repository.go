package mocks

import (
	"context"

	"github.com/janapc/event-tickets/clients/internal/domain"
	testMock "github.com/stretchr/testify/mock"
)

type MockClientRepository struct {
	testMock.Mock
}

func (m *MockClientRepository) Save(ctx context.Context, client *domain.Client) (*domain.Client, error) {
	args := m.Called(ctx, client)
	return args.Get(0).(*domain.Client), args.Error(1)
}

func (m *MockClientRepository) GetByEmail(ctx context.Context, email string) (*domain.Client, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*domain.Client), args.Error(1)
}
