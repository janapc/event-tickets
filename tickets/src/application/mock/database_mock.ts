import { type ITicketRepository } from '@domain/repository'
import { type Ticket } from '@domain/ticket'

export class DatabaseMock implements ITicketRepository {
  public tickets: Ticket[] = []
  async save(ticket: Ticket): Promise<void> {
    this.tickets.push(ticket)
  }
}
