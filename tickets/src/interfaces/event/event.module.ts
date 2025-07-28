import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import {
  TicketModel,
  TicketSchema,
} from '@infra/database/ticket/ticket.schema';
import { TicketRepository } from '@infra/database/ticket/ticket.repository';
import { TicketAbstractRepository } from '@domain/ticket-abstract.repository';
import { MailService } from '@infra/mail/service.mail';
import { MailAbstract } from '@domain/mail';
import { CreateTicketHandler } from '@commands/create-ticket/create-ticket.handler';
import { EventController } from './event.controller';

@Module({
  imports: [
    MongooseModule.forFeature([
      { name: TicketModel.name, schema: TicketSchema },
    ]),
  ],
  controllers: [EventController],
  providers: [
    CreateTicketHandler,
    {
      provide: TicketAbstractRepository,
      useClass: TicketRepository,
    },
    {
      provide: MailAbstract,
      useClass: MailService,
    },
  ],
})
export class EventModule {}
