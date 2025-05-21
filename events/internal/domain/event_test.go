package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateEvent(t *testing.T) {
	eventDate := time.Now().Add(48 * time.Hour).Format(time.RFC3339)
	params := EventParams{
		Name:        "Show Test",
		Description: "Description test",
		ImageUrl:    "http://localhost:3000/image.png",
		Price:       20.99,
		EventDate:   eventDate,
		Currency:    "BRL",
	}
	event, err := NewEvent(params)
	assert.NoError(t, err)
	assert.Equal(t, event.Name, params.Name)
	assert.Equal(t, event.Description, params.Description)
	assert.Equal(t, event.ImageUrl, params.ImageUrl)
	assert.Equal(t, event.Price, params.Price)
	assert.NotNil(t, event.EventDate)
	assert.Equal(t, event.Currency, params.Currency)
}

func TestShouldErrorIfTheEventFieldsAreIncorrect(t *testing.T) {
	eventDate := time.Now().Add(48 * time.Hour).Format(time.RFC3339)
	eventDateWrong := "20/10/2001"
	eventDatePast := time.Now().Add(-48 * time.Hour).Format(time.RFC3339)

	type DataType struct {
		EventParams   EventParams
		ExpectedError string
	}

	data := []DataType{
		{EventParams{Description: "test", ImageUrl: "http://test.png", EventDate: eventDate, Currency: "BRL", Price: 20.99}, "the name field is mandatory"},
		{EventParams{Name: "test", ImageUrl: "http://test.png", EventDate: eventDate, Currency: "BRL", Price: 20.99}, "the description field is mandatory"},
		{EventParams{Name: "test", Description: "test", ImageUrl: "http://test.png", EventDate: eventDate, Currency: "BRL"}, "the price field cannot be less than or equal to zero"},
		{EventParams{Name: "test", Description: "test", Price: 20.99, EventDate: eventDate, Currency: "BRL"}, "the image_url field is mandatory"},
		{EventParams{Name: "test", Description: "test", ImageUrl: "http://test.png", EventDate: eventDate, Price: 20.99}, "the currency field is mandatory"},
		{EventParams{Name: "test", Description: "test", ImageUrl: "http://test.png", Price: 20.99, Currency: "BRL"}, "the event_date field is mandatory"},
		{EventParams{Name: "test", Description: "test", ImageUrl: "http://test.png", Price: 20.99, Currency: "BRL", EventDate: eventDateWrong}, "the event_date field must have this format '2024-09-25T00:00:00.000Z'"},
		{EventParams{Name: "test", Description: "test", ImageUrl: "http://test.png", Price: 20.99, Currency: "BRL", EventDate: eventDatePast}, "the event_date field cannot be less than the current date"},
	}
	for _, d := range data {
		p, err := NewEvent(d.EventParams)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), d.ExpectedError)
		}
		assert.Empty(t, p)
	}
}
