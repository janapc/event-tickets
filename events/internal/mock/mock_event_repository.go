package mock

import (
	"github.com/janapc/event-tickets/events/internal/domain"
	testMock "github.com/stretchr/testify/mock"
)

type EventRepositoryMock struct {
	testMock.Mock
}

func (m *EventRepositoryMock) Register(event *domain.Event) (*domain.Event, error) {
	args := m.Called(event)
	return args.Get(0).(*domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) Update(event *domain.Event) error {
	args := m.Called(event)
	return args.Error(0)
}

func (m *EventRepositoryMock) Remove(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *EventRepositoryMock) List() ([]*domain.Event, error) {
	args := m.Called()
	return args.Get(0).([]*domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) FindByID(id string) (*domain.Event, error) {
	args := m.Called(id)
	return args.Get(0).(*domain.Event), args.Error(1)
}
