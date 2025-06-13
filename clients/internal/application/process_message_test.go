package application

import (
	"context"
	"testing"

	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/janapc/event-tickets/clients/internal/mocks"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestProcessMessage_Execute_NewClient(t *testing.T) {
	mockRepo := new(mocks.MockClientRepository)
	mockMessaging := new(mocks.MockMessaging)

	input := `{
  "messageId": "123",
  "name": "John Doe",
  "email": "john.doe@example.com",
  "eventId": "event123",
  "eventName": "Concert",
  "eventDescription": "A great concert",
  "eventImageUrl": "http://example.com/image.jpg",
  "language": "en"
 }`

	client := &domain.Client{
		ID:    "client123",
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	mockRepo.On("GetByEmail", testMock.Anything, testMock.AnythingOfType("string")).Return((*domain.Client)(nil), nil)
	mockRepo.On("Save", testMock.Anything, testMock.AnythingOfType("*domain.Client")).Return(client, nil)
	mockMessaging.On("Producer", testMock.Anything, testMock.Anything, testMock.Anything, testMock.Anything).Return(nil)

	processMessage := NewProcessMessage(mockRepo, mockMessaging, "clientCreatedQueue", "sendTicketQueue")
	ctx := context.Background()
	err := processMessage.Execute(ctx, input)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "GetByEmail", testMock.Anything, "john.doe@example.com")
	mockRepo.AssertCalled(t, "Save", testMock.Anything, testMock.Anything)
	mockMessaging.AssertNumberOfCalls(t, "Producer", 2)
}

func TestProcessMessage_Execute_ExistingClient(t *testing.T) {
	mockRepo := new(mocks.MockClientRepository)
	mockMessaging := new(mocks.MockMessaging)

	input := `{
  "messageId": "123",
  "name": "John Doe",
  "email": "john.doe@example.com",
  "eventId": "event123",
  "eventName": "Concert",
  "eventDescription": "A great concert",
  "eventImageUrl": "http://example.com/image.jpg",
  "language": "en"
 }`

	client := &domain.Client{
		ID:    "client123",
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	mockRepo.On("GetByEmail", testMock.Anything, "john.doe@example.com").Return(client, nil)
	mockMessaging.On("Producer", testMock.Anything, testMock.Anything, testMock.Anything, testMock.Anything).Return(nil)

	processMessage := NewProcessMessage(mockRepo, mockMessaging, "clientCreatedQueue", "sendTicketQueue")
	ctx := context.Background()

	err := processMessage.Execute(ctx, input)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "GetByEmail", testMock.Anything, "john.doe@example.com")
	mockRepo.AssertNotCalled(t, "Save", testMock.Anything, testMock.Anything)
	mockMessaging.AssertNumberOfCalls(t, "Producer", 1)
}

func TestProcessMessage_Execute_InvalidInput(t *testing.T) {
	mockRepo := new(mocks.MockClientRepository)
	mockMessaging := new(mocks.MockMessaging)

	input := `{
  "messageId": "123",
  "name": "John Doe",
  "email": "",
  "eventId": "event123",
  "eventName": "Concert",
  "eventDescription": "A great concert",
  "eventImageUrl": "http://example.com/image.jpg",
  "language": "en"
 }`

	processMessage := NewProcessMessage(mockRepo, mockMessaging, "clientCreatedQueue", "sendTicketQueue")
	ctx := context.Background()

	err := processMessage.Execute(ctx, input)

	assert.Error(t, err)
	assert.Equal(t, "email is required", err.Error())
	mockRepo.AssertNotCalled(t, "GetByEmail", testMock.Anything, testMock.Anything)
	mockRepo.AssertNotCalled(t, "Save", testMock.Anything, testMock.Anything)
	mockMessaging.AssertNotCalled(t, "Producer", testMock.Anything, testMock.Anything, testMock.Anything)
}
