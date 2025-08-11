package transaction

import "context"

type ITransactionRepository interface {
	Save(ctx context.Context, transaction *Transaction) error
	FindByID(ctx context.Context, ID string) (*Transaction, error)
	Update(ctx context.Context, transaction *Transaction) error
}
