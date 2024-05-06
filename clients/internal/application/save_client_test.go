package application

import (
	"testing"

	"github.com/janapc/event-tickets/clients/internal/application/mocks"
	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveAClient(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	saveClient := NewSaveClient(repository)
	input := InputSaveClient{
		Name:  "test",
		Email: "test@test.com",
	}
	client, err := saveClient.Execute(input)
	assert.NoError(t, err)
	assert.NotEmpty(t, client)
	assert.NotEmpty(t, client.ID)
}

func TestErrorIfClientAlreadyExists(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	client, _ := domain.NewClient("test", "test@test.com")
	_ = repository.Save(client)
	saveClient := NewSaveClient(repository)
	input := InputSaveClient{
		Name:  "test2",
		Email: "test@test.com",
	}
	c, err := saveClient.Execute(input)
	if assert.Error(t, err) {
		assert.Equal(t, "client already exists", err.Error())
	}
	assert.Empty(t, c)
}
