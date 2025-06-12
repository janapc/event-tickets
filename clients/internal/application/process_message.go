package application

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"github.com/janapc/event-tickets/clients/internal/domain"
)

type ProcessMessage struct {
	Repository         domain.IClientRepository
	Messaging          domain.IMessaging
	ClientCreatedQueue string
	SendTicketQueue    string
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

type ClientCreated struct {
	MessageID string `json:"messageId"`
	Email     string `json:"email"`
	HasClient bool   `json:"hasClient"`
}

func NewProcessMessage(repo domain.IClientRepository, messaging domain.IMessaging, clientCreatedQueue, sendTicketQueue string) *ProcessMessage {
	return &ProcessMessage{
		Repository:         repo,
		Messaging:          messaging,
		ClientCreatedQueue: clientCreatedQueue,
		SendTicketQueue:    sendTicketQueue,
	}
}

func (p *ProcessMessage) Execute(input string) error {
	var data InputProcessMessage
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		return err
	}
	if err := data.Validate(); err != nil {
		return err
	}

	if err := p.processClient(data); err != nil {
		return err
	}
	return p.sendTicket(data)
}

func (p *ProcessMessage) processClient(input InputProcessMessage) error {
	existsClient, err := p.Repository.GetByEmail(input.Email)
	if existsClient == nil || err != nil {
		newClient, err := p.CreateClient(input)
		if err != nil {
			return err
		}
		return p.notifyClientCreated(newClient)
	}
	return nil
}

func (p *ProcessMessage) notifyClientCreated(client *domain.Client) error {
	clientCreated := ClientCreated{
		MessageID: uuid.New().String(),
		Email:     client.Email,
		HasClient: true,
	}
	clientCreatedJson, err := json.Marshal(clientCreated)
	if err != nil {
		return err
	}
	return p.Messaging.Producer(p.ClientCreatedQueue, []byte(clientCreated.MessageID), clientCreatedJson)
}

func (p *ProcessMessage) sendTicket(input InputProcessMessage) error {
	sendTicketJson, err := json.Marshal(input)
	if err != nil {
		return err
	}
	return p.Messaging.Producer(p.SendTicketQueue, []byte(input.MessageID), sendTicketJson)
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

func (p *ProcessMessage) CreateClient(input InputProcessMessage) (*domain.Client, error) {
	client, err := domain.NewClient(domain.ClientParams{
		Name:  input.Name,
		Email: input.Email,
	})
	if err != nil {
		return nil, err
	}
	return p.Repository.Save(client)
}
