import { type Model } from 'mongoose'
import { type ITicketRepository } from '@domain/repository'
import { type Ticket } from '@domain/ticket'

export class TicketRepository implements ITicketRepository {
  constructor(private readonly TicketModel: Model<Ticket>) {}

  async save(ticket: Ticket): Promise<void> {
    const data = new this.TicketModel(ticket)
    await data.save()
  }
}
