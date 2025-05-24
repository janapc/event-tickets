package application

import (
	"context"
	"errors"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type OutputGetEventByIdDTO struct {
	ID          int64     `json:"id"`
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
	Repository domain.IEventRepository
}

func NewGetEventById(repo domain.IEventRepository) *GetEventById {
	return &GetEventById{
		Repository: repo,
	}
}

func (g *GetEventById) Execute(ctx context.Context, id int64) (*OutputGetEventByIdDTO, error) {
	event, err := g.Repository.FindByID(ctx, id)
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
