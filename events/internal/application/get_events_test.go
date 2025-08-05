package application

import (
	"context"
	"testing"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/janapc/event-tickets/events/internal/mock"
	"github.com/janapc/event-tickets/events/pkg/pagination"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestGetEvents(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockEvent := []domain.Event{
		{ID: 1,
			Name:        "Test Event",
			Description: "Test Description",
			ImageUrl:    "http://test.com/image.jpg",
			Price:       99.99,
			EventDate:   time.Now().UTC(),
			Currency:    "USD",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now()},
	}
	mockPagination := pagination.NewPagination(1, 10, 10)
	mockRepo.On("List", testMock.Anything, testMock.AnythingOfType("int"), testMock.AnythingOfType("int")).Return(mockEvent, mockPagination, nil)
	getEvents := NewGetEvents(mockRepo)
	ctx := context.Background()
	events, err := getEvents.Execute(ctx, 1, 10)
	assert.NoError(t, err)
	assert.Len(t, events.Events, 1)

	event := events.Events[0]
	assert.Equal(t, int64(1), event.ID)
	assert.Equal(t, "Test Event", event.Name)
	assert.Equal(t, "Test Description", event.Description)
	assert.Equal(t, "http://test.com/image.jpg", event.ImageUrl)
	assert.Equal(t, 99.99, event.Price)
	assert.Equal(t, "USD", event.Currency)
	assert.Equal(t, 1, events.Pagination.Page)
	assert.Equal(t, 10, events.Pagination.Size)
	mockRepo.AssertExpectations(t)
	mockRepo.AssertNumberOfCalls(t, "List", 1)
	mockRepo.AssertCalled(t, "List", ctx, 1, 10)
}

func TestReturnErrorIfRepositoryListCallFails(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockRepo.On("List", testMock.Anything, testMock.AnythingOfType("int"), testMock.AnythingOfType("int")).Return([]domain.Event{}, pagination.Pagination{}, assert.AnError)
	getEvents := NewGetEvents(mockRepo)
	ctx := context.Background()
	events, err := getEvents.Execute(ctx, 1, 10)
	assert.Error(t, err)
	assert.Empty(t, events.Pagination)
	assert.Len(t, events.Events, 0)
}

func TestListEmpty(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockRepo.On("List", testMock.Anything, testMock.AnythingOfType("int"), testMock.AnythingOfType("int")).Return([]domain.Event{}, pagination.Pagination{}, nil)
	getEvents := NewGetEvents(mockRepo)
	ctx := context.Background()
	events, err := getEvents.Execute(ctx, 1, 10)
	assert.Empty(t, err)
	assert.Empty(t, events.Pagination)
	assert.Len(t, events.Events, 0)
	mockRepo.AssertCalled(t, "List", ctx, 1, 10)
}

func TestInvalidPaginationInputs(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockEvent := []domain.Event{
		{ID: 1,
			Name:        "Test Event",
			Description: "Test Description",
			ImageUrl:    "http://test.com/image.jpg",
			Price:       99.99,
			EventDate:   time.Now().UTC(),
			Currency:    "USD",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now()},
	}
	mockPagination := pagination.NewPagination(1, 10, 1)
	mockRepo.On("List", testMock.Anything, 1, 10).Return(mockEvent, mockPagination, nil)
	getEvents := NewGetEvents(mockRepo)
	ctx := context.Background()
	events, err := getEvents.Execute(ctx, -1, -5)
	assert.NoError(t, err)
	assert.NotEmpty(t, events.Events)
	assert.Equal(t, 1, events.Pagination.Page)
	assert.Equal(t, 10, events.Pagination.Size)
	mockRepo.AssertCalled(t, "List", ctx, 1, 10)
}

func TestLimitPageSize(t *testing.T) {
	mockRepo := new(mock.EventRepositoryMock)
	mockEvent := []domain.Event{
		{ID: 1,
			Name:        "Test Event",
			Description: "Test Description",
			ImageUrl:    "http://test.com/image.jpg",
			Price:       99.99,
			EventDate:   time.Now().UTC(),
			Currency:    "USD",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now()},
	}
	mockPagination := pagination.NewPagination(1, 100, 1)
	mockRepo.On("List", testMock.Anything, 1, 100).Return(mockEvent, mockPagination, nil)
	getEvents := NewGetEvents(mockRepo)
	ctx := context.Background()
	events, err := getEvents.Execute(ctx, 1, 200)
	assert.NoError(t, err)
	assert.NotEmpty(t, events.Events)
	assert.Equal(t, 100, events.Pagination.Size)
	mockRepo.AssertCalled(t, "List", ctx, 1, 100)
}
