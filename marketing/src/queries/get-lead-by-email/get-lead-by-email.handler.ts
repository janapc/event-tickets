import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { IQueryHandler, QueryHandler } from '@nestjs/cqrs';
import { Lead } from '@domain/lead';
import { Inject } from '@nestjs/common';
import { GetLeadByEmailQuery } from './get-lead-by-email.query';

@QueryHandler(GetLeadByEmailQuery)
export class GetLeadByEmailHandler
  implements IQueryHandler<GetLeadByEmailQuery>
{
  constructor(
    @Inject(LeadAbstractRepository)
    private readonly leadRepository: LeadAbstractRepository,
  ) {}

  async execute(query: GetLeadByEmailQuery): Promise<Lead> {
    return await this.leadRepository.getByEmail(query.email);
  }
}
