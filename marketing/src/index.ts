import 'dotenv/config'
import type amqplib from 'amqplib'
import type * as http from 'http'
import { connectRabbitmq } from '@infra/message-queue/connect_rabbitmq'
import { QueueRabbitmq } from '@infra/message-queue/queue_rabbitmq'
import { ConsumeMessagesQueue } from '@application/consume_messages_queue'
import { LeadPrismaRepository } from '@infra/dabatase/lead_prisma_repository'
import prisma from '@infra/dabatase/client'
import { logger } from '@infra/logger/logger'
import { server } from '@infra/api'

export async function start(): Promise<void> {
  try {
    const { conn, channel } = await connectRabbitmq()
    const messageQueueRabbitmq = new QueueRabbitmq(channel)
    const repository = new LeadPrismaRepository(prisma)
    const application = new ConsumeMessagesQueue(repository)
    logger.info(`starting message consumption`)
    await messageQueueRabbitmq.Consumer(
      process.env.QUEUE_CLIENT_CREATED,
      application.execute.bind(application),
    )
    const s = server()
    process.on('SIGINT', () => {
      gracefulShutdown(conn, channel, s)
    })
    process.on('SIGTERM', () => {
      gracefulShutdown(conn, channel, s)
    })
  } catch (error) {
    let message = 'internal server error'
    if (error instanceof Error) message = error.message
    logger.error(message)
  }
}

function gracefulShutdown(
  conn: amqplib.Connection,
  channel: amqplib.Channel,
  server: http.Server,
): void {
  channel
    .close()
    .catch((e) =>
      logger.error(`error disconnect channel: ${e.message as string}`),
    )
  conn
    .close()
    .catch((e) =>
      logger.error(`error disconnect connection: ${e.message as string}`),
    )
  server.close((err) => {
    logger.info('Http server closed.')
    process.exit(err ? 1 : 0)
  })
}
void start()
