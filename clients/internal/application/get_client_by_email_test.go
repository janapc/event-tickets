package application

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/janapc/event-tickets/clients/internal/mocks"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestGetClientByEmail(t *testing.T) {
	mockRepo := new(mocks.MockClientRepository)
	mockClient := &domain.Client{
		ID:        "123",
		Name:      "new",
		Email:     "new@test.com",
		CreatedAt: time.Now(),
	}
	mockRepo.On("GetByEmail", testMock.Anything, testMock.AnythingOfType("string")).Return(mockClient, nil)
	getClient := NewGetClientByEmail(mockRepo)
	input := "new@test.com"
	ctx := context.Background()
	client, err := getClient.Execute(ctx, input)
	assert.NoError(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, client.ID, client.ID)
	assert.Equal(t, client.Name, client.Name)
	assert.Equal(t, client.Email, client.Email)
	mockRepo.AssertCalled(t, "GetByEmail", testMock.Anything, "new@test.com")
}

func TestErrorIfClientIsNotExists(t *testing.T) {
	mockRepo := new(mocks.MockClientRepository)
	expectedError := errors.New("client is not found")
	mockRepo.On("GetByEmail", testMock.Anything, testMock.AnythingOfType("string")).Return((*domain.Client)(nil), expectedError)
	getClient := NewGetClientByEmail(mockRepo)
	input := "new@test.com"
	ctx := context.Background()
	client, err := getClient.Execute(ctx, input)
	assert.Error(t, expectedError, err)
	assert.Empty(t, client)
	mockRepo.AssertCalled(t, "GetByEmail", testMock.Anything, "new@test.com")
}
