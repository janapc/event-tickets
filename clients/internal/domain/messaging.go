package domain

import (
	"context"
)

type ProducerParameters struct {
	Value   Event
	Context context.Context
	Topic   string
	Key     string
}

func (p ProducerParameters) Name() string {
	return p.Value.Name()
}

type IMessaging interface {
	Consumer(topic, groupID string, handler func(context.Context, string) error)
	Producer(params ProducerParameters) error
}
