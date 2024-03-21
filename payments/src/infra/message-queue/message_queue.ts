export abstract class IProducer {
  abstract addToQueue(message: string, queueName: string): Promise<void>;
}
