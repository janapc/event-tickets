import 'dotenv/config'
import { connectRabbitmq } from '@infra/message-queue/connect_rabbitmq'
import { QueueRabbitmq } from '@infra/message-queue/queue_rabbitmq'
import { ConsumeMessagesQueue } from '@application/consume_messages_queue'
import { LeadPrismaRepository } from '@infra/dabatase/lead_prisma_repository'
import prisma from '@infra/dabatase/client'

export async function start(): Promise<void> {
  try {
    const { conn, channel } = await connectRabbitmq()
    process.once('SIGINT', () => {
      channel.close().catch((e) => console.error)
      conn.close().catch((e) => console.error)
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
    console.error(
      `${new Date().toISOString()} [marketing] error init - ${message}`,
    )
  }
}
void start()
