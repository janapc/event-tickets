import { Module } from '@nestjs/common';
import { ApiController } from './api.controller';
import { ConfigModule } from '@nestjs/config';
import { RegisterPayment } from 'src/application/register_payment';
import { DatabaseModule } from '../database/database.module';
import { MessageQueueModule } from '../message-queue/message_queue.module';
import { ProcessPayment } from '@application/process_payment';
import { MailServiceModule } from '@infra/mail-service/mail_service.module';

@Module({
  imports: [
    ConfigModule.forRoot(),
    DatabaseModule,
    MessageQueueModule,
    MailServiceModule,
  ],
  controllers: [ApiController],
  providers: [RegisterPayment, ProcessPayment],
})
export class ApiModule {}
