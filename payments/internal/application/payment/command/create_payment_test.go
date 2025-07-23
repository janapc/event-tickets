package command

import (
	"errors"
	"testing"

	"github.com/janapc/event-tickets/payments/internal/mock"
	"github.com/stretchr/testify/assert"
	testMock "github.com/stretchr/testify/mock"
)

func TestCreatePayment_Success(t *testing.T) {
	mockRepo := new(mock.PaymentRepositoryMock)
	mockBus := new(mock.EventBusMock)

	mockRepo.On("Save", testMock.AnythingOfType("*payment.Payment")).Return(nil)
	mockBus.On("Publish", testMock.AnythingOfType("*payment.CreatedEvent"))

	handler := NewCreatePaymentHandler(mockRepo, mockBus)
	cmd := CreatePaymentCommand{
		UserName:         "Alice",
		UserEmail:        "alice@example.com",
		EventId:          "event1",
		EventAmount:      100.0,
		PaymentToken:     "TOKEN123",
		EventName:        "Concert",
		EventDescription: "A great concert",
		EventImageUrl:    "http://img.com/1.png",
		UserLanguage:     "en",
	}

	err := handler.Handle(cmd)
	assert.NoError(t, err)
	mockRepo.AssertCalled(t, "Save", testMock.AnythingOfType("*payment.Payment"))
	mockBus.AssertCalled(t, "Publish", testMock.AnythingOfType("*payment.CreatedEvent"))
}

func TestCreatePayment_Failed_NegativeAmount(t *testing.T) {
	mockRepo := new(mock.PaymentRepositoryMock)
	mockBus := new(mock.EventBusMock)

	handler := NewCreatePaymentHandler(mockRepo, mockBus)
	cmd := CreatePaymentCommand{
		UserName:         "Bob",
		UserEmail:        "bob@example.com",
		EventId:          "event2",
		EventAmount:      -10.0,
		PaymentToken:     "TOKEN456",
		EventName:        "Festival",
		EventDescription: "A fun festival",
		EventImageUrl:    "http://img.com/2.png",
		UserLanguage:     "es",
	}

	err := handler.Handle(cmd)
	assert.EqualError(t, err, "amount must be greater than 0")
	mockRepo.AssertNotCalled(t, "Save", testMock.Anything)
	mockBus.AssertNotCalled(t, "Publish", testMock.Anything)
}

func TestCreatePayment_Failed_SaveError(t *testing.T) {
	mockRepo := new(mock.PaymentRepositoryMock)
	mockBus := new(mock.EventBusMock)

	mockRepo.On("Save", testMock.AnythingOfType("*payment.Payment")).Return(errors.New("db error"))

	handler := NewCreatePaymentHandler(mockRepo, mockBus)
	cmd := CreatePaymentCommand{
		UserName:         "Carol",
		UserEmail:        "carol@example.com",
		EventId:          "event3",
		EventAmount:      50.0,
		PaymentToken:     "TOKEN789",
		EventName:        "Seminar",
		EventDescription: "A learning seminar",
		EventImageUrl:    "http://img.com/3.png",
		UserLanguage:     "fr",
	}

	err := handler.Handle(cmd)
	assert.EqualError(t, err, "db error")
	mockRepo.AssertCalled(t, "Save", testMock.AnythingOfType("*payment.Payment"))
	mockBus.AssertNotCalled(t, "Publish", testMock.Anything)
}
