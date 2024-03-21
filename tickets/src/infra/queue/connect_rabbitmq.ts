import amqplib from 'amqplib'

interface OutputQueue {
  channel: amqplib.Channel
  conn: amqplib.Connection
}

export async function connectRabbitmq(): Promise<OutputQueue> {
  const conn = await amqplib.connect(process.env.RABBITMQ_URL)
  const channel = await conn.createChannel()
  await channel.assertQueue(process.env.QUEUE_NAME)
  return { channel, conn }
}
