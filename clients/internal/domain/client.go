package domain

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Client struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewClient(name string, email string) (*Client, error) {
	c := &Client{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now().UTC(),
	}
	if err := c.isValid(); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Client) isValid() error {
	if c.Name == "" {
		return errors.New("the name field is mandatory")
	}
	if c.Email == "" {
		return errors.New("the email field is mandatory")
	}
	return nil
}
