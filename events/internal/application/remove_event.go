package application

import (
	"context"
	"errors"

	"github.com/janapc/event-tickets/events/internal/domain"
	"github.com/janapc/event-tickets/events/internal/infra/logger"
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
	logger.Logger.WithContext(ctx).Infof("removing event with id: %d", id)
	_, err := r.Repository.FindByID(ctx, id)
	if err != nil {
		return errors.New("event is not found")
	}
	err = r.Repository.Remove(ctx, id)
	if err != nil {
		return err
	}
	logger.Logger.WithContext(ctx).Infof("event with id %d removed successfully", id)
	return nil
}
