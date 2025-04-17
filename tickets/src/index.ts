import 'dotenv/config'
import { ProcessMessageTicket } from '@application/process_message_ticket'
import {
  closeDatabase,
  initDatabase,
} from '@infra/database/database_connection'
import { ticketModel } from '@infra/database/schema'
import { TicketRepository } from '@infra/database/ticket_repository'
import { rabbitMQConnection } from '@infra/queue/rabbitmq_connection'
import { QueueRabbitmq } from '@infra/queue/queue_rabbitmq'
import { MailService } from '@infra/mail/mail_service'
import { logger } from '@infra/logger/logger'

process.once('SIGINT', () => {
  rabbitMQConnection
    .closeRabbitmq()
    .catch((error) =>
      logger.error(`close rabbitmq connection - ${error.message as string}`),
    )
  closeDatabase().catch((error) =>
    logger.error(`close database connection - ${error.message as string}`),
  )
})

async function start(): Promise<void> {
  await initDatabase()
  const channel = await rabbitMQConnection.init()
  const queue = new QueueRabbitmq(channel)
  const repository = new TicketRepository(ticketModel)
  const mail = new MailService()
  const application = new ProcessMessageTicket(repository, mail)
  logger.info(`starting message consumption`)
  await queue.Consumer(
    process.env.QUEUE_SEND_TICKET,
    application.execute.bind(application),
  )
}

void start()
