import { IProducer } from '@infra/message-queue/message_queue';

export class ProducerMockMessageQueue implements IProducer {
  public messages: string[] = [];

  async addToQueue(message: string): Promise<void> {
    this.messages.push(message);
  }
}
