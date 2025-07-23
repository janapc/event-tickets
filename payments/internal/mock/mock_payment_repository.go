package mock

import (
	domainPayment "github.com/janapc/event-tickets/payments/internal/domain/payment"

	testMock "github.com/stretchr/testify/mock"
)

type PaymentRepositoryMock struct {
	testMock.Mock
}

func (m *PaymentRepositoryMock) FindByID(ID string) (*domainPayment.Payment, error) {
	args := m.Called(ID)
	return args.Get(0).(*domainPayment.Payment), args.Error(1)
}

func (m *PaymentRepositoryMock) Save(payment *domainPayment.Payment) error {
	args := m.Called(payment)
	return args.Error(0)
}

func (m *PaymentRepositoryMock) Update(payment *domainPayment.Payment) error {
	args := m.Called(payment)
	return args.Error(0)
}
