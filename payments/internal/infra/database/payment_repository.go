package database

import (
	"database/sql"
	domainPayment "github.com/janapc/event-tickets/payments/internal/domain/payment"
	"time"
)

type PaymentRepository struct {
	DB *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepository {
	return &PaymentRepository{
		DB: db,
	}
}

func (r *PaymentRepository) FindByID(ID string) (*domainPayment.Payment, error) {
	stmt, err := r.DB.Prepare(`SELECT id, user_email, status, event_id, amount FROM payments WHERE id=$1`)
	if err != nil {
		return &domainPayment.Payment{}, err
	}
	defer stmt.Close()
	var payment domainPayment.Payment
	err = stmt.QueryRow(ID).Scan(&payment.ID, &payment.UserEmail, &payment.Status, &payment.EventId, &payment.Amount)
	if err != nil {
		return &domainPayment.Payment{}, err
	}
	return &payment, nil
}

func (r *PaymentRepository) Save(payment *domainPayment.Payment) error {
	stmt, err := r.DB.Prepare(`INSERT INTO payments(id, user_email, status, event_id, amount) VALUES($1,$2,$3,$4,$5) RETURNING id`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(payment.ID, payment.UserEmail, payment.Status, payment.EventId, payment.Amount)
	if err != nil {
		return err
	}
	return nil
}

func (r *PaymentRepository) Update(payment *domainPayment.Payment) error {
	stmt, err := r.DB.Prepare("UPDATE payments SET status=$1,  updated_at=$2 WHERE id=$3")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(payment.Status, time.Now().UTC(), payment.ID)
	if err != nil {
		return err
	}
	return nil
}
