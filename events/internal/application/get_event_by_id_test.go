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

func TestGetEventById(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockEvent := &domain.Event{
		ID:          1,
		Name:        "Test Event",
		Description: "Test Description",
		ImageUrl:    "http://test.com/image.jpg",
		Price:       99.99,
		EventDate:   time.Now().UTC(),
		Currency:    "USD",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	mockRepo.On("FindByID", testMock.Anything, testMock.AnythingOfType("int64")).Return(mockEvent, nil)
	getEventById := NewGetEventById(mockRepo)
	id := int64(1)
	ctx := context.Background()
	result, err := getEventById.Execute(ctx, id)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
	assert.NotEmpty(t, result.ID)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertCalled(t, "FindByID", ctx, id)
}
func TestReturnErrorIfRepositoryFindByIDCallFails(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockRepo.On("FindByID", testMock.Anything, testMock.AnythingOfType("int64")).Return(&domain.Event{}, assert.AnError)
	getEventById := NewGetEventById(mockRepo)
	id := int64(1)
	ctx := context.Background()
	event, err := getEventById.Execute(ctx, id)
	if assert.Error(t, err) {
		assert.Equal(t, err.Error(), "event is not found")
	}
	assert.Empty(t, event)
	mockRepo.AssertCalled(t, "FindByID", ctx, id)
}
