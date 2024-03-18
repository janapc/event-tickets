package domain

type IClientRepository interface {
	Save(client *Client) error
	GetByEmail(email string) (*Client, error)
}
