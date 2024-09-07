package domain

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"image_url"`
	Price       float64   `json:"price"`
	Currency    string    `json:"currency"`
	EventDate   time.Time `json:"event_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewEvent(name, description, imageUrl string, price float64, eventDate string, currency string) (*Event, error) {
	if err := IsValidEventDate(eventDate); err != nil {
		return nil, err
	}
	event := &Event{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		ImageUrl:    imageUrl,
		Price:       price,
		Currency:    currency,
		CreatedAt:   FormatDate(time.Now(), false),
		UpdatedAt:   FormatDate(time.Now(), false),
	}
	event.EventDate = FormatEventDate(eventDate)
	if err := event.isValid(); err != nil {
		return nil, err
	}
	return event, nil
}

func FormatDate(d time.Time, finalDate bool) time.Time {
	if finalDate {
		return time.Date(d.Year(), d.Month(),
			d.Day(), 23, 59, 59, 00, time.UTC)
	}
	return time.Date(d.Year(), d.Month(),
		d.Day(), d.Hour(), d.Minute(), d.Second(), 00, d.Location()).UTC()
}

func FormatEventDate(eventDate string) time.Time {
	r := regexp.MustCompile(`(\d{2})/(\d{2})/(\d{4})`)
	eventDateRaw := r.ReplaceAllString(eventDate, "${3}-${2}-$1")
	eventDateFormatted, _ := time.Parse("2006-01-02", eventDateRaw)
	return FormatDate(eventDateFormatted, true)
}

func IsValidEventDate(eventDate string) error {
	match, err := regexp.MatchString(`\d{2}/\d{2}/\d{4}`, eventDate)
	if !match || err != nil {
		return errors.New("the field event_date is mandatory and should is this format DD/MM/YYYY")
	}
	return nil
}

func (p *Event) isValid() error {
	if p.Name == "" {
		return errors.New("the name field is mandatory")
	}
	if p.Currency == "" {
		return errors.New("the currency field is mandatory")
	}
	if p.Description == "" {
		return errors.New("the description field is mandatory")
	}
	if p.ImageUrl == "" {
		return errors.New("the image_url field is mandatory")
	}
	if p.Price <= 0 {
		return errors.New("the price field cannot be less than or equal to zero")
	}
	currentDate := time.Now().UTC()
	if p.EventDate.Before(currentDate) {
		return errors.New("the event_date field cannot be less than current date")
	}
	return nil
}
