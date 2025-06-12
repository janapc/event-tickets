package domain

type IMessaging interface {
	Consumer(topic, groupID string, handler func(string) error)
	Producer(topic string, key, value []byte) error
}
