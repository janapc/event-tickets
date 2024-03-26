import 'dotenv/config'
import { connectRabbitmq } from '@infra/message-queue/connect_rabbitmq'
import { QueueRabbitmq } from '@infra/message-queue/queue_rabbitmq'
import { ConsumeMessagesQueue } from '@application/consume_messages_queue'
import { LeadPrismaRepository } from '@infra/dabatase/lead_prisma_repository'
import prisma from '@infra/dabatase/client'
import { logger } from '@infra/logger/logger'

export async function start(): Promise<void> {
  try {
    const { conn, channel } = await connectRabbitmq()
    process.once('SIGINT', () => {
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
    })
    const messageQueueRabbitmq = new QueueRabbitmq(channel)
    const repository = new LeadPrismaRepository(prisma)
    const application = new ConsumeMessagesQueue(repository)
    await messageQueueRabbitmq.Consumer(
      process.env.QUEUE_CLIENT_CREATED,
      application.execute.bind(application),
    )
  } catch (error) {
    let message = 'internal server error'
    if (error instanceof Error) message = error.message
    logger.error(`error start queue: ${message}`)
  }
}
void start()
