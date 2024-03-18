import { type ILeadRepository } from '@domain/repository'

interface InputConsumerMessageQueue {
  email: string
  hasClient: boolean
}

export class ConsumeMessagesQueue {
  constructor(private readonly repository: ILeadRepository) {}

  async execute(message: string): Promise<void> {
    console.log(`${new Date().toISOString()} [leads] message - ${message}`)
    const { email, hasClient: converted } = JSON.parse(
      message,
    ) as InputConsumerMessageQueue
    await this.repository.update(email, converted)
  }
}
