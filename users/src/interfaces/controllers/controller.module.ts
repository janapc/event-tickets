import { MiddlewareConsumer, Module, NestModule } from '@nestjs/common';
import { UserController } from './user.controller';
import { UserAbstractRepository } from '@domain/user-abstract.repository';
import { UserRepository } from '@infra/database/user/user.repository';
import { MongooseModule } from '@nestjs/mongoose';
import { UserModel, UserSchema } from '@infra/database/user/user.schema';
import {
  RemoveUserHandler,
  RegisterUserHandler,
  GenerateUserTokenHandler,
} from '@application/commands';
import { HealthController } from './health.controller';
import { TerminusModule } from '@nestjs/terminus';
import { HttpModule } from '@nestjs/axios';
import { MetricsService } from '@infra/metrics/metrics.service';
import { MetricsMiddleware } from '@interfaces/middleware/metrics.middleware';
@Module({
  imports: [
    MongooseModule.forFeature([{ name: UserModel.name, schema: UserSchema }]),
    TerminusModule,
    HttpModule,
  ],
  controllers: [UserController, HealthController],
  providers: [
    RegisterUserHandler,
    RemoveUserHandler,
    GenerateUserTokenHandler,
    {
      provide: UserAbstractRepository,
      useClass: UserRepository,
    },
    MetricsService,
  ],
})
export class ControllerModule implements NestModule {
  configure(consumer: MiddlewareConsumer) {
    consumer.apply(MetricsMiddleware).forRoutes(UserController);
  }
}
