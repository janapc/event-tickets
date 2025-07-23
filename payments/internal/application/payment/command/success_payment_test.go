package command

import (
	"errors"
	"testing"

	"github.com/janapc/event-tickets/payments/internal/domain/payment"
	"github.com/janapc/event-tickets/payments/internal/mock"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestSuccessPayment_Success(t *testing.T) {
	mockRepo := new(mock.PaymentRepositoryMock)
	mockBus := new(mock.EventBusMock)

	p := &payment.Payment{
		ID:     "pay_1",
		Status: payment.StatusPending,
	}
	mockRepo.On("FindByID", "pay_1").Return(p, nil)
	mockRepo.On("Update", p).Return(nil)
	mockBus.On("Publish", testMock.AnythingOfType("*payment.SucceededEvent"))

	handler := NewSuccessPaymentHandler(mockRepo, mockBus)
	cmd := SuccessPaymentCommand{
		PaymentID:        "pay_1",
		UserName:         "Alice",
		UserEmail:        "alice@example.com",
		EventId:          "event1",
		EventName:        "Concert",
		EventDescription: "A great concert",
		EventImageUrl:    "http://img.com/1.png",
		UserLanguage:     "en",
	}

	err := handler.Handle(cmd)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "FindByID", "pay_1")
	mockRepo.AssertCalled(t, "Update", p)
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*payment.SucceededEvent"))
}

func TestSuccessPayment_FindByID_Error(t *testing.T) {
	mockRepo := new(mock.PaymentRepositoryMock)
	mockBus := new(mock.EventBusMock)

	mockRepo.On("FindByID", "pay_2").Return(&payment.Payment{}, errors.New("not found"))

	handler := NewSuccessPaymentHandler(mockRepo, mockBus)
	cmd := SuccessPaymentCommand{
		PaymentID:        "pay_2",
		UserName:         "Bob",
		UserEmail:        "bob@example.com",
		EventId:          "event2",
		EventName:        "Festival",
		EventDescription: "A fun festival",
		EventImageUrl:    "http://img.com/2.png",
		UserLanguage:     "es",
	}

	err := handler.Handle(cmd)
	assert.EqualError(t, err, "not found")
	mockRepo.AssertCalled(t, "FindByID", "pay_2")
	mockRepo.AssertNotCalled(t, "Update", testMock.Anything)
	mockBus.AssertNotCalled(t, "Publish", testMock.Anything)
}
