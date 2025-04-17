import { type ITicketRepository } from '@domain/repository'
import { Ticket } from '@domain/ticket'
import { type IMail } from '@infra/mail/mail'
import { intlMail } from '@infra/mail/intl'

interface InputProcessMessageTicket {
  name: string
  email: string
  eventId: string
  eventName: string
  eventDescription: string
  eventImageUrl: string
  language: string
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
    const body = intlMail(
      input.name,
      ticket.passport,
      input.eventDescription,
      input.eventName,
      input.eventImageUrl,
      input.language,
    )
    await this.mail.sendMail(input.email, body.subject, body.html)
  }
}
