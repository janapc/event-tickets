import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { IQueryHandler, QueryHandler } from '@nestjs/cqrs';
import { Lead } from '@domain/lead';
import { Inject } from '@nestjs/common';
import { GetLeadByEmailQuery } from './get-lead-by-email.query';
import { LeadNotFoundException } from '@domain/exceptions/lead-not-found.exception';

@QueryHandler(GetLeadByEmailQuery)
export class GetLeadByEmailHandler
  implements IQueryHandler<GetLeadByEmailQuery>
{
  constructor(
    @Inject(LeadAbstractRepository)
    private readonly leadRepository: LeadAbstractRepository,
  ) {}

  async execute(query: GetLeadByEmailQuery): Promise<Lead> {
    const lead = await this.leadRepository.getByEmail(query.email);
    if (!lead) {
      throw new LeadNotFoundException(query.email);
    }
    return lead;
  }
}
