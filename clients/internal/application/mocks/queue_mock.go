package mocks

type QueueMock struct {
	Messages []string
}

func NewQueueMock() *QueueMock {
	return &QueueMock{
		Messages: []string{},
	}
}

func (q *QueueMock) Consumer(queueName string, workerPoolSize int) error {
	return nil
}

func (q *QueueMock) Producer(queueName string, message []byte) error {
	q.Messages = append(q.Messages, string(message))
	return nil
}
