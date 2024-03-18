package application

import (
	"encoding/json"
	"os"

	"github.com/janapc/event-tickets/clients/internal/domain"
)

type ProcessMessage struct {
	Repository domain.IClientRepository
}

type InputProcessMessage struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	EventId          string `json:"eventId"`
	EventName        string `json:"eventName"`
	EventDescription string `json:"eventDescription"`
	EventImageUrl    string `json:"eventImageUrl"`
}

type MessageClientCreated struct {
	Email     string `json:"email"`
	HasClient bool   `json:"hasClient"`
}

func NewProcessMessage(repo domain.IClientRepository) *ProcessMessage {
	return &ProcessMessage{
		Repository: repo,
	}
}

func (p *ProcessMessage) Execute(input InputProcessMessage, fn domain.IQueue) error {
	_, err := p.Repository.GetByEmail(input.Email)
	if err != nil {
		client, err := domain.NewClient(input.Name, input.Email)
		if err != nil {
			return err
		}
		err = p.Repository.Save(client)
		if err != nil {
			return err
		}
		messageClientCreated := MessageClientCreated{
			Email:     input.Email,
			HasClient: true,
		}
		msgQueueClientCreated, _ := json.Marshal(messageClientCreated)
		queueClientCreated := os.Getenv("QUEUE_CLIENT_CREATED")
		fn.Producer(queueClientCreated, msgQueueClientCreated)
	}

	msgQueueSendTicket, _ := json.Marshal(input)
	queueSendTicket := os.Getenv("QUEUE_SEND_TICKET")
	fn.Producer(queueSendTicket, msgQueueSendTicket)
	return nil
}
