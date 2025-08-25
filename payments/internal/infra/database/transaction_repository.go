package database

import (
	"context"
	"database/sql"
	"time"

	domainTransaction "github.com/janapc/event-tickets/payments/internal/domain/transaction"
	"github.com/janapc/event-tickets/payments/internal/infra/logger"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{
		DB: db,
	}
}

func (t *TransactionRepository) Save(ctx context.Context, transaction *domainTransaction.Transaction) error {
	stmt, err := t.DB.PrepareContext(ctx, `INSERT INTO transactions(id, payment_id, status, reason) VALUES($1,$2, $3, $4)`)
	if err != nil {
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			logger.Logger.WithContext(ctx).Errorf("Error closing statement: %v", err)
		}
	}()
	_, err = stmt.ExecContext(ctx, transaction.ID, transaction.PaymentID, transaction.Status, transaction.Reason)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepository) FindByID(ctx context.Context, ID string) (*domainTransaction.Transaction, error) {
	stmt, err := t.DB.PrepareContext(ctx, `SELECT id, payment_id, status, reason FROM transactions WHERE id=$1`)
	if err != nil {
		return &domainTransaction.Transaction{}, err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			logger.Logger.WithContext(ctx).Errorf("Error closing statement: %v", err)
		}
	}()
	var transaction domainTransaction.Transaction
	err = stmt.QueryRowContext(ctx, ID).Scan(&transaction.ID, &transaction.PaymentID, &transaction.Status, &transaction.Reason)
	if err != nil {
		return &domainTransaction.Transaction{}, err
	}
	return &transaction, nil
}

func (t *TransactionRepository) Update(ctx context.Context, transaction *domainTransaction.Transaction) error {
	stmt, err := t.DB.PrepareContext(ctx, "UPDATE transactions SET status=$1, reason=$2, updated_at=$3 WHERE id=$4")
	if err != nil {
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			logger.Logger.WithContext(ctx).Errorf("Error closing statement: %v", err)
		}
	}()
	_, err = stmt.ExecContext(ctx, transaction.Status, transaction.Reason, time.Now().UTC(), transaction.ID)
	if err != nil {
		return err
	}
	return nil
}
