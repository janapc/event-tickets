import 'dotenv/config'
import { rabbitMQConnection } from '@infra/message-queue/rabbitmq_connection'
import { QueueRabbitmq } from '@infra/message-queue/queue_rabbitmq'
import { ConsumeMessagesQueue } from '@application/consume_messages_queue'
import { LeadPrismaRepository } from '@infra/dabatase/lead_prisma_repository'
import prisma from '@infra/dabatase/client'
import { logger } from '@infra/logger/logger'
import { server } from '@infra/api'

process.on('SIGINT', () => {
  rabbitMQConnection.closeRabbitmq().catch((error) => {
    logger.error(`close rabbitmq connection - ${error.message as string}`)
  })
})

export async function init(): Promise<void> {
  try {
    const channel = await rabbitMQConnection.init()
    const messageQueueRabbitmq = new QueueRabbitmq(channel)
    const repository = new LeadPrismaRepository(prisma)
    const application = new ConsumeMessagesQueue(repository)
    await messageQueueRabbitmq.Consumer(
      process.env.QUEUE_CLIENT_CREATED,
      application.execute.bind(application),
    )
    logger.info(`starting message consumption`)
    server()
  } catch (error) {
    let message = 'internal server error'
    if (error instanceof Error) message = error.message
    logger.error(`error init - ${message}`)
  }
}

void init()
