import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { Inject, Logger } from '@nestjs/common';
import { CreateTicketCommand } from './create-ticket.command';
import { TicketAbstractRepository } from '@domain/ticket-abstract.repository';
import { Ticket } from '@domain/ticket.entity';
import { MailAbstract } from '@domain/mail';
import { generateTicketMail } from '@infra/mail/generate-ticket.mail';

@CommandHandler(CreateTicketCommand)
export class CreateTicketHandler
  implements ICommandHandler<CreateTicketCommand>
{
  private readonly Logger = new Logger(CreateTicketHandler.name);
  constructor(
    @Inject(TicketAbstractRepository)
    private readonly ticketRepository: TicketAbstractRepository,
    @Inject(MailAbstract)
    private readonly mailService: MailAbstract,
  ) {}

  async execute(command: CreateTicketCommand): Promise<void> {
    this.Logger.log(`Creating ticket for eventId: ${command.eventId}`);
    const ticket = new Ticket({
      email: command.email,
      eventId: command.eventId,
    });
    const newTicket = await this.ticketRepository.save(ticket);
    const mail = generateTicketMail({
      name: command.name,
      passport: newTicket.passport,
      eventDescription: command.eventDescription,
      eventName: command.eventName,
      eventImageUrl: command.eventImageUrl,
      language: command.language,
    });
    await this.mailService.sendMail({
      to: command.email,
      subject: mail.subject,
      html: mail.html,
      from: process.env.MAIL_FROM!,
    });
  }
}
