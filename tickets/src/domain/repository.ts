import { type Ticket } from './ticket'

export interface ITicketRepository {
  save: (ticket: Ticket) => Promise<void>
}
