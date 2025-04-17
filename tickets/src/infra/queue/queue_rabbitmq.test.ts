import type amqplib from 'amqplib'
import { QueueRabbitmq } from './queue_rabbitmq'

const mockData = {
  content: 'testing',
}
const mockChannel: amqplib.Channel = {
  ack: jest.fn(),
  consume: jest.fn().mockImplementation((queue, callback) => {
    callback(mockData)
  }),
} as any
jest.useFakeTimers()

describe('Queue Rabbitmq', () => {
  it('should consumer messages', async () => {
    const queue = new QueueRabbitmq(mockChannel)
    const spyConsume = jest.spyOn(mockChannel, 'consume')
    const spyAck = jest.spyOn(mockChannel, 'ack')
    const Fn = jest.fn().mockResolvedValueOnce('ok')
    await expect(queue.Consumer('queue_name', Fn)).resolves.toBeUndefined()
    expect(Fn).toHaveBeenCalledTimes(1)
    expect(Fn).toHaveBeenCalledWith(mockData.content)
    expect(spyConsume).toHaveBeenCalledTimes(1)
    expect(spyAck).toHaveBeenCalledTimes(1)
  })

  it('should not consumer message if message is wrong', async () => {
    const queue = new QueueRabbitmq(mockChannel)
    const Fn = jest.fn().mockRejectedValue(new Error('message is wrong'))
    await queue.Consumer('queue_name', Fn)
    expect(Fn).toHaveBeenCalledTimes(1)
    expect(Fn).toHaveBeenCalledWith(mockData.content)
  })
})
