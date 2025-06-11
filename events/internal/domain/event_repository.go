package domain

import (
	"context"

	"github.com/janapc/event-tickets/events/pkg/pagination"
)

type IEventRepository interface {
	Register(ctx context.Context, event *Event) (*Event, error)
	Update(ctx context.Context, event *Event) error
	Remove(ctx context.Context, id int64) error
	List(ctx context.Context, page, size int) ([]Event, pagination.Pagination, error)
	FindByID(ctx context.Context, id int64) (*Event, error)
}
