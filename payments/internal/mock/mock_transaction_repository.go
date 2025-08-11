package mock

import (
	"context"

	"github.com/janapc/event-tickets/payments/internal/domain/transaction"

	testMock "github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	testMock.Mock
}

func (m *TransactionRepositoryMock) FindByID(ctx context.Context, ID string) (*transaction.Transaction, error) {
	args := m.Called(ctx, ID)
	return args.Get(0).(*transaction.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) Save(ctx context.Context, transaction *transaction.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) Update(ctx context.Context, transaction *transaction.Transaction) error {
	args := m.Called(ctx, transaction)
	return args.Error(0)
}
