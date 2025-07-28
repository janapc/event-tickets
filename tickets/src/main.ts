import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const configService = app.get(ConfigService);
  app.useGlobalPipes(
    new ValidationPipe({
      transform: true,
      whitelist: true,
    }),
  );
  const kafkaBroker = configService.getOrThrow<string>('KAFKA_BROKER');
  const kafkaGroupId = configService.getOrThrow<string>('KAFKA_GROUP_ID');
  app.connectMicroservice<MicroserviceOptions>({
    transport: Transport.KAFKA,
    options: {
      client: {
        brokers: [kafkaBroker],
      },
      consumer: {
        groupId: kafkaGroupId,
      },
    },
  });
  await app.startAllMicroservices();
  const port = configService.getOrThrow<string>('PORT');
  await app.listen(port ?? 3000);
}
bootstrap();
