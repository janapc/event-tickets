import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { Inject, Logger } from '@nestjs/common';
import { ProcessCreatedClientCommand } from './process-created-client.query';
import { Lead } from '@domain/lead';

@CommandHandler(ProcessCreatedClientCommand)
export class ProcessCreatedClientHandler
  implements ICommandHandler<ProcessCreatedClientCommand>
{
  private readonly logger = new Logger(ProcessCreatedClientHandler.name);
  constructor(
    @Inject(LeadAbstractRepository)
    private readonly leadRepository: LeadAbstractRepository,
  ) {}

  async execute(command: ProcessCreatedClientCommand): Promise<void> {
    this.logger.log('processing message:', {
      messageId: command.messageId,
    });
    const lead = await this.leadRepository.getByEmail(command.email);
    if (lead) {
      await this.leadRepository.converted(command.email);
      this.logger.log('message processed:', {
        messageId: command.messageId,
        leadId: lead.id,
      });
      return;
    }
    const newLead = await this.leadRepository.save(
      new Lead({
        email: command.email,
        converted: true,
      }),
    );
    this.logger.log('lead created and message processed:', {
      messageId: command.messageId,
      leadId: newLead.id,
    });
    return;
  }
}
