package mocks

import (
	testMock "github.com/stretchr/testify/mock"
)

type MockMessaging struct {
	testMock.Mock
}

func (m *MockMessaging) Consumer(topic, groupID string, handler func(string) error) {
	m.Called(topic, groupID, handler)
}

func (m *MockMessaging) Producer(topic string, key, value []byte) error {
	args := m.Called(topic, key, value)
	return args.Error(0)
}
