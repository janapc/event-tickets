package mocks

import (
	"github.com/janapc/event-tickets/clients/internal/domain"
	testMock "github.com/stretchr/testify/mock"
)

type MockEventBus struct {
	testMock.Mock
}

func (m *MockEventBus) Dispatch(event domain.Event) {
	m.Called(event)
}

func (m *MockEventBus) Register(eventName string, handler func(domain.Event)) {
	m.Called(eventName, handler)
}
