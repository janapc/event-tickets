package application

import (
	"context"
	"errors"
	"log/slog"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type RemoveEvent struct {
	Repository domain.IEventRepository
}

func NewRemoveEvent(repo domain.IEventRepository) *RemoveEvent {
	return &RemoveEvent{
		Repository: repo,
	}
}

func (r *RemoveEvent) Execute(ctx context.Context, id int64) error {
	slog.InfoContext(ctx, "starting handling of remove an event", "id", id)
	_, err := r.Repository.FindByID(ctx, id)
	if err != nil {
		return errors.New("event is not found")
	}
	err = r.Repository.Remove(ctx, id)
	if err != nil {
		return err
	}
	slog.InfoContext(ctx, "finished handling of remove an event", "id", id)
	return nil
}
