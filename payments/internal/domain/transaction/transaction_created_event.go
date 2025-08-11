package transaction

import "context"

const CreatedEventName = "TRANSACTION_CREATED"

type CreatedEvent struct {
	TransactionID    string          `json:"transaction_id"`
	UserName         string          `json:"user_name"`
	UserEmail        string          `json:"user_email"`
	EventId          string          `json:"event_id"`
	PaymentToken     string          `json:"payment_token"`
	EventName        string          `json:"event_name"`
	EventDescription string          `json:"event_description"`
	EventImageUrl    string          `json:"event_image_url"`
	UserLanguage     string          `json:"user_language"`
	PaymentID        string          `json:"payment_id"`
	Amount           float64         `json:"amount"`
	Context          context.Context `json:"-"`
}

func (e CreatedEvent) Name() string {
	return CreatedEventName
}
