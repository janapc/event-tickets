package application

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/janapc/event-tickets/clients/internal/domain/events"
)

type ProcessMessage struct {
	Repository domain.IClientRepository
	Messaging  domain.IMessaging
	Bus        domain.Bus
}

type InputProcessMessage struct {
	MessageID        string `json:"messageId"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	EventId          string `json:"eventId"`
	EventName        string `json:"eventName"`
	EventDescription string `json:"eventDescription"`
	EventImageUrl    string `json:"eventImageUrl"`
	Language         string `json:"language"`
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
	existsClient, err := p.Repository.GetByEmail(ctx, input.Email)
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
	event := events.NewClientCreatedEvent(uuid.New().String(),
		client.Email, ctx)
	p.Bus.Dispatch(event)
	return nil
}

func (p *ProcessMessage) sendTicket(ctx context.Context, input InputProcessMessage) error {
	event := events.NewSendTicketEvent(events.SendTicketEvent{
		MessageID:        input.MessageID,
		ClientName:       input.Name,
		Email:            input.Email,
		EventId:          input.EventId,
		EventName:        input.EventName,
		EventDescription: input.EventDescription,
		EventImageUrl:    input.EventImageUrl,
		Language:         input.Language,
		Context:          ctx,
	})
	p.Bus.Dispatch(event)
	return nil
}

func (input *InputProcessMessage) Validate() error {
	if input.Email == "" {
		return errors.New("email is required")
	}
	if input.EventId == "" {
		return errors.New("eventId is required")
	}
	return nil
}

func (p *ProcessMessage) CreateClient(ctx context.Context, input InputProcessMessage) (*domain.Client, error) {
	client, err := domain.NewClient(domain.ClientParams{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}
	return p.Repository.Save(ctx, client)
}
