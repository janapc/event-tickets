package application

import (
	"testing"
	"time"

	"github.com/janapc/event-tickets/events/internal/application/mocks"
	"github.com/stretchr/testify/assert"
)

func TestShouldRegisterAEvent(t *testing.T) {
	DDMMYYYY := "02/01/2006"
	repository, _ := mocks.NewDatabaseMockRepository()
	registerEvent := NewRegisterEvent(repository)
	eventDate := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	input := InputRegisterEventDTO{
		Name:        "test",
		Description: "muito legal",
		ImageUrl:    "http://test.png",
		Price:       150.99,
		EventDate:   eventDate,
		Currency:    "BRL",
	}
	event, err := registerEvent.Execute(input)
	assert.NoError(t, err)
	assert.NotEmpty(t, event)
	assert.Equal(t, event.Name, input.Name)
}

func TestShouldErrorIfFieldsAreInvalid(t *testing.T) {
	DDMMYYYY := "02/01/2006"
	repository, _ := mocks.NewDatabaseMockRepository()
	registerEvent := NewRegisterEvent(repository)
	eventDate := time.Now().Add(-48 * time.Hour).Format(DDMMYYYY)
	input := InputRegisterEventDTO{
		Name:        "test",
		Description: "muito legal",
		ImageUrl:    "http://test.png",
		Price:       150.99,
		EventDate:   eventDate,
		Currency:    "BRL",
	}
	event, err := registerEvent.Execute(input)
	if assert.Error(t, err) {
		assert.Equal(t, err.Error(), "the event_date field cannot be less than current date")
	}
	assert.Empty(t, event)
}
