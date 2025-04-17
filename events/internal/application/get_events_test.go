package application

import (
	"testing"
	"time"

	"github.com/janapc/event-tickets/events/internal/application/mocks"
	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetEvents(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	err := repository.Register(&domain.Event{
		ID:          "1",
		Name:        "test",
		ImageUrl:    "http://test.png",
		Description: "test",
		Price:       1.0,
		Currency:    "BRL",
		EventDate:   time.Now().Add(1 * time.Hour),
	})
	assert.NoError(t, err)
	getEvents := NewGetEvents(repository)
	events, err := getEvents.Execute()
	assert.NoError(t, err)
	assert.Len(t, events, 1)
	assert.NotEmpty(t, events[0].ID)
}
