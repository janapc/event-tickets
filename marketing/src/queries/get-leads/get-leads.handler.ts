import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { IQueryHandler, QueryHandler } from '@nestjs/cqrs';
import { Lead } from '@domain/lead';
import { Inject } from '@nestjs/common';
import { GetLeadsQuery } from './get-leads.query';

@QueryHandler(GetLeadsQuery)
export class GetLeadsHandler implements IQueryHandler<GetLeadsQuery> {
  constructor(
    @Inject(LeadAbstractRepository)
    private readonly leadRepository: LeadAbstractRepository,
  ) {}

  //eslint-disable-next-line @typescript-eslint/no-unused-vars
  async execute(query: GetLeadsQuery): Promise<Lead[]> {
    return await this.leadRepository.getAll();
  }
}
