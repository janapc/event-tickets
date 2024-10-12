import 'dotenv/config'
import queueConnection from '@infra/messageQueue/queue_connection'
import { QueueRabbitmq } from '@infra/messageQueue/queue_rabbitmq'
import { ConsumeMessagesQueue } from '@application/consume_messages_queue'
import Database from '@infra/dabatase'
import { LeadPrismaRepository } from '@infra/dabatase/lead_prisma_repository'
import { logger } from '@infra/logger/logger'

process.on('SIGINT', () => {
  queueConnection.close()
  Database.getInstance().close()
})

if (process.env.NODE_ENV !== 'development') {
  import('@infra/trace')
}

async function init(): Promise<void> {
  try {
    const channel = await queueConnection.init()
    const queueRabbitmq = new QueueRabbitmq(channel)
    const database = Database.getInstance()
    const repository = new LeadPrismaRepository(database.connection)
    const application = new ConsumeMessagesQueue(repository)
    await queueRabbitmq.Consumer(
      process.env.QUEUE_CLIENT_CREATED,
      application.execute.bind(application),
    )
  } catch (error) {
    logger.error(`error init queue ${String(error)}`)
  }
}

void init()
