import { Module } from '@nestjs/common';
import { CqrsModule } from '@nestjs/cqrs';
import { MongooseModule } from '@nestjs/mongoose';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { EventModule } from './interfaces/event/event.module';
import { MailerModule } from '@nestjs-modules/mailer';
import { OpenTelemetryModule } from 'nestjs-otel';

@Module({
  imports: [
    OpenTelemetryModule.forRoot({
      metrics: {
        hostMetrics: true,
        apiMetrics: {
          enable: true,
          ignoreUndefinedRoutes: false,
        },
      },
    }),
    ConfigModule.forRoot({
      isGlobal: true,
    }),
    CqrsModule.forRoot(),
    MongooseModule.forRootAsync({
      imports: [ConfigModule],
      useFactory: (configService: ConfigService) => {
        const uri = configService.get<string>('MONGODB_URL');
        return { uri };
      },
      inject: [ConfigService],
    }),
    MailerModule.forRootAsync({
      imports: [ConfigModule],
      useFactory: (configService: ConfigService) => {
        const mailHost = configService.get<string>('MAIL_HOST');
        const mailPort = configService.get<string>('MAIL_PORT');
        const mailAuthUser = configService.get<string>('MAIL_AUTH_USER');
        const mailAuthPass = configService.get<string>('MAIL_AUTH_PASS');

        return {
          transport: {
            host: mailHost,
            port: Number(mailPort),
            secure: false,
            auth: {
              user: mailAuthUser,
              pass: mailAuthPass,
            },
          },
        };
      },
      inject: [ConfigService],
    }),
    EventModule,
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
