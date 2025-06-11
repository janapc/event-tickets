package mock

import (
	"context"

	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/janapc/event-tickets/events/pkg/pagination"
	testMock "github.com/stretchr/testify/mock"
)

type EventRepositoryMock struct {
	testMock.Mock
}

func (m *EventRepositoryMock) Register(ctx context.Context, event *domain.Event) (*domain.Event, error) {
	args := m.Called(ctx, event)
	return args.Get(0).(*domain.Event), args.Error(1)
}

func (m *EventRepositoryMock) Update(ctx context.Context, event *domain.Event) error {
	args := m.Called(ctx, event)
	return args.Error(0)
}

func (m *EventRepositoryMock) Remove(ctx context.Context, id int64) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *EventRepositoryMock) List(ctx context.Context, page, size int) ([]domain.Event, pagination.Pagination, error) {
	args := m.Called(ctx, page, size)
	return args.Get(0).([]domain.Event), args.Get(1).(pagination.Pagination), args.Error(2)
}

func (m *EventRepositoryMock) FindByID(ctx context.Context, id int64) (*domain.Event, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Event), args.Error(1)
}
