package application

import (
	"context"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type InputRegisterEventDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	EventDate   string  `json:"event_date"`
}

type OutputRegisterEventDTO struct {
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

type RegisterEvent struct {
	Repository domain.IEventRepository
}

func NewRegisterEvent(repo domain.IEventRepository) *RegisterEvent {
	return &RegisterEvent{
		Repository: repo,
	}
}

func (r *RegisterEvent) Execute(ctx context.Context, input InputRegisterEventDTO) (*OutputRegisterEventDTO, error) {
	event, err := domain.NewEvent(domain.EventParams{
		Name: input.Name, Description: input.Description, ImageUrl: input.ImageUrl, Price: input.Price, EventDate: input.EventDate, Currency: input.Currency,
	})
	if err != nil {
		return nil, err
	}
	newEvent, err := r.Repository.Register(ctx, event)
	if err != nil {
		return nil, err
	}
	return &OutputRegisterEventDTO{
		ID:          newEvent.ID,
		Name:        newEvent.Name,
		Description: newEvent.Description,
		ImageUrl:    newEvent.ImageUrl,
		Currency:    newEvent.Currency,
		EventDate:   newEvent.EventDate,
		Price:       newEvent.Price,
		CreatedAt:   newEvent.CreatedAt,
		UpdatedAt:   newEvent.UpdatedAt,
	}, err
}
