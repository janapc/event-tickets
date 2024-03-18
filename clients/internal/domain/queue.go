package domain

type IQueue interface {
	Consumer(queueName string, workerPoolSize int) error
	Producer(queueName string, message []byte) error
}
