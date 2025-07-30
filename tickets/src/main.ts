import otelSDK from '@infra/telemetry';
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';
import { AsyncMicroserviceOptions, Transport } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';
import { logger } from '@infra/logger';

async function bootstrap() {
  // eslint-disable-next-line @typescript-eslint/await-thenable
  await otelSDK.start();
  const app = await NestFactory.createMicroservice<AsyncMicroserviceOptions>(
    AppModule,
    {
      logger,
      useFactory: (configService: ConfigService) => ({
        transport: Transport.KAFKA,
        options: {
          client: {
            brokers: [configService.getOrThrow<string>('KAFKA_BROKER')],
          },
          consumer: {
            groupId: configService.getOrThrow<string>('KAFKA_GROUP_ID'),
          },
        },
      }),
      inject: [ConfigService],
    },
  );
  app.useGlobalPipes(
    new ValidationPipe({
      transform: true,
      whitelist: true,
    }),
  );
  await app.listen();
}
bootstrap();
