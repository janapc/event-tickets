package domain

import (
	"errors"
	"time"
)

type ClientParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Client struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func NewClient(params ClientParams) (*Client, error) {
	client := &Client{
		Name:  params.Name,
		Email: params.Email,
	}
	if err := client.isValid(); err != nil {
		return nil, err
	}
	return client, nil
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
