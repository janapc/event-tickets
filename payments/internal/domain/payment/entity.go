package payment

import "github.com/google/uuid"

type Status string

const (
	StatusPending Status = "PENDING"
	StatusSuccess Status = "SUCCESS"
	StatusFailed  Status = "FAILED"
)

type Payment struct {
	ID           string
	UserEmail    string
	Status       Status
	EventId      string
	Amount       float64
	PaymentToken string
}

func NewPayment(userEmail string, status Status, eventId string, amount float64, paymentToken string) *Payment {
	return &Payment{
		ID:           uuid.NewString(),
		UserEmail:    userEmail,
		Status:       status,
		EventId:      eventId,
		Amount:       amount,
		PaymentToken: paymentToken,
	}
}

func (p *Payment) MarkSuccess() {
	p.Status = StatusSuccess
}

func (p *Payment) MarkFailed() {
	p.Status = StatusFailed
}
