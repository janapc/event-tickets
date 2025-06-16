import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { Inject, Logger } from '@nestjs/common';
import { ProcessCreatedClientCommand } from './process-created-client.query';

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
    await this.leadRepository.converted(command.email);
    this.logger.log('message processed:', { messageId: command.messageId });
  }
}
