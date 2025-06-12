package mocks

import (
	"github.com/janapc/event-tickets/clients/internal/domain"
	testMock "github.com/stretchr/testify/mock"
)

type MockClientRepository struct {
	testMock.Mock
}

func (m *MockClientRepository) Save(client *domain.Client) (*domain.Client, error) {
	args := m.Called(client)
	return args.Get(0).(*domain.Client), args.Error(1)
}

func (m *MockClientRepository) GetByEmail(email string) (*domain.Client, error) {
	args := m.Called(email)
	return args.Get(0).(*domain.Client), args.Error(1)
}
