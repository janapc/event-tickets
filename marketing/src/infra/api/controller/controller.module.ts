import { Module } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { LeadModel, LeadSchema } from '@infra/database/lead/lead.schema';
import { LeadController } from './lead.controller';
import { CreateLeadHandler } from 'src/commands/create-lead/create-lead.handler';
import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { LeadRepository } from '@infra/database/lead/lead.repository';
import { GetLeadByEmailHandler } from '@queries/get-lead-by-email/get-lead-by-email.handler';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: LeadModel.name, schema: LeadSchema }]),
  ],
  controllers: [LeadController],
  providers: [
    CreateLeadHandler,
    GetLeadByEmailHandler,
    {
      provide: LeadAbstractRepository,
      useClass: LeadRepository,
    },
  ],
})
export class ControllerModule {}
