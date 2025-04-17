package application

import (
	"testing"
	"time"

	"github.com/janapc/event-tickets/events/internal/application/mocks"
	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestShouldUpdateAEvent(t *testing.T) {
	DDMMYYYY := "02/01/2006"
	eventDate := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	repository, _ := mocks.NewDatabaseMockRepository()
	updateEvent := NewUpdateEvent(repository)
	event, _ := domain.NewEvent(
		"test",
		"test",
		"http://test.png",
		1.0,
		eventDate,
		"BRL",
	)
	event.CreatedAt = time.Now().Add(-48 * time.Hour)
	event.UpdatedAt = time.Now().Add(-48 * time.Hour)
	err := repository.Register(event)
	assert.NoError(t, err)
	input := InputUpdateEventDTO{
		ID:   event.ID,
		Name: "teste",
	}
	err = updateEvent.Execute(input)
	assert.NoError(t, err)
	p, _ := repository.List()
	assert.Equal(t, p[0].Name, input.Name)
	assert.Equal(t, p[0].Description, event.Description)
	assert.Equal(t, p[0].EventDate, event.EventDate)
	assert.NotEqual(t, p[0].UpdatedAt, event.CreatedAt)
}

func TestShouldErrorifItDoesNotFindTheEvent(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	updateEvent := NewUpdateEvent(repository)
	input := InputUpdateEventDTO{
		ID:   "2",
		Name: "teste",
	}
	err := updateEvent.Execute(input)
	if assert.Error(t, err) {
		assert.Equal(t, err.Error(), "event is not found")
	}
}
