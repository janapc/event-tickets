package command

import (
	"errors"
	"testing"

	"github.com/janapc/event-tickets/payments/internal/mock"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestCreateTransaction(t *testing.T) {
	mockRepo := new(mock.TransactionRepositoryMock)
	mockBus := new(mock.EventBusMock)
	mockRepo.On("Save", testMock.AnythingOfType("*transaction.Transaction")).Return(nil)
	mockBus.On("Publish", testMock.AnythingOfType("*transaction.CreatedEvent"))
	command := NewCreateTransactionHandler(mockRepo, mockBus)
	err := command.Handle(CreateTransactionCommand{
		UserName:         "test",
		UserEmail:        "test@test.com",
		UserLanguage:     "en",
		PaymentID:        "PAYMENT_ID",
		PaymentToken:     "PAYMENT_TOKEN",
		EventDescription: "EVENT_DESCRIPTION",
		EventImageUrl:    "EVENT_IMAGE_URL",
		EventId:          "EVENT_ID",
		EventAmount:      20.90,
		EventName:        "EVENT_NAME",
	})
	assert.NoError(t, err)
	mockRepo.AssertNumberOfCalls(t, "Save", 1)
	mockBus.AssertNumberOfCalls(t, "Publish", 1)
	mockRepo.AssertCalled(t, "Save", testMock.AnythingOfType("*transaction.Transaction"))
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*transaction.CreatedEvent"))
}

func TestCreateTransactionWithError(t *testing.T) {
	mockRepo := new(mock.TransactionRepositoryMock)
	mockBus := new(mock.EventBusMock)
	mockRepo.On("Save", testMock.AnythingOfType("*transaction.Transaction")).Return(errors.New("error"))
	command := NewCreateTransactionHandler(mockRepo, mockBus)
	err := command.Handle(CreateTransactionCommand{
		UserName:         "test",
		UserEmail:        "test@test.com",
		UserLanguage:     "en",
		PaymentID:        "PAYMENT_ID",
		PaymentToken:     "PAYMENT_TOKEN",
		EventDescription: "EVENT_DESCRIPTION",
		EventImageUrl:    "EVENT_IMAGE_URL",
		EventId:          "EVENT_ID",
		EventAmount:      20.90,
		EventName:        "EVENT_NAME",
	})
	assert.EqualError(t, err, "error")
	mockRepo.AssertNumberOfCalls(t, "Save", 1)
	mockBus.AssertNumberOfCalls(t, "Publish", 0)
}
