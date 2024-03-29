package domain

type EventRepository interface {
	Register(event *Event) error
	Update(event *Event) error
	Remove(id string) error
	List() ([]Event, error)
	FindById(id string) (*Event, error)
}
