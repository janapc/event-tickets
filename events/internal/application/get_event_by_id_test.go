package application

import (
	"testing"
	"time"

	"github.com/janapc/event-tickets/events/internal/application/mocks"
	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetEventById(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	DDMMYYYY := "02/01/2006"
	expirateAt := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	event, _ := domain.NewEvent("test", "test", "http://test.png", 10.0, expirateAt)
	err := repository.Register(event)
	assert.NoError(t, err)
	getEventById := NewGetEventById(repository)
	result, err := getEventById.Execute(event.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}
func TestShouldErrorIfEventByIdIsNotExists(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	getEventById := NewGetEventById(repository)
	event, err := getEventById.Execute("1")
	if assert.Error(t, err) {
		assert.Equal(t, err.Error(), "event is not found")
	}
	assert.Empty(t, event)
}
