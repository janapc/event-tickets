import { type ITicketRepository } from '@domain/repository'
import { Ticket } from '@domain/ticket'
import { type IMail } from '@infra/mail/mail'
import { formatTemplate } from '@infra/mail/template'

interface InputProcessMessageTicket {
  name: string
  email: string
  eventId: string
  eventName: string
  eventDescription: string
  eventImageUrl: string
}

export class ProcessMessageTicket {
  constructor(
    private readonly ticketRepository: ITicketRepository,
    private readonly mail: IMail,
  ) {}

  async execute(message: string): Promise<void> {
    const input = JSON.parse(message) as InputProcessMessageTicket
    const ticket = new Ticket(input.email, input.eventId)
    await this.ticketRepository.save(ticket)
    const template = formatTemplate(
      input.name,
      ticket.passport,
      input.eventDescription,
      input.eventName,
      input.eventImageUrl,
    )
    await this.mail.sendMail(input.email, 'Your ticket is here =)', template)
  }
}
