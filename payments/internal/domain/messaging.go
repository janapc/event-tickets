package domain

import (
	"context"
)

type NamedEvent interface {
	Name() string
}

type ProducerParameters struct {
	Value NamedEvent
	Topic string
	Key   string
}

func (p ProducerParameters) Name() string {
	return p.Value.Name()
}

type IMessaging interface {
	Consumer(topic, groupID string, handler func(context.Context, string) error)
	Producer(params ProducerParameters) error
}
