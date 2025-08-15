package application

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/janapc/event-tickets/clients/internal/domain/events"
)

type ProcessMessage struct {
	Repository domain.IClientRepository
	Messaging  domain.IMessaging
	Bus        domain.Bus
}

type InputProcessMessage struct {
	UserName         string `json:"user_name"`
	UserEmail        string `json:"user_email"`
	EventId          string `json:"event_id"`
	EventName        string `json:"event_name"`
	EventDescription string `json:"event_description"`
	EventImageUrl    string `json:"event_image_url"`
	UserLanguage     string `json:"user_language"`
}

func NewProcessMessage(repo domain.IClientRepository, messaging domain.IMessaging, bus domain.Bus) *ProcessMessage {
	return &ProcessMessage{
		Repository: repo,
		Messaging:  messaging,
		Bus:        bus,
	}
}

func (p *ProcessMessage) Execute(ctx context.Context, input string) error {
	var data InputProcessMessage
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return err
	}
	if err := data.Validate(); err != nil {
		return err
	}

	if err := p.processClient(ctx, data); err != nil {
		return err
	}
	return p.sendTicket(ctx, data)
}

func (p *ProcessMessage) processClient(ctx context.Context, input InputProcessMessage) error {
	existsClient, err := p.Repository.GetByEmail(ctx, input.UserEmail)
	if existsClient == nil || err != nil {
		newClient, err := p.CreateClient(ctx, input)
		if err != nil {
			return err
		}
		return p.notifyClientCreated(ctx, newClient)
	}
	return nil
}

func (p *ProcessMessage) notifyClientCreated(ctx context.Context, client *domain.Client) error {
	event := events.NewClientCreatedEvent(events.ClientCreatedEventPayload{
		Email: client.Email,
	}, ctx)
	p.Bus.Dispatch(event)
	return nil
}

func (p *ProcessMessage) sendTicket(ctx context.Context, input InputProcessMessage) error {
	event := events.NewSendTicketEvent(events.SendTicketEventPayload{
		ClientName:       input.UserName,
		Email:            input.UserEmail,
		EventId:          input.EventId,
		EventName:        input.EventName,
		EventDescription: input.EventDescription,
		EventImageUrl:    input.EventImageUrl,
		Language:         input.UserLanguage,
	}, ctx)
	p.Bus.Dispatch(event)
	return nil
}

func (input *InputProcessMessage) Validate() error {
	if input.UserEmail == "" {
		return errors.New("email is required")
	}
	if input.EventId == "" {
		return errors.New("eventId is required")
	}
	return nil
}

func (p *ProcessMessage) CreateClient(ctx context.Context, input InputProcessMessage) (*domain.Client, error) {
	client, err := domain.NewClient(domain.ClientParams{
		Name:  input.UserName,
		Email: input.UserEmail,
	})
	if err != nil {
		return nil, err
	}
	return p.Repository.Save(ctx, client)
}
