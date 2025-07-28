import { Module } from '@nestjs/common';
import { TicketAbstractRepository } from '@domain/ticket-abstract.repository';
import { TicketRepository } from './ticket/ticket.repository';

@Module({
  imports: [],
  providers: [
    {
      provide: TicketAbstractRepository,
      useClass: TicketRepository,
    },
  ],
  exports: [
    {
      provide: TicketAbstractRepository,
      useClass: TicketRepository,
    },
  ],
})
export class DatabaseModule {}
