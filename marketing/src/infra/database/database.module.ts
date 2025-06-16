import { Module } from '@nestjs/common';
import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { LeadRepository } from './lead/lead.repository';

@Module({
  imports: [],
  providers: [
    {
      provide: LeadAbstractRepository,
      useClass: LeadRepository,
    },
  ],
  exports: [
    {
      provide: LeadAbstractRepository,
      useClass: LeadRepository,
    },
  ],
})
export class DatabaseModule {}
