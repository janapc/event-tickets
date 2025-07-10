package eventbus

import "sync"

type Event interface {
	Name() string
}

type HandlerFunc func(Event)

type EventBus struct {
	handlers map[string][]HandlerFunc
	mu       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]HandlerFunc),
	}
}

func (b *EventBus) Subscribe(eventName string, handler HandlerFunc) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *EventBus) Publish(event Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if hs, ok := b.handlers[event.Name()]; ok {
		for _, handler := range hs {
			go handler(event)
		}
	}
}
