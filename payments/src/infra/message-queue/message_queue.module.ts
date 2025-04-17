import { Module } from '@nestjs/common';
import { ProducerRabbitMq } from './producer_rabbitmq';
import { IProducer } from './message_queue';
import { ClientsModule, Transport } from '@nestjs/microservices';
import { ConfigModule, ConfigService } from '@nestjs/config';

@Module({
  imports: [
    ClientsModule.registerAsync([
      {
        name: 'QUEUE_PROCESS',
        imports: [ConfigModule],
        useFactory: (configService: ConfigService) => ({
          transport: Transport.RMQ,
          options: {
            urls: [`${configService.get('QUEUE_URL')}`],
            queue: `${configService.get('QUEUE_PROCESS_PAYMENT')}`,
            queueOptions: {
              durable: true,
            },
          },
        }),
        inject: [ConfigService],
      },
      {
        name: 'QUEUE_SUCCESS',
        imports: [ConfigModule],
        useFactory: (configService: ConfigService) => ({
          transport: Transport.RMQ,
          options: {
            urls: [`${configService.get('QUEUE_URL')}`],
            queue: `${configService.get('QUEUE_SUCCESS_PAYMENT')}`,
            queueOptions: {
              durable: true,
            },
          },
        }),
        inject: [ConfigService],
      },
    ]),
  ],
  exports: [IProducer],
  providers: [
    {
      provide: IProducer,
      useClass: ProducerRabbitMq,
    },
  ],
})
export class MessageQueueModule {}
