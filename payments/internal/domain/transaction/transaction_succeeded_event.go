package transaction

import (
	"context"
	"encoding/json"
)

const SucceededEventName = "TRANSACTION_SUCCEEDED"

type SucceededEventPayload struct {
	PaymentID        string `json:"payment_id"`
	UserName         string `json:"user_name"`
	UserEmail        string `json:"user_email"`
	EventId          string `json:"event_id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_description"`
	EventImageUrl    string `json:"event_image_url"`
	UserLanguage     string `json:"user_language"`
}

type SucceededEvent struct {
	Payload SucceededEventPayload
	Context context.Context `json:"-"`
}

func NewSucceededEvent(payload SucceededEventPayload,
	ctx context.Context) *SucceededEvent {
	return &SucceededEvent{
		Payload: payload,
		Context: ctx,
	}
}

func (e SucceededEvent) ToMessage() ([]byte, error) {
	message, err := json.Marshal(e.Payload)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (e SucceededEvent) Name() string {
	return SucceededEventName
}
