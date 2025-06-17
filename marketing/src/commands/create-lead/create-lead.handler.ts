import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { CreateLeadCommand } from './create-lead.command';
import { Lead } from '@domain/lead';
import { Inject } from '@nestjs/common';
import { MetricsService } from '@infra/telemetry/metrics';

@CommandHandler(CreateLeadCommand)
export class CreateLeadHandler implements ICommandHandler<CreateLeadCommand> {
  constructor(
    @Inject(LeadAbstractRepository)
    private readonly leadRepository: LeadAbstractRepository,
    private readonly metricsService: MetricsService,
  ) {}

  async execute(command: CreateLeadCommand): Promise<Lead> {
    const lead = new Lead({
      email: command.email,
      converted: command.converted,
      language: command.language,
    });
    const newLead = await this.leadRepository.save(lead);
    this.metricsService.incrementLeadCreated();
    return newLead;
  }
}
