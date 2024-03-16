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
	ExpirateAt  string  `json:"expirate_at"`
}

type OutputRegisterEventDTO struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Price       float64   `json:"price"`
	ExpirateAt  time.Time `json:"expirate_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type RegisterEvent struct {
	Repository domain.EventRepository
}

func NewRegisterEvent(repo domain.EventRepository) *RegisterEvent {
	return &RegisterEvent{
		Repository: repo,
	}
}

func (r *RegisterEvent) Execute(input InputRegisterEventDTO) (*OutputRegisterEventDTO, error) {
	event, err := domain.NewEvent(input.Name, input.Description, input.ImageUrl, input.Price, input.ExpirateAt)
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
		ExpirateAt:  event.ExpirateAt,
		Price:       event.Price,
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}, err
}
