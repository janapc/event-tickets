package mocks

import (
	"context"

	"github.com/janapc/event-tickets/clients/internal/domain"
	testMock "github.com/stretchr/testify/mock"
)

type MockMessaging struct {
	testMock.Mock
}

func (m *MockMessaging) Consumer(topic, groupID string, handler func(context.Context, string) error) {
	m.Called(topic, groupID, handler)
}

func (m *MockMessaging) Producer(params domain.ProducerParameters) error {
	args := m.Called(params)
	return args.Error(0)
}
