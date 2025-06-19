import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { CreateLeadCommand } from './create-lead.command';
import { Lead } from '@domain/lead';
import { Inject, Logger } from '@nestjs/common';

@CommandHandler(CreateLeadCommand)
export class CreateLeadHandler implements ICommandHandler<CreateLeadCommand> {
  private readonly Logger = new Logger(CreateLeadHandler.name);
  constructor(
    @Inject(LeadAbstractRepository)
    private readonly leadRepository: LeadAbstractRepository,
  ) {}

  async execute(command: CreateLeadCommand): Promise<Lead> {
    const lead = new Lead({
      email: command.email,
      converted: command.converted,
      language: command.language,
    });
    const newLead = await this.leadRepository.save(lead);
    this.Logger.log('created lead', { id: newLead.id });
    return newLead;
  }
}
