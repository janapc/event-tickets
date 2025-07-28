import { type Ticket } from './ticket.entity';

export abstract class TicketAbstractRepository {
  abstract save: (ticket: Ticket) => Promise<Ticket>;
}
