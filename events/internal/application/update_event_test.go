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

func TestUpdateEvent(t *testing.T) {
	logger.Init()
	mockRepo := new(mock.EventRepositoryMock)
	mockEvent := &domain.Event{
		ID:          1,
		Name:        "Test Event",
		Description: "Test Description",
		ImageUrl:    "http://test.com/image.jpg",
		Price:       99.99,
		EventDate:   time.Now().Add(48 * time.Hour),
		Currency:    "USD",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockRepo.On("FindByID", testMock.Anything, testMock.AnythingOfType("int64")).Return(mockEvent, nil)
	mockRepo.On("Update", testMock.Anything, testMock.AnythingOfType("*domain.Event")).Return(nil)
	updateEvent := NewUpdateEvent(mockRepo)
	input := InputUpdateEventDTO{
		Name: "teste",
	}
	id := int64(1)
	ctx := context.Background()
	err := updateEvent.Execute(ctx, id, input)
	assert.NoError(t, err)
	mockRepo.AssertNumberOfCalls(t, "FindByID", 1)
	mockRepo.AssertCalled(t, "FindByID", ctx, id)
	mockRepo.AssertNumberOfCalls(t, "Update", 1)
	mockRepo.AssertCalled(t, "Update", ctx, mockEvent)
}

func TestErrorIfEventNotFound(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockRepo.On("FindByID", testMock.Anything, testMock.AnythingOfType("int64")).Return(&domain.Event{}, assert.AnError)
	updateEvent := NewUpdateEvent(mockRepo)
	id := int64(1)
	input := InputUpdateEventDTO{
		Name: "teste",
	}
	ctx := context.Background()
	err := updateEvent.Execute(ctx, id, input)
	if assert.Error(t, err) {
		assert.Equal(t, err.Error(), "event is not found")
	}
}
