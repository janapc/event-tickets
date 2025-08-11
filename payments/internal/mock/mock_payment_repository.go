package mock

import (
	"context"

	domainPayment "github.com/janapc/event-tickets/payments/internal/domain/payment"

	testMock "github.com/stretchr/testify/mock"
)

type PaymentRepositoryMock struct {
	testMock.Mock
}

func (m *PaymentRepositoryMock) FindByID(ctx context.Context, ID string) (*domainPayment.Payment, error) {
	args := m.Called(ctx, ID)
	return args.Get(0).(*domainPayment.Payment), args.Error(1)
}

func (m *PaymentRepositoryMock) Save(ctx context.Context, payment *domainPayment.Payment) error {
	args := m.Called(ctx, payment)
	return args.Error(0)
}

func (m *PaymentRepositoryMock) Update(ctx context.Context, payment *domainPayment.Payment) error {
	args := m.Called(ctx, payment)
	return args.Error(0)
}
