package domain

import (
	"errors"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Price       float64   `json:"price"`
	Currency    string    `json:"currency"`
	EventDate   time.Time `json:"event_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type EventParams struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency"`
	EventDate   string  `json:"event_date"`
}

func NewEvent(params EventParams) (*Event, error) {
	eventDate, err := FormatDate(params.EventDate)
	if err != nil {
		return &Event{}, err
	}
	event := &Event{
		Name:        params.Name,
		Description: params.Description,
		ImageUrl:    params.ImageUrl,
		Price:       params.Price,
		Currency:    params.Currency,
		EventDate:   eventDate,
	}
	if err = event.isValid(); err != nil {
		return &Event{}, err
	}
	return event, nil
}

func FormatDate(eventDate string) (time.Time, error) {
	if eventDate == "" {
		return time.Time{}, errors.New("the event_date field is mandatory")
	}
	date, err := time.Parse(time.RFC3339, eventDate)
	if err != nil {
		return time.Time{}, errors.New("the event_date field must have this format '2024-09-25T00:00:00.000Z'")
	}
	return date, nil
}

func (e *Event) isValid() error {
	if e.Name == "" {
		return errors.New("the name field is mandatory")
	}
	if e.Currency == "" {
		return errors.New("the currency field is mandatory")
	}
	if e.Description == "" {
		return errors.New("the description field is mandatory")
	}
	if e.ImageUrl == "" {
		return errors.New("the image_url field is mandatory")
	}
	if e.Price <= 0 {
		return errors.New("the price field cannot be less than or equal to zero")
	}
	currentDate := time.Now().UTC()
	if e.EventDate.Before(currentDate) {
		return errors.New("the event_date field cannot be less than the current date")
	}
	return nil
}
