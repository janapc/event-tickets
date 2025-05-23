package application

import (
	"errors"
	"time"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type InputUpdateEventDTO struct {
	Name        string  `json:"name,omitempty"`
	ImageUrl    string  `json:"image_url,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Currency    string  `json:"currency,omitempty"`
	EventDate   string  `json:"event_date,omitempty"`
}

type UpdateEvent struct {
	Repository domain.IEventRepository
}

func NewUpdateEvent(repo domain.IEventRepository) *UpdateEvent {
	return &UpdateEvent{
		Repository: repo,
	}
}

func (u *UpdateEvent) Execute(id int64, input InputUpdateEventDTO) error {
	event, err := u.Repository.FindByID(id)
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
	if input.EventDate != "" {
		eventDate, err := domain.FormatDate(input.EventDate)
		if err != nil {
			return err
		}
		currentDate := time.Now().UTC()
		if eventDate.Before(currentDate) {
			return errors.New("the event_date field cannot be less than the current date")
		}
		event.EventDate = eventDate
	}
	event.UpdatedAt = time.Now()
	err = u.Repository.Update(event)
	if err != nil {
		return err
	}
	return nil
}
