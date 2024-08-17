import amqplib from 'amqplib'

class RabbitMQConnection {
  connection!: amqplib.Connection
  channel!: amqplib.Channel

  async init(): Promise<amqplib.Channel> {
    this.connection = await amqplib.connect(process.env.RABBITMQ_URL)
    this.channel = await this.connection.createChannel()
    await this.channel.assertQueue(process.env.QUEUE_SEND_TICKET)
    return this.channel
  }

  async closeRabbitmq(): Promise<void> {
    await this.connection.close()
    await this.channel.close()
  }
}

export const rabbitMQConnection = new RabbitMQConnection()
