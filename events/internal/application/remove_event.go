package application

import (
	"errors"

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

func (r *RemoveEvent) Execute(id int64) error {
	_, err := r.Repository.FindByID(id)
	if err != nil {
		return errors.New("event is not found")
	}
	err = r.Repository.Remove(id)
	if err != nil {
		return err
	}
	return nil
}
