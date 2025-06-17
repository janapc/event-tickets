package events

import "context"

const CLIENT_CREATED_EVENT = "CLIENT_CREATED"

type ClientCreatedEvent struct {
	MessageID string          `json:"messageId"`
	Email     string          `json:"email"`
	Context   context.Context `json:"-"`
}

func NewClientCreatedEvent(messageId, email string, context context.Context) *ClientCreatedEvent {
	return &ClientCreatedEvent{
		MessageID: messageId,
		Email:     email,
		Context:   context,
	}
}

func (c ClientCreatedEvent) Name() string {
	return CLIENT_CREATED_EVENT
}
