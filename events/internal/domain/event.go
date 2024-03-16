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
	ExpirateAt  time.Time `json:"expirate_at"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

const YYYYMMDD = "2006-01-02"

func NewEvent(name, description, imageUrl string, price float64, expirateAt string) (*Event, error) {
	if err := IsValidExpirateAt(expirateAt); err != nil {
		return nil, err
	}
	event := &Event{
		ID:          uuid.New().String(),
		Name:        name,
		Description: description,
		ImageUrl:    imageUrl,
		Price:       price,
		CreatedAt:   FormatDate(time.Now(), false),
		UpdatedAt:   FormatDate(time.Now(), false),
	}
	event.ExpirateAt = FormatExpiratedAt(expirateAt)
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

func FormatExpiratedAt(expirateAt string) time.Time {
	r := regexp.MustCompile(`(\d{2})\/(\d{2})\/(\d{4})`)
	expirateAtRaw := r.ReplaceAllString(expirateAt, "${3}-${2}-$1")
	expirateAtFormatted, _ := time.Parse("2006-01-02", expirateAtRaw)
	return FormatDate(expirateAtFormatted, true)
}

func IsValidExpirateAt(expirateAt string) error {
	match, err := regexp.MatchString(`\d{2}\/\d{2}\/\d{4}`, expirateAt)
	if !match || err != nil {
		return errors.New("the field expirate_at is mandatory and should is this format DD/MM/YYYY")
	}
	return nil
}

func (p *Event) isValid() error {
	if p.Name == "" {
		return errors.New("the name field is mandatory")
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
	if p.ExpirateAt.Before(currentDate) {
		return errors.New("the expirate_at field cannot be less than current date")
	}
	return nil
}
