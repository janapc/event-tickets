import { Inject, Injectable, Logger } from '@nestjs/common';
import { IProducer } from './message_queue';
import { ClientProxy } from '@nestjs/microservices';
import { lastValueFrom } from 'rxjs';

@Injectable()
export class ProducerRabbitMq implements IProducer {
  private readonly logger = new Logger(ProducerRabbitMq.name);

  constructor(
    @Inject('QUEUE_PROCESS') private readonly clientProcess: ClientProxy,
    @Inject('QUEUE_SUCCESS') private readonly clientSuccess: ClientProxy,
  ) {
    this.clientProcess.connect();
    this.clientSuccess.connect();
  }

  async addToQueue(message: string, queueName: string) {
    try {
      if (process.env.QUEUE_PROCESS_PAYMENT === queueName) {
        await lastValueFrom(this.clientProcess.emit(queueName, message));
      } else {
        await lastValueFrom(
          this.clientSuccess.emit(queueName, JSON.parse(message)),
        );
      }
      this.logger.log(`message sent to queue: ${queueName}`);
    } catch (error) {
      this.logger.error(
        `Problem with RABBITMQ in ${queueName} error: ${JSON.stringify(error)}`,
      );
    }
  }
}
