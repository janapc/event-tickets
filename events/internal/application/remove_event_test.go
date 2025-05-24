package application

import (
	"context"
	"testing"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/janapc/event-tickets/events/internal/mock"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestRemoveEvent(t *testing.T) {
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
	mockRepo.On("Remove", testMock.Anything, testMock.AnythingOfType("int64")).Return(nil)
	removeEvent := NewRemoveEvent(mockRepo)
	id := int64(1)
	ctx := context.Background()
	err := removeEvent.Execute(ctx, id)
	assert.NoError(t, err)
	mockRepo.AssertNumberOfCalls(t, "Remove", 1)
	mockRepo.AssertNumberOfCalls(t, "FindByID", 1)
	mockRepo.AssertCalled(t, "Remove", ctx, id)
	mockRepo.AssertCalled(t, "FindByID", ctx, id)

}

func TestReturnErrorIfEventIsNotFound(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockRepo.On("FindByID", testMock.Anything, testMock.AnythingOfType("int64")).Return(&domain.Event{}, assert.AnError)
	removeEvent := NewRemoveEvent(mockRepo)
	ctx := context.Background()
	err := removeEvent.Execute(ctx, int64(1))
	if assert.Error(t, err) {
		assert.Equal(t, err.Error(), "event is not found")
	}
}
