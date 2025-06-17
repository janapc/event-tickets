import { MiddlewareConsumer, Module, NestModule } from '@nestjs/common';
import { MongooseModule } from '@nestjs/mongoose';
import { LeadModel, LeadSchema } from '@infra/database/lead/lead.schema';
import { LeadController } from './lead.controller';
import { CreateLeadHandler } from 'src/commands/create-lead/create-lead.handler';
import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { LeadRepository } from '@infra/database/lead/lead.repository';
import { GetLeadByEmailHandler } from '@queries/get-lead-by-email/get-lead-by-email.handler';
import { GetLeadsHandler } from '@queries/get-leads/get-leads.handler';
import { ProcessCreatedClientHandler } from '@commands/process-created-client/process-created-client.handler';
import { MetricsService } from '@infra/telemetry/metrics';
import { MetricsMiddleware } from '../middlewares/metrics.middleware';

@Module({
  imports: [
    MongooseModule.forFeature([{ name: LeadModel.name, schema: LeadSchema }]),
  ],
  controllers: [LeadController],
  providers: [
    CreateLeadHandler,
    GetLeadByEmailHandler,
    GetLeadsHandler,
    ProcessCreatedClientHandler,
    {
      provide: LeadAbstractRepository,
      useClass: LeadRepository,
    },
    MetricsService,
  ],
})
export class ControllerModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(MetricsMiddleware).forRoutes(LeadController);
  }
}
