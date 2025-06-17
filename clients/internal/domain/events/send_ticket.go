package events

import "context"

const SEND_TICKET_EVENT = "SEND_TICKET"

type SendTicketEvent struct {
	MessageID        string          `json:"messageId"`
	ClientName       string          `json:"name"`
	Email            string          `json:"email"`
	EventId          string          `json:"eventId"`
	EventName        string          `json:"eventName"`
	EventDescription string          `json:"eventDescription"`
	EventImageUrl    string          `json:"eventImageUrl"`
	Language         string          `json:"language"`
	Context          context.Context `json:"-"`
}

func NewSendTicketEvent(params SendTicketEvent) *SendTicketEvent {
	return &SendTicketEvent{
		MessageID:        params.MessageID,
		ClientName:       params.ClientName,
		Email:            params.Email,
		EventId:          params.EventId,
		EventName:        params.EventName,
		EventDescription: params.EventDescription,
		EventImageUrl:    params.EventImageUrl,
		Language:         params.Language,
		Context:          params.Context,
	}
}

func (s SendTicketEvent) Name() string {
	return SEND_TICKET_EVENT
}
