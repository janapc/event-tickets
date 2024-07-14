import 'dotenv/config'
import { ProcessMessageTicket } from '@application/process_message_ticket'
import { connectDatabase } from '@infra/database/connect_database'
import { ticketModel } from '@infra/database/schema'
import { TicketRepository } from '@infra/database/ticket_repository'
import { connectRabbitmq } from '@infra/queue/connect_rabbitmq'
import { QueueRabbitmq } from '@infra/queue/queue_rabbitmq'
import { MailService } from '@infra/mail/mail_service'
import { logger } from '@infra/logger/logger'

async function start(): Promise<void> {
  await connectDatabase()
  const { channel, conn } = await connectRabbitmq()
  process.once('SIGINT', () => {
    channel
      .close()
      .catch((error) =>
        logger.error(`close queue - ${error.message as string}`),
      )
    conn
      .close()
      .catch((error) =>
        logger.error(`close database connection - ${error.message as string}`),
      )
  })
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
