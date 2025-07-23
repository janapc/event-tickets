package domain

type Event interface {
	Name() string
}

type IEventBus interface {
	Publish(event Event)
	Subscribe(name string, handler func(Event))
}
