package domain

import (
	"context"
)

type IMessaging interface {
	Consumer(topic, groupID string, handler func(context.Context, string) error)
	Producer(message []byte, ctx context.Context, topic string, key string) error
}
