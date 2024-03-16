package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const DDMMYYYY = "02/01/2006"

func TestCreateEvent(t *testing.T) {
	expirateAt := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	event, err := NewEvent("Show banana", "show banana", "http://test.png", 600.40, expirateAt)
	assert.NoError(t, err)
	assert.NotEmpty(t, event.ID)
	assert.NotEmpty(t, event.CreatedAt)
	assert.Equal(t, event.Name, "Show banana")
}

func TestShouldErrorIfTheEventFieldsAreIncorrect(t *testing.T) {
	expirateAt := time.Now().Add(48 * time.Hour).Format(DDMMYYYY)
	expirateAtWrong := time.Now().Add(-48 * time.Hour).Format(DDMMYYYY)

	type Data struct {
		name, description, imageUrl string
		price                       float64
		expirateAt                  string
		ExpectedError               string
	}

	data := []Data{
		{"", "test", "http://test.png", 600.40, expirateAt, "the name field is mandatory"},
		{"test", "", "http://test.png", 600.40, expirateAt, "the description field is mandatory"},
		{"test", "test", "http://test.png", 0.0, expirateAt, "the price field cannot be less than or equal to zero"},
		{"test", "test", "http://test.png", 10.0, "", "the field expirate_at is mandatory and should is this format DD/MM/YYYY"},
		{"test", "test", "http://test.png", 10.0, expirateAtWrong, "the expirate_at field cannot be less than current date"},
		{"test", "test", "http://test.png", 10.0, "20/1/10", "the field expirate_at is mandatory and should is this format DD/MM/YYYY"},
		{"test", "test", "http://test.png", 10.0, "20/1/10", "the field expirate_at is mandatory and should is this format DD/MM/YYYY"},
		{"test", "test", "", 10.0, expirateAt, "the image_url field is mandatory"},
	}
	for _, d := range data {
		p, err := NewEvent(d.name, d.description, d.imageUrl, d.price, d.expirateAt)
		if assert.Error(t, err) {
			assert.Equal(t, err.Error(), d.ExpectedError)
		}
		assert.Empty(t, p)
	}
}
