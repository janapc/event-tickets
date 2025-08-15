package events

import (
	"context"
	"encoding/json"
)

const SEND_TICKET_EVENT = "SEND_TICKET"

type SendTicketEventPayload struct {
	ClientName       string `json:"name"`
	Email            string `json:"email"`
	EventId          string `json:"eventId"`
	EventName        string `json:"eventName"`
	EventDescription string `json:"eventDescription"`
	EventImageUrl    string `json:"eventImageUrl"`
	Language         string `json:"language"`
}
type SendTicketEvent struct {
	Payload SendTicketEventPayload
	Context context.Context `json:"-"`
}

func NewSendTicketEvent(payload SendTicketEventPayload, context context.Context) *SendTicketEvent {
	return &SendTicketEvent{
		Payload: payload,
		Context: context,
	}
}

func (s SendTicketEvent) Name() string {
	return SEND_TICKET_EVENT
}

func (s SendTicketEvent) ToMessage() ([]byte, error) {
	message, err := json.Marshal(s.Payload)
	if err != nil {
		return nil, err
	}
	return message, nil
}
