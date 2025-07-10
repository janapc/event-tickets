package database

import (
	"database/sql"
	domainTransaction "github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"time"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (t *TransactionRepository) Save(transaction *domainTransaction.Transaction) error {
	stmt, err := t.DB.Prepare(`INSERT INTO transactions(id, payment_id, status, reason) VALUES($1,$2, $3, $4)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(transaction.ID, transaction.PaymentID, transaction.Status, transaction.Reason)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepository) FindByID(ID string) (*domainTransaction.Transaction, error) {
	stmt, err := t.DB.Prepare(`SELECT id, payment_id, status, reason FROM transactions WHERE id=$1`)
	if err != nil {
		return &domainTransaction.Transaction{}, err
	}
	defer stmt.Close()
	var transaction domainTransaction.Transaction
	err = stmt.QueryRow(ID).Scan(&transaction.ID, &transaction.PaymentID, &transaction.Status, &transaction.Reason)
	if err != nil {
		return &domainTransaction.Transaction{}, err
	}
	return &transaction, nil
}

func (t *TransactionRepository) Update(transaction *domainTransaction.Transaction) error {
	stmt, err := t.DB.Prepare("UPDATE transactions SET status=$1, reason=$2, updated_at=$3 WHERE id=$4")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(transaction.Status, transaction.Reason, time.Now().UTC(), transaction.ID)
	if err != nil {
		return err
	}
	return nil
}
