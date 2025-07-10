package transaction

import "github.com/google/uuid"

type Status string

const (
	StatusPending Status = "PENDING"
	StatusSuccess Status = "SUCCESS"
	StatusFailed  Status = "FAILED"
)

type Transaction struct {
	ID        string
	Status    Status
	PaymentID string
	Reason    string
}

func NewTransaction(paymentId string) *Transaction {
	return &Transaction{
		ID:        uuid.NewString(),
		PaymentID: paymentId,
		Status:    StatusPending,
		Reason:    "-",
	}
}

func (t *Transaction) MarkSuccess() {
	t.Status = StatusSuccess
}

func (t *Transaction) MarkFailed(reason string) {
	t.Status = StatusFailed
	t.Reason = reason
}
