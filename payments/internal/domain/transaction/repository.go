package transaction

type ITransactionRepository interface {
	Save(transaction *Transaction) error
	FindByID(ID string) (*Transaction, error)
	Update(transaction *Transaction) error
}
