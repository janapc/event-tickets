package domain

type IEventRepository interface {
	Register(event *Event) (*Event, error)
	Update(event *Event) error
	Remove(id int64) error
	List() ([]*Event, error)
	FindByID(id int64) (*Event, error)
}
