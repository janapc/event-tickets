package application

import (
	"errors"

	"github.com/janapc/event-tickets/events/internal/domain"
)

type RemoveEvent struct {
	Repository domain.EventRepository
}

func NewRemoveEvent(repo domain.EventRepository) *RemoveEvent {
	return &RemoveEvent{
		Repository: repo,
	}
}

func (r *RemoveEvent) Execute(id string) error {
	_, err := r.Repository.FindById(id)
	if err != nil {
		return errors.New("event is not found")
	}
	err = r.Repository.Remove(id)
	if err != nil {
		return err
	}
	return nil
}
