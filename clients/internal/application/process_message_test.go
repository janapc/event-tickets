package application

import (
	"testing"

	"github.com/janapc/event-tickets/clients/internal/application/mocks"
	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestProcessMessage(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	queue := mocks.NewQueueMock()
	processMessage := NewProcessMessage(repository)
	input := InputProcessMessage{
		Name:             "banana",
		Email:            "banana@banana.com",
		EventId:          "123",
		EventName:        "test",
		EventDescription: "description",
		EventImageUrl:    "http://image.png",
		Language:         "pt",
	}
	err := processMessage.Execute(input, queue)
	assert.NoError(t, err)
	assert.Len(t, queue.Messages, 2)
	client, err := repository.GetByEmail("banana@banana.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, client.ID)
}

func TestProcessMessageWithClientExists(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	c, _ := domain.NewClient("banana", "banana@banana.com")
	_ = repository.Save(c)
	queue := mocks.NewQueueMock()
	processMessage := NewProcessMessage(repository)
	input := InputProcessMessage{
		Name:             "banana",
		Email:            "banana@banana.com",
		EventId:          "123",
		EventName:        "test",
		EventDescription: "description",
		EventImageUrl:    "http://image.png",
		Language:         "pt",
	}
	err := processMessage.Execute(input, queue)
	assert.NoError(t, err)
	assert.Len(t, queue.Messages, 1)
	client, err := repository.GetByEmail("banana@banana.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, client.ID)
}
