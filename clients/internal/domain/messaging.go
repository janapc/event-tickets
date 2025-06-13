package domain

import "context"

type IMessaging interface {
	Consumer(topic, groupID string, handler func(context.Context, string) error)
	Producer(topic string, key, value []byte, ctx context.Context) error
}
