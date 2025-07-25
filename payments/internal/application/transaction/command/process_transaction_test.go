package command

import (
	"errors"
	"testing"

	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/internal/mock"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestProcessTransaction_Success(t *testing.T) {
	mockRepo := new(mock.TransactionRepositoryMock)
	mockBus := new(mock.EventBusMock)

	tx := &transaction.Transaction{
		ID:        "tx123",
		Status:    transaction.StatusPending,
		PaymentID: "pay_1",
		Reason:    "-",
	}
	mockRepo.On("FindByID", "tx123").Return(tx, nil)
	mockRepo.On("Update", tx).Return(nil)
	mockBus.On("Publish", testMock.AnythingOfType("*transaction.SucceededEvent"))

	handler := NewProcessTransactionHandler(mockRepo, mockBus)
	cmd := ProcessTransactionCommand{
		TransactionID:    "tx123",
		UserName:         "Alice",
		UserEmail:        "alice@example.com",
		EventId:          "event1",
		PaymentToken:     "TOKEN123",
		EventName:        "Concert",
		EventDescription: "A great concert",
		EventImageUrl:    "http://img.com/1.png",
		UserLanguage:     "en",
		PaymentID:        "pay_1",
		Amount:           100.0,
	}

	err := handler.Handle(cmd)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "FindByID", "tx123")
	mockRepo.AssertCalled(t, "Update", tx)
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*transaction.SucceededEvent"))
	mockBus.AssertNotCalled(t, "Publish", testMock.AnythingOfType("*transaction.FailedEvent"))
}

func TestProcessTransaction_Failed_InvalidPaymentData(t *testing.T) {
	mockRepo := new(mock.TransactionRepositoryMock)
	mockBus := new(mock.EventBusMock)

	tx := &transaction.Transaction{
		ID:        "tx123",
		Status:    transaction.StatusPending,
		PaymentID: "pay_1",
		Reason:    "-",
	}
	mockRepo.On("FindByID", "tx124").Return(tx, nil)
	mockRepo.On("Update", tx).Return(nil)
	mockBus.On("Publish", testMock.AnythingOfType("*transaction.FailedEvent"))

	handler := NewProcessTransactionHandler(mockRepo, mockBus)
	cmd := ProcessTransactionCommand{
		TransactionID:    "tx124",
		UserName:         "Bob",
		UserEmail:        "bob@example.com",
		EventId:          "event2",
		PaymentToken:     "",
		EventName:        "Festival",
		EventDescription: "A fun festival",
		EventImageUrl:    "http://img.com/2.png",
		UserLanguage:     "es",
		PaymentID:        "pay_2",
		Amount:           0,
	}

	err := handler.Handle(cmd)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "FindByID", "tx124")
	mockRepo.AssertCalled(t, "Update", tx)
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*transaction.FailedEvent"))
	mockBus.AssertNotCalled(t, "Publish", testMock.AnythingOfType("*transaction.SucceededEvent"))
}

func TestProcessTransaction_Failed_GatewayReject(t *testing.T) {
	mockRepo := new(mock.TransactionRepositoryMock)
	mockBus := new(mock.EventBusMock)

	tx := &transaction.Transaction{
		ID:        "tx123",
		Status:    transaction.StatusPending,
		PaymentID: "pay_1",
		Reason:    "-",
	}
	mockRepo.On("FindByID", "tx125").Return(tx, nil)
	mockRepo.On("Update", tx).Return(nil)
	mockBus.On("Publish", testMock.AnythingOfType("*transaction.FailedEvent"))

	handler := NewProcessTransactionHandler(mockRepo, mockBus)
	cmd := ProcessTransactionCommand{
		TransactionID:    "tx125",
		UserName:         "Carol",
		UserEmail:        "carol@example.com",
		EventId:          "event3",
		PaymentToken:     "ODD12",
		EventName:        "Seminar",
		EventDescription: "A learning seminar",
		EventImageUrl:    "http://img.com/3.png",
		UserLanguage:     "fr",
		PaymentID:        "pay_3",
		Amount:           50.0,
	}

	err := handler.Handle(cmd)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "FindByID", "tx125")
	mockRepo.AssertCalled(t, "Update", tx)
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*transaction.FailedEvent"))
	mockBus.AssertNotCalled(t, "Publish", testMock.AnythingOfType("*transaction.SucceededEvent"))
}

func TestProcessTransaction_FindByID_Error(t *testing.T) {
	mockRepo := new(mock.TransactionRepositoryMock)
	mockBus := new(mock.EventBusMock)

	mockRepo.On("FindByID", "tx126").Return(&transaction.Transaction{}, errors.New("not found"))

	handler := NewProcessTransactionHandler(mockRepo, mockBus)
	cmd := ProcessTransactionCommand{
		TransactionID:    "tx126",
		UserName:         "Dave",
		UserEmail:        "dave@example.com",
		EventId:          "event4",
		PaymentToken:     "TOKEN5678",
		EventName:        "Talk",
		EventDescription: "A tech talk",
		EventImageUrl:    "http://img.com/4.png",
		UserLanguage:     "de",
		PaymentID:        "pay_4",
		Amount:           30.0,
	}

	err := handler.Handle(cmd)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertCalled(t, "FindByID", "tx126")
	mockRepo.AssertNotCalled(t, "Update", testMock.Anything)
	mockBus.AssertNotCalled(t, "Publish", testMock.Anything)
}
