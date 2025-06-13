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

func (m *MockMessaging) Producer(topic string, key, value []byte, ctx context.Context) error {
	args := m.Called(topic, key, value, ctx)
	return args.Error(0)
}
