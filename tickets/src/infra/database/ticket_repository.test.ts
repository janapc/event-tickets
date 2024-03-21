import { Ticket } from '@domain/ticket'
import { ticketModel } from './schema'
import { TicketRepository } from './ticket_repository'

describe('Ticket Repository', () => {
  it('should save a ticket', async () => {
    const spySave = jest
      .spyOn(ticketModel.prototype, 'save')
      .mockResolvedValue({ _id: 'asd-123' })
    const repository = new TicketRepository(ticketModel)
    await expect(
      repository.save(new Ticket('test@test.com', '123-asd')),
    ).resolves.toBeUndefined()
    expect(spySave).toHaveBeenCalledTimes(1)
  })
})
