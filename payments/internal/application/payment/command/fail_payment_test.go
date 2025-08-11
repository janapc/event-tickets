package command

import (
	"context"
	"errors"
	"testing"

	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/internal/infra/logger"
	"github.com/janapc/event-tickets/payments/internal/mock"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestFailPayment_Success(t *testing.T) {
	logger.Init()
	mockRepo := new(mock.PaymentRepositoryMock)
	mockBus := new(mock.EventBusMock)

	p := &payment.Payment{
		ID:     "pay_1",
		Status: payment.StatusPending,
	}
	mockRepo.On("FindByID", testMock.Anything, "pay_1").Return(p, nil)
	mockRepo.On("Update", testMock.Anything, p).Return(nil)
	mockBus.On("Publish", testMock.AnythingOfType("*payment.FailedEvent"))

	handler := NewFailPaymentHandler(mockRepo, mockBus)
	cmd := FailPaymentCommand{
		PaymentID:    "pay_1",
		UserName:     "Alice",
		UserEmail:    "alice@example.com",
		UserLanguage: "en",
	}
	ctx := context.Background()
	err := handler.Handle(ctx, cmd)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "FindByID", ctx, "pay_1")
	mockRepo.AssertCalled(t, "Update", ctx, p)
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*payment.FailedEvent"))
}

func TestFailPayment_FindByID_Error(t *testing.T) {
	logger.Init()
	mockRepo := new(mock.PaymentRepositoryMock)
	mockBus := new(mock.EventBusMock)
	mockRepo.On("FindByID", testMock.Anything, "pay_2").Return(&payment.Payment{}, errors.New("not found"))

	handler := NewFailPaymentHandler(mockRepo, mockBus)
	cmd := FailPaymentCommand{
		PaymentID:    "pay_2",
		UserName:     "Bob",
		UserEmail:    "bob@example.com",
		UserLanguage: "es",
	}

	ctx := context.Background()
	err := handler.Handle(ctx, cmd)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertCalled(t, "FindByID", ctx, "pay_2")
	mockRepo.AssertNotCalled(t, "Update", ctx, testMock.Anything)
	mockBus.AssertNotCalled(t, "Publish", testMock.Anything)
}
