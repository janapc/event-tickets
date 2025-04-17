package application

import (
	"testing"
	"time"

	"github.com/janapc/event-tickets/events/internal/application/mocks"
	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestShouldRemoveAEvent(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	err := repository.Register(&domain.Event{
		ID:          "1",
		Name:        "test",
		Description: "test",
		ImageUrl:    "http://test.png",
		Price:       1.0,
		EventDate:   time.Now().Add(1 * time.Hour),
		Currency:    "BRL",
	})
	assert.NoError(t, err)
	removeEvent := NewRemoveEvent(repository)
	err = removeEvent.Execute("1")
	assert.NoError(t, err)
	event, _ := repository.List()
	assert.Len(t, event, 0)
}

func TestShouldErrorIfEventIsNotFound(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	removeEvent := NewRemoveEvent(repository)
	err := removeEvent.Execute("1")
	if assert.Error(t, err) {
		assert.Equal(t, err.Error(), "event is not found")
	}
}
