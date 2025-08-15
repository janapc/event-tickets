package mocks

import (
	"context"

	testMock "github.com/stretchr/testify/mock"
)

type MockMessaging struct {
	testMock.Mock
}

func (m *MockMessaging) Consumer(topic, groupID string, handler func(context.Context, string) error) {
	m.Called(topic, groupID, handler)
}

func (m *MockMessaging) Producer(message []byte, ctx context.Context, topic string, key string) error {
	args := m.Called(message, ctx, topic, key)
	return args.Error(0)
}
