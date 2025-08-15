package events

import (
	"context"
	"encoding/json"
)

const CLIENT_CREATED_EVENT = "CLIENT_CREATED"

type ClientCreatedEventPayload struct {
	Email string `json:"email"`
}

type ClientCreatedEvent struct {
	Payload ClientCreatedEventPayload
	Context context.Context `json:"-"`
}

func NewClientCreatedEvent(payload ClientCreatedEventPayload, context context.Context) *ClientCreatedEvent {
	return &ClientCreatedEvent{
		Payload: payload,
		Context: context,
	}
}

func (c ClientCreatedEvent) Name() string {
	return CLIENT_CREATED_EVENT
}

func (c ClientCreatedEvent) ToMessage() ([]byte, error) {
	message, err := json.Marshal(c.Payload)
	if err != nil {
		return nil, err
	}
	return message, nil
}
