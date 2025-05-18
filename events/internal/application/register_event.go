package application

import (
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

type RegisterEvent struct {
	Repository domain.IEventRepository
}

func NewRegisterEvent(repo domain.IEventRepository) *RegisterEvent {
	return &RegisterEvent{
		Repository: repo,
	}
}

func (r *RegisterEvent) Execute(input InputRegisterEventDTO) (*OutputRegisterEventDTO, error) {
	event, err := domain.NewEvent(input.Name, input.Description, input.ImageUrl, input.Price, input.EventDate, input.Currency)
	if err != nil {
		return nil, err
	}
	err = r.Repository.Register(event)
	if err != nil {
		return nil, err
	}
	return &OutputRegisterEventDTO{
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
