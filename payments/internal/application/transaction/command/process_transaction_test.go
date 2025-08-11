package command

import (
	"context"
	"errors"
	"testing"

	"github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/internal/infra/logger"
	"github.com/janapc/event-tickets/payments/internal/mock"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestProcessTransaction_Success(t *testing.T) {
	logger.Init()
	mockRepo := new(mock.TransactionRepositoryMock)
	mockBus := new(mock.EventBusMock)

	tx := &transaction.Transaction{
		ID:        "tx123",
		Status:    transaction.StatusPending,
		PaymentID: "pay_1",
		Reason:    "-",
	}
	mockRepo.On("FindByID", testMock.Anything, "tx123").Return(tx, nil)
	mockRepo.On("Update", testMock.Anything, tx).Return(nil)
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
	ctx := context.Background()
	err := handler.Handle(ctx, cmd)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "FindByID", ctx, "tx123")
	mockRepo.AssertCalled(t, "Update", ctx, tx)
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*transaction.SucceededEvent"))
	mockBus.AssertNotCalled(t, "Publish", testMock.AnythingOfType("*transaction.FailedEvent"))
}

func TestProcessTransaction_Failed_InvalidPaymentData(t *testing.T) {
	logger.Init()
	mockRepo := new(mock.TransactionRepositoryMock)
	mockBus := new(mock.EventBusMock)

	tx := &transaction.Transaction{
		ID:        "tx123",
		Status:    transaction.StatusPending,
		PaymentID: "pay_1",
		Reason:    "-",
	}
	mockRepo.On("FindByID", testMock.Anything, "tx124").Return(tx, nil)
	mockRepo.On("Update", testMock.Anything, tx).Return(nil)
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

	ctx := context.Background()
	err := handler.Handle(ctx, cmd)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "FindByID", ctx, "tx124")
	mockRepo.AssertCalled(t, "Update", ctx, tx)
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*transaction.FailedEvent"))
	mockBus.AssertNotCalled(t, "Publish", testMock.AnythingOfType("*transaction.SucceededEvent"))
}

func TestProcessTransaction_Failed_GatewayReject(t *testing.T) {
	logger.Init()
	mockRepo := new(mock.TransactionRepositoryMock)
	mockBus := new(mock.EventBusMock)

	tx := &transaction.Transaction{
		ID:        "tx123",
		Status:    transaction.StatusPending,
		PaymentID: "pay_1",
		Reason:    "-",
	}
	mockRepo.On("FindByID", testMock.Anything, "tx125").Return(tx, nil)
	mockRepo.On("Update", testMock.Anything, tx).Return(nil)
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

	ctx := context.Background()
	err := handler.Handle(ctx, cmd)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "FindByID", ctx, "tx125")
	mockRepo.AssertCalled(t, "Update", ctx, tx)
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*transaction.FailedEvent"))
	mockBus.AssertNotCalled(t, "Publish", testMock.AnythingOfType("*transaction.SucceededEvent"))
}

func TestProcessTransaction_FindByID_Error(t *testing.T) {
	logger.Init()
	mockRepo := new(mock.TransactionRepositoryMock)
	mockBus := new(mock.EventBusMock)

	mockRepo.On("FindByID", testMock.Anything, "tx126").Return(&transaction.Transaction{}, errors.New("not found"))

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

	ctx := context.Background()
	err := handler.Handle(ctx, cmd)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertCalled(t, "FindByID", ctx, "tx126")
	mockRepo.AssertNotCalled(t, "Update", ctx, testMock.Anything)
	mockBus.AssertNotCalled(t, "Publish", testMock.Anything)
}
