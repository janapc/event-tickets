package application

import (
	"errors"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type OutputGetEventByIdDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Currency    string    `json:"currency"`
	Price       float64   `json:"price"`
	EventDate   time.Time `json:"event_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetEventById struct {
	Repository domain.EventRepository
}

func NewGetEventById(repo domain.EventRepository) *GetEventById {
	return &GetEventById{
		Repository: repo,
	}
}

func (g *GetEventById) Execute(id string) (*OutputGetEventByIdDTO, error) {
	event, err := g.Repository.FindById(id)
	if err != nil {
		return nil, errors.New("event is not found")
	}
	return &OutputGetEventByIdDTO{
		ID:          event.ID,
		Name:        event.Name,
		Description: event.Description,
		ImageUrl:    event.ImageUrl,
		Currency:    event.Currency,
		EventDate:   event.EventDate,
		Price:       event.Price,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}, err
}
