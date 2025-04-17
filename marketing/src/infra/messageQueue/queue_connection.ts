import amqplib from 'amqplib'
import { logger } from '@infra/logger/logger'

class QueueConnection {
  connection!: amqplib.Connection

  async init(): Promise<amqplib.Channel> {
    this.connection = await amqplib.connect(process.env.RABBITMQ_URL)
    const channel = await this.connection.createChannel()
    await channel.assertQueue(process.env.QUEUE_CLIENT_CREATED)
    return channel
  }

  close(): void {
    this.connection
      .close()
      .catch((error) =>
        logger.error(`queue connection close error ${String(error)}`),
      )
    logger.info(`queue connection closed`)
  }
}

export default new QueueConnection()
