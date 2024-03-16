package application

import (
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type OutputGetEventsDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Price       float64   `json:"price"`
	ExpirateAt  time.Time `json:"expirate_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetEvents struct {
	Repository domain.EventRepository
}

func NewGetEvents(repo domain.EventRepository) *GetEvents {
	return &GetEvents{
		Repository: repo,
	}
}

func (g *GetEvents) Execute() ([]OutputGetEventsDTO, error) {
	events, err := g.Repository.List()
	if err != nil {
		return nil, err
	}
	var output []OutputGetEventsDTO
	for _, event := range events {
		output = append(output, OutputGetEventsDTO{
			ID:          event.ID,
			Name:        event.Name,
			Description: event.Description,
			ImageUrl:    event.ImageUrl,
			ExpirateAt:  event.ExpirateAt,
			Price:       event.Price,
			CreatedAt:   event.CreatedAt,
			UpdatedAt:   event.UpdatedAt,
		})
	}
	if len(output) == 0 {
		output = []OutputGetEventsDTO{}
	}
	return output, nil
}
