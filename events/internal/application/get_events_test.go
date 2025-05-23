package application

import (
	"testing"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/janapc/event-tickets/events/internal/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetEvents(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockEvent := []*domain.Event{
		&domain.Event{ID: 1,
			Name:        "Test Event",
			Description: "Test Description",
			ImageUrl:    "http://test.com/image.jpg",
			Price:       99.99,
			EventDate:   time.Now().UTC(),
			Currency:    "USD",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now()},
	}
	mockRepo.On("List").Return(mockEvent, nil)
	getEvents := NewGetEvents(mockRepo)
	events, err := getEvents.Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.NotEmpty(t, events[0].ID)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "List", 1)
}

func TestReturnErrorIfRepositoryListCallFails(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockRepo.On("List").Return([]*domain.Event{}, assert.AnError)
	getEvents := NewGetEvents(mockRepo)
	events, err := getEvents.Execute()
	assert.Error(t, err)
	assert.Empty(t, events)
}

func TestListEmpty(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockRepo.On("List").Return([]*domain.Event{}, nil)
	getEvents := NewGetEvents(mockRepo)
	events, err := getEvents.Execute()
	assert.Empty(t, err)
	assert.Empty(t, events)
}
