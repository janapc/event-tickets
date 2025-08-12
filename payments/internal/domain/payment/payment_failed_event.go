package payment

import (
	"context"
	"encoding/json"
)

const FailedEventName = "PAYMENT_FAILED"

type FailedEventPayload struct {
	UserName     string `json:"user_name"`
	UserEmail    string `json:"user_email"`
	UserLanguage string `json:"user_language"`
}

type FailedEvent struct {
	Payload FailedEventPayload
	Context context.Context `json:"-"`
}

func NewFailedEvent(payload FailedEventPayload,
	ctx context.Context) *FailedEvent {
	return &FailedEvent{
		Payload: payload,
		Context: ctx,
	}
}

func (e FailedEvent) ToMessage() ([]byte, error) {
	message, err := json.Marshal(e.Payload)
	if err != nil {
		return nil, err
	}
	return message, nil
}

func (e FailedEvent) Name() string {
	return FailedEventName
}
