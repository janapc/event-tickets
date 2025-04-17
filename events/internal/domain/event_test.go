package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const DDMMYYYY = "02/01/2006"

func TestCreateEvent(t *testing.T) {
	eventDate := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	event, err := NewEvent("Show banana", "show banana", "http://test.png", 600.40, eventDate, "BRL")
	assert.NoError(t, err)
	assert.NotEmpty(t, event.ID)
	assert.NotEmpty(t, event.CreatedAt)
	assert.Equal(t, event.Name, "Show banana")
}

func TestShouldErrorIfTheEventFieldsAreIncorrect(t *testing.T) {
	eventDate := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	eventDateWrong := time.Now().Add(-48 * time.Hour).Format(DDMMYYYY)

	type Data struct {
		Name, Description, ImageUrl string
		Price                       float64
		EventDate                   string
		Currency                    string
		ExpectedError               string
	}

	data := []Data{
		{"", "test", "http://test.png", 600.40, eventDate, "BRL", "the name field is mandatory"},
		{"test", "", "http://test.png", 600.40, eventDate, "BRL", "the description field is mandatory"},
		{"test", "test", "http://test.png", 0.0, eventDate, "BRL", "the price field cannot be less than or equal to zero"},
		{"test", "test", "http://test.png", 10.0, "", "BRL", "the field event_date is mandatory and should is this format DD/MM/YYYY"},
		{"test", "test", "http://test.png", 10.0, eventDateWrong, "BRL", "the event_date field cannot be less than current date"},
		{"test", "test", "http://test.png", 10.0, "20/1/10", "BRL", "the field event_date is mandatory and should is this format DD/MM/YYYY"},
		{"test", "test", "http://test.png", 10.0, "20/1/10", "BRL", "the field event_date is mandatory and should is this format DD/MM/YYYY"},
		{"test", "test", "", 10.0, eventDate, "BRL", "the image_url field is mandatory"},
		{"test", "test", "http://test.png", 600.40, eventDate, "", "the currency field is mandatory"},
	}
	for _, d := range data {
		p, err := NewEvent(d.Name, d.Description, d.ImageUrl, d.Price, d.EventDate, d.Currency)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), d.ExpectedError)
		}
		assert.Empty(t, p)
	}
}
