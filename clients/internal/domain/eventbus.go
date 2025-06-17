package domain

type Event interface {
	Name() string
}

type Bus interface {
	Dispatch(event Event)
	Register(eventName string, handler func(Event))
}
