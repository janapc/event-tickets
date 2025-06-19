import otelSdk from '@infra/telemetry';
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';

async function bootstrap() {
  // eslint-disable-next-line @typescript-eslint/await-thenable
  await otelSdk.start();
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
  const config = new DocumentBuilder()
    .setTitle('Marketing API')
    .setDescription('Marketing api description')
    .setVersion('1.0')
    .build();
  const documentFactory = () => SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('api', app, documentFactory);
  await app.startAllMicroservices();
  const port = configService.getOrThrow<string>('PORT');
  await app.listen(port ?? 3000);
}
bootstrap();
