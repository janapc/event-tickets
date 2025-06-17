package eventbus

import (
	"sync"

	"github.com/janapc/event-tickets/clients/internal/domain"
	"github.com/janapc/event-tickets/clients/internal/infra/logger"
)

type EventBus struct {
	handlers map[string][]func(domain.Event)
	mu       sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		handlers: make(map[string][]func(domain.Event)),
	}
}

func (b *EventBus) Register(eventName string, handler func(domain.Event)) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[eventName] = append(b.handlers[eventName], handler)
}

func (b *EventBus) Dispatch(event domain.Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if handlers, ok := b.handlers[event.Name()]; ok {
		for _, handler := range handlers {
			handler(event)
		}
	} else {
		logger.Logger.Warnf("[EventBus] No handler for event: %s", event.Name())
	}
}
