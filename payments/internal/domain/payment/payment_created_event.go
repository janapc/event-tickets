package payment

import (
	"context"
	"encoding/json"
)

var CreatedEventName = "PAYMENT_CREATED"

type CreatedEventPayload struct {
	UserName         string  `json:"user_name"`
	UserEmail        string  `json:"user_email"`
	EventId          string  `json:"event_id"`
	EventAmount      float64 `json:"event_amount"`
	PaymentToken     string  `json:"payment_token"`
	EventName        string  `json:"event_name"`
	EventDescription string  `json:"event_description"`
	EventImageUrl    string  `json:"event_image_url"`
	UserLanguage     string  `json:"user_language"`
	PaymentID        string  `json:"payment_id"`
}

type CreatedEvent struct {
	Payload CreatedEventPayload
	Context context.Context `json:"-"`
}

func NewCreatedEvent(payload CreatedEventPayload,
	ctx context.Context) *CreatedEvent {
	return &CreatedEvent{
		Payload: payload,
		Context: ctx,
	}
}

func (e CreatedEvent) Name() string {
	return CreatedEventName
}

func (e CreatedEvent) ToMessage() ([]byte, error) {
	message, err := json.Marshal(e.Payload)
	if err != nil {
		return nil, err
	}
	return message, nil
}
