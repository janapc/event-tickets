package application

import (
	"context"
	"testing"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/janapc/event-tickets/events/internal/infra/logger"
	"github.com/janapc/event-tickets/events/internal/mock"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestCreateEvent(t *testing.T) {
	logger.Init()
	mockRepo := new(mock.EventRepositoryMock)
	eventDate := time.Now().Add(24 * time.Hour)
	mockEvent := &domain.Event{
		ID:          1,
		Name:        "Test Event",
		Description: "Test Description",
		ImageUrl:    "http://test.com/image.jpg",
		Price:       99.99,
		EventDate:   eventDate,
		Currency:    "USD",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockRepo.On("Register", testMock.Anything, testMock.AnythingOfType("*domain.Event")).Return(mockEvent, nil)
	registerEvent := NewRegisterEvent(mockRepo)
	input := InputRegisterEventDTO{
		Name:        "Test Event",
		Description: "Test Description",
		ImageUrl:    "http://test.com/image.jpg",
		Price:       99.99,
		EventDate:   eventDate.Format(time.RFC3339),
		Currency:    "USD",
	}
	ctx := context.Background()
	event, err := registerEvent.Execute(ctx, input)
	assert.Nil(t, err)
	assert.NotNil(t, event)
	assert.Equal(t, input.Name, event.Name)
	assert.Equal(t, input.Description, event.Description)
	assert.Equal(t, input.ImageUrl, event.ImageUrl)
	assert.Equal(t, input.Price, event.Price)
	assert.Equal(t, input.Currency, event.Currency)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "Register", ctx, testMock.AnythingOfType("*domain.Event"))

}

func TestReturnErrorIfFieldsAreInvalid(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	registerEvent := NewRegisterEvent(mockRepo)
	eventDate := time.Now().Add(-48 * time.Hour)
	input := InputRegisterEventDTO{
		Name:        "test",
		Description: "muito legal",
		ImageUrl:    "http://test.png",
		Price:       150.99,
		EventDate:   eventDate.Format(time.RFC3339),
		Currency:    "BRL",
	}
	ctx := context.Background()
	event, err := registerEvent.Execute(ctx, input)
	if assert.Error(t, err) {
		assert.Equal(t, err.Error(), "the event_date field cannot be less than the current date")
	}
	assert.Empty(t, event)
}

func TestRegistrationFailure(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	eventDate := time.Now().Add(24 * time.Hour)
	mockRepo.On("Register", testMock.Anything, testMock.AnythingOfType("*domain.Event")).Return(&domain.Event{}, assert.AnError)
	registerEvent := NewRegisterEvent(mockRepo)
	input := InputRegisterEventDTO{
		Name:        "Test Event",
		Description: "Test Description",
		ImageUrl:    "http://test.com/image.jpg",
		Price:       99.99,
		EventDate:   eventDate.Format(time.RFC3339),
		Currency:    "USD",
	}
	ctx := context.Background()
	event, err := registerEvent.Execute(ctx, input)
	assert.Error(t, err)
	assert.Empty(t, event)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "Register", ctx, testMock.AnythingOfType("*domain.Event"))
}
