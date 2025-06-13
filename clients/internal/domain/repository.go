package domain

import "context"

type IClientRepository interface {
	Save(ctx context.Context, client *Client) (*Client, error)
	GetByEmail(ctx context.Context, email string) (*Client, error)
}
