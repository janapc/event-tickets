package domain

type IClientRepository interface {
	Save(client *Client) (*Client, error)
	GetByEmail(email string) (*Client, error)
}
