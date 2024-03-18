package application

import (
	"testing"

	"github.com/janapc/event-tickets/clients/internal/application/mocks"
	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestGetClientByEmail(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	c, _ := domain.NewClient("test", "test@test.com")
	repository.Save(c)
	getClientByEmail := NewGetClientByEmail(repository)
	client, err := getClientByEmail.Execute("test@test.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, client)
	assert.NotEmpty(t, client.ID)
}

func TestErrorIfClientIsNotExists(t *testing.T) {
	repository, _ := mocks.NewDatabaseMockRepository()
	getClientByEmail := NewGetClientByEmail(repository)
	client, err := getClientByEmail.Execute("test@test.com")
	if assert.Error(t, err) {
		assert.Equal(t, "client is not found", err.Error())
	}
	assert.Empty(t, client)
}
