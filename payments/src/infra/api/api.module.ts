import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { TerminusModule } from '@nestjs/terminus';
import { HttpModule } from '@nestjs/axios';

import { ApiController } from './api.controller';
import { RegisterPayment } from 'src/application/register_payment';
import { DatabaseModule } from '../database/database.module';
import { MessageQueueModule } from '../message-queue/message_queue.module';
import { ProcessPayment } from '@application/process_payment';
import { MailServiceModule } from '@infra/mail-service/mail_service.module';
import { HealthController } from './health.controller';

@Module({
  imports: [
    ConfigModule.forRoot(),
    DatabaseModule,
    MessageQueueModule,
    MailServiceModule,
    TerminusModule,
    HttpModule,
  ],
  controllers: [ApiController, HealthController],
  providers: [RegisterPayment, ProcessPayment],
})
export class ApiModule {}
