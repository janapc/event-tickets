package application

import (
	"errors"
	"testing"
	"time"

	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/janapc/event-tickets/clients/internal/mocks"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestSaveClient(t *testing.T) {
	mockRepo := new(mocks.MockClientRepository)
	mockClient := &domain.Client{
		ID:        "123",
		Name:      "new",
		Email:     "new@test.com",
		CreatedAt: time.Now(),
	}
	mockRepo.On("Save", testMock.AnythingOfType("*domain.Client")).Return(mockClient, nil)
	saveClient := NewSaveClient(mockRepo)
	input := InputSaveClient{
		Name:  "new",
		Email: "new@test.com",
	}
	client, err := saveClient.Execute(input)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, mockClient.ID, client.ID)
	assert.Equal(t, mockClient.Name, client.Name)
	assert.Equal(t, mockClient.Email, client.Email)
	mockRepo.AssertCalled(t, "Save", testMock.AnythingOfType("*domain.Client"))
}

func TestErrorIfClientAlreadyExists(t *testing.T) {
	mockRepo := new(mocks.MockClientRepository)
	expectedError := errors.New("Error: pq: duplicate key value violates unique constraint \"clients_email_key\"")
	mockRepo.On("Save", testMock.AnythingOfType("*domain.Client")).Return(&domain.Client{}, expectedError)
	saveClient := NewSaveClient(mockRepo)
	input := InputSaveClient{
		Name:  "new",
		Email: "existing@test.com",
	}
	client, err := saveClient.Execute(input)
	assert.Error(t, expectedError, err)
	assert.Nil(t, client)
	mockRepo.AssertCalled(t, "Save", testMock.AnythingOfType("*domain.Client"))
}
