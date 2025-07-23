package mock

import (
	"github.com/janapc/event-tickets/payments/internal/domain/transaction"

	testMock "github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	testMock.Mock
}

func (m *TransactionRepositoryMock) FindByID(ID string) (*transaction.Transaction, error) {
	args := m.Called(ID)
	return args.Get(0).(*transaction.Transaction), args.Error(1)
}

func (m *TransactionRepositoryMock) Save(transaction *transaction.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func (m *TransactionRepositoryMock) Update(transaction *transaction.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}
