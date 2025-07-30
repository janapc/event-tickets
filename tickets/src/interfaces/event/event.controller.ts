import { CreateTicketCommand } from '@commands/create-ticket/create-ticket.command';
import { Controller } from '@nestjs/common';
import { CommandBus } from '@nestjs/cqrs';
import { MessagePattern, Payload } from '@nestjs/microservices';
import { SendTicketDto } from './dto/send-ticket.dto';

@Controller()
export class EventController {
  constructor(private readonly commandBus: CommandBus) {}

  @MessagePattern('SEND_TICKET_TOPIC')
  async sendTicket(@Payload() message: SendTicketDto): Promise<void> {
    await this.commandBus.execute(
      new CreateTicketCommand({
        messageId: message.messageId,
        name: message.name,
        email: message.email,
        eventId: message.eventId,
        eventName: message.eventName,
        eventDescription: message.eventDescription,
        eventImageUrl: message.eventImageUrl,
        language: message.language,
      }),
    );
  }
}
