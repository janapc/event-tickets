import { type ILeadRepository } from '@domain/repository'
import { logger } from '@infra/logger/logger'

interface InputConsumerMessageQueue {
  email: string
  hasClient: boolean
}

export class ConsumeMessagesQueue {
  constructor(private readonly repository: ILeadRepository) {}

  async execute(message: string): Promise<void> {
    logger.info(`receive message: ${message}`)
    const { email, hasClient: converted } = JSON.parse(
      message,
    ) as InputConsumerMessageQueue
    await this.repository.update(email, converted)
  }
}
