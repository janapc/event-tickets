package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAClient(t *testing.T) {
	client, err := NewClient("test", "test@test.com")
	assert.NoError(t, err)
	assert.NotEmpty(t, client)
	assert.NotEmpty(t, client.ID)
}

func TestErrorIfParametersAreWrong(t *testing.T) {
	type Data struct {
		name, email, expectedError string
	}

	data := []Data{
		{name: "", email: "test@test.com", expectedError: "the name field is mandatory"},
		{name: "test", email: "", expectedError: "the email field is mandatory"},
	}
	for _, d := range data {
		client, err := NewClient(d.name, d.email)
		if assert.Error(t, err) {
			assert.Equal(t, d.expectedError, err.Error())
		}
		assert.Empty(t, client)
	}
}
