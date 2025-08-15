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
	mockEventBus := new(mocks.MockEventBus)

	input := `{
  "user_name": "John Doe",
  "user_email": "john.doe@example.com",
  "event_id": "event123",
  "event_name":  "Concert",
  "event_description": "A great concert",
  "event_image_url": "http://example.com/image.jpg",
  "user_language": "en"
  }`

	client := &domain.Client{
		ID:    "client123",
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	mockRepo.On("GetByEmail", testMock.Anything, testMock.AnythingOfType("string")).Return((*domain.Client)(nil), nil)
	mockRepo.On("Save", testMock.Anything, testMock.AnythingOfType("*domain.Client")).Return(client, nil)
	mockEventBus.On("Dispatch", testMock.AnythingOfType("*events.ClientCreatedEvent")).Once()
	mockEventBus.On("Dispatch", testMock.AnythingOfType("*events.SendTicketEvent")).Once()

	processMessage := NewProcessMessage(mockRepo, mockMessaging, mockEventBus)
	ctx := context.Background()
	err := processMessage.Execute(ctx, input)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "GetByEmail", testMock.Anything, "john.doe@example.com")
	mockRepo.AssertCalled(t, "Save", testMock.Anything, testMock.Anything)
	mockEventBus.AssertNumberOfCalls(t, "Dispatch", 2)
}

func TestProcessMessage_Execute_ExistingClient(t *testing.T) {
	mockRepo := new(mocks.MockClientRepository)
	mockMessaging := new(mocks.MockMessaging)
	mockEventBus := new(mocks.MockEventBus)

	input := `{
	"user_name": "John Doe",
  "user_email": "john.doe@example.com",
  "event_id": "event123",
  "event_name":  "Concert",
  "event_description": "A great concert",
  "event_image_url": "http://example.com/image.jpg",
  "user_language": "en"
 }`

	client := &domain.Client{
		ID:    "client123",
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	mockRepo.On("GetByEmail", testMock.Anything, "john.doe@example.com").Return(client, nil)
	mockEventBus.On("Dispatch", testMock.AnythingOfType("*events.SendTicketEvent")).Once()

	processMessage := NewProcessMessage(mockRepo, mockMessaging, mockEventBus)
	ctx := context.Background()

	err := processMessage.Execute(ctx, input)

	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "GetByEmail", testMock.Anything, "john.doe@example.com")
	mockRepo.AssertNotCalled(t, "Save", testMock.Anything, testMock.Anything)
	mockEventBus.AssertNumberOfCalls(t, "Dispatch", 1)
}

func TestProcessMessage_Execute_InvalidInput(t *testing.T) {
	mockRepo := new(mocks.MockClientRepository)
	mockMessaging := new(mocks.MockMessaging)
	mockEventBus := new(mocks.MockEventBus)

	input := `{
	"user_name": "John Doe",
  "user_email": "",
  "event_id": "event123",
  "event_name":  "Concert",
  "event_description": "A great concert",
  "event_image_url": "http://example.com/image.jpg",
  "user_language": "en"
 }`

	processMessage := NewProcessMessage(mockRepo, mockMessaging, mockEventBus)
	ctx := context.Background()

	err := processMessage.Execute(ctx, input)

	assert.Error(t, err)
	assert.Equal(t, "email is required", err.Error())
	mockRepo.AssertNotCalled(t, "GetByEmail")
	mockRepo.AssertNotCalled(t, "Save")
	mockMessaging.AssertNotCalled(t, "Dispatch")
}
