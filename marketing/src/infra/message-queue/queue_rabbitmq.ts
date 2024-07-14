import type amqplib from 'amqplib'
import { type IQueue } from './interface'
import { logger } from '@infra/logger/logger'

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
            logger.error(`consume message: ${e.message as string}`)
          })
      }
    })
  }
}
