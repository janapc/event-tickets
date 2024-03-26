package application

import (
	"errors"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type InputUpdateEventDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name,omitempty"`
	ImageUrl    string  `json:"image_url,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	EventDate   string  `json:"event_date,omitempty"`
}

type UpdateEvent struct {
	Repository domain.EventRepository
}

func NewUpdateEvent(repo domain.EventRepository) *UpdateEvent {
	return &UpdateEvent{
		Repository: repo,
	}
}

func (u *UpdateEvent) Execute(input InputUpdateEventDTO) error {
	event, err := u.Repository.FindById(input.ID)
	if err != nil {
		return errors.New("event is not found")
	}
	if input.Name != "" {
		event.Name = input.Name
	}
	if input.Description != "" {
		event.Description = input.Description
	}
	if input.ImageUrl != "" {
		event.ImageUrl = input.ImageUrl
	}
	if input.Currency != "" {
		event.Currency = input.Currency
	}
	if input.Price > 0 {
		event.Price = input.Price
	}
	currentDate := domain.FormatDate(time.Now(), false)
	if input.EventDate != "" {
		if err := domain.IsValidEventDate(input.EventDate); err != nil {
			return err
		}
		eventDate := domain.FormatEventDate(input.EventDate)
		if eventDate.After(currentDate) {
			event.EventDate = eventDate
		}
	}
	event.UpdatedAt = currentDate
	err = u.Repository.Update(event)
	if err != nil {
		return err
	}
	return nil
}
