package application

import (
	"context"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type OutputGetEventsDTO struct {
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

type GetEvents struct {
	Repository domain.IEventRepository
}

func NewGetEvents(repo domain.IEventRepository) *GetEvents {
	return &GetEvents{
		Repository: repo,
	}
}

func (g *GetEvents) Execute(ctx context.Context) ([]OutputGetEventsDTO, error) {
	events, err := g.Repository.List(ctx)
	if err != nil {
		return nil, err
	}
	if len(events) == 0 {
		return []OutputGetEventsDTO{}, nil
	}
	var output []OutputGetEventsDTO
	for _, event := range events {
		output = append(output, OutputGetEventsDTO{
			ID:          event.ID,
			Name:        event.Name,
			Description: event.Description,
			ImageUrl:    event.ImageUrl,
			EventDate:   event.EventDate,
			Currency:    event.Currency,
			Price:       event.Price,
			CreatedAt:   event.CreatedAt,
			UpdatedAt:   event.UpdatedAt,
		})
	}
	return output, nil
}
