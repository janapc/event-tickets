package domain

type EventRepository interface {
	Register(event *Event) (*Event, error)
	Update(event *Event) error
	Remove(id string) error
	List() ([]*Event, error)
	FindByID(id string) (*Event, error)
}
