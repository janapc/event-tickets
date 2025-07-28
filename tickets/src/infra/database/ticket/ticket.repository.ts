import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { TicketAbstractRepository } from '@domain/ticket-abstract.repository';
import { TicketDocument, TicketModel } from './ticket.schema';
import { Ticket } from '@domain/ticket.entity';

@Injectable()
export class TicketRepository implements TicketAbstractRepository {
  constructor(
    @InjectModel(TicketModel.name) private ticketModel: Model<TicketDocument>,
  ) {}

  async save(ticket: Ticket): Promise<Ticket> {
    const newTicket = await this.ticketModel.create(ticket);
    return new Ticket({
      id: newTicket._id.toString(),
      eventId: newTicket.eventId,
      email: newTicket.email,
      passport: newTicket.passport,
    });
  }
}
