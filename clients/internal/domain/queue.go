package domain

type IQueue interface {
	Consumer(queueName string, workerPoolSize int)
	Producer(queueName string, message []byte) error
}
