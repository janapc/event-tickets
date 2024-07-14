import type amqplib from 'amqplib'
import { type IQueue } from './queue'
import { logger } from '../logger/logger'

export class QueueRabbitmq implements IQueue {
  constructor(private readonly channel: amqplib.Channel) {}

  async Consumer(
    queueName: string,
    fn: (message: string) => Promise<void>,
  ): Promise<void> {
    await this.channel.consume(queueName, (msg): void => {
      if (msg !== null) {
        const content = msg.content.toString()
        logger.info(`consumer message - ${content}`)
        fn(content)
          .then(() => {
            this.channel.ack(msg)
          })
          .catch((error) => {
            logger.error(`consumer message - ${error.message as string}`)
          })
      }
    })
  }
}
