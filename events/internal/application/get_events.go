package application

import (
	"context"
	"math"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/janapc/event-tickets/events/pkg/pagination"
)

type EventsOutputDTO struct {
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

type OutputGetEventsDTO struct {
	Events     []EventsOutputDTO     `json:"events"`
	Pagination pagination.Pagination `json:"pagination"`
}

type GetEvents struct {
	Repository domain.IEventRepository
}

func NewGetEvents(repo domain.IEventRepository) *GetEvents {
	return &GetEvents{
		Repository: repo,
	}
}

func formatPagination(page, size int) (int, int) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 10
	}
	return page, LimitPageSize(size)
}

func LimitPageSize(requestedSize int) int {
	const maxPageSize = 100
	return int(math.Min(float64(requestedSize), float64(maxPageSize)))
}

func (g *GetEvents) Execute(ctx context.Context, page, size int) (OutputGetEventsDTO, error) {
	page, size = formatPagination(page, size)
	events, pagination, err := g.Repository.List(ctx, page, size)
	if err != nil {
		return OutputGetEventsDTO{}, err
	}
	if len(events) == 0 {
		return OutputGetEventsDTO{}, nil
	}
	var eventsResponse []EventsOutputDTO
	for _, event := range events {
		eventsResponse = append(eventsResponse, EventsOutputDTO{
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
	output := OutputGetEventsDTO{
		Events:     eventsResponse,
		Pagination: pagination,
	}
	return output, nil
}
