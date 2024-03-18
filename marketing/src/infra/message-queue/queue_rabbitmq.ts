import type amqplib from 'amqplib'
import { type IQueue } from './interface'

export class QueueRabbitmq implements IQueue {
  constructor(private readonly channel: amqplib.Channel) {}

  async Consumer(
    queueName: string,
    fn: (message: string) => Promise<void>,
  ): Promise<void> {
    await this.channel.consume(queueName, (msg): void => {
      if (msg !== null) {
        const content = msg.content.toString()
        fn(content)
          .then((r) => {
            this.channel.ack(msg)
          })
          .catch((e) => {
            throw new Error(e.message as string)
          })
      }
    })
  }
}
