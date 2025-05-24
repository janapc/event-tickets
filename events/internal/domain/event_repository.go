package domain

import "context"

type IEventRepository interface {
	Register(ctx context.Context, event *Event) (*Event, error)
	Update(ctx context.Context, event *Event) error
	Remove(ctx context.Context, id int64) error
	List(ctx context.Context) ([]*Event, error)
	FindByID(ctx context.Context, id int64) (*Event, error)
}
