package mock

import (
	"github.com/janapc/event-tickets/payments/internal/domain"
	testMock "github.com/stretchr/testify/mock"
)

type EventBusMock struct {
	testMock.Mock
}

func (m *EventBusMock) Publish(event domain.Event) {
	m.Called(event)
}

func (m *EventBusMock) Subscribe(eventType string, handler func(domain.Event)) {
	m.Called(eventType, handler)
}
