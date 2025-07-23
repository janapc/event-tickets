package command

import (
	"errors"
	"testing"

	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/internal/mock"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestFailPayment_Success(t *testing.T) {
	mockRepo := new(mock.PaymentRepositoryMock)
	mockBus := new(mock.EventBusMock)

	p := &payment.Payment{
		ID:     "pay_1",
		Status: payment.StatusPending,
	}
	mockRepo.On("FindByID", "pay_1").Return(p, nil)
	mockRepo.On("Update", p).Return(nil)
	mockBus.On("Publish", testMock.AnythingOfType("*payment.FailedEvent"))

	handler := NewFailPaymentHandler(mockRepo, mockBus)
	cmd := FailPaymentCommand{
		PaymentID:    "pay_1",
		UserName:     "Alice",
		UserEmail:    "alice@example.com",
		UserLanguage: "en",
	}

	err := handler.Handle(cmd)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "FindByID", "pay_1")
	mockRepo.AssertCalled(t, "Update", p)
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*payment.FailedEvent"))
}

func TestFailPayment_FindByID_Error(t *testing.T) {
	mockRepo := new(mock.PaymentRepositoryMock)

	mockBus := new(mock.EventBusMock)

	mockRepo.On("FindByID", "pay_2").Return(&payment.Payment{}, errors.New("not found"))

	handler := NewFailPaymentHandler(mockRepo, mockBus)
	cmd := FailPaymentCommand{
		PaymentID:    "pay_2",
		UserName:     "Bob",
		UserEmail:    "bob@example.com",
		UserLanguage: "es",
	}

	err := handler.Handle(cmd)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertCalled(t, "FindByID", "pay_2")
	mockRepo.AssertNotCalled(t, "Update", testMock.Anything)
	mockBus.AssertNotCalled(t, "Publish", testMock.Anything)
}
