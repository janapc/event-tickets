import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { CreateLeadCommand } from './create-lead.command';
import { Lead } from '@domain/lead';
import { Inject } from '@nestjs/common';

@CommandHandler(CreateLeadCommand)
export class CreateLeadHandler implements ICommandHandler<CreateLeadCommand> {
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
    return this.leadRepository.save(lead);
  }
}
