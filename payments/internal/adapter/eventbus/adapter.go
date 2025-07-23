package eventbus

import (
	"github.com/janapc/event-tickets/payments/internal/domain"
	pkgEventBus "github.com/janapc/event-tickets/payments/pkg/eventbus"
)

type adapter struct {
	bus *pkgEventBus.EventBus
}

func NewBusAdapter(bus *pkgEventBus.EventBus) domain.IEventBus {
	return &adapter{bus: bus}
}

func (a *adapter) Publish(e domain.Event) {
	a.bus.Publish(e)
}

func (a *adapter) Subscribe(name string, handler func(domain.Event)) {
	a.bus.Subscribe(name, func(e pkgEventBus.Event) {
		handler(e)
	})
}
