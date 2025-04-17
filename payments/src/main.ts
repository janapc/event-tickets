import { NestFactory } from '@nestjs/core';
import { ApiModule } from './infra/api/api.module';
import { ValidationPipe } from '@nestjs/common';
import { Transport } from '@nestjs/microservices';
import { ConfigService } from '@nestjs/config';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';

async function bootstrap() {
  const app = await NestFactory.create(ApiModule);
  app.enableCors();
  const configService = app.get(ConfigService);

  app.connectMicroservice({
    transport: Transport.RMQ,
    options: {
      noAck: false,
      urls: [`${configService.get('QUEUE_URL')}`],
      queue: `${configService.get('QUEUE_PROCESS_PAYMENT')}`,
      queueOptions: { durable: true },
    },
  });

  await app.startAllMicroservices();
  app.useGlobalPipes(new ValidationPipe());
  const config = new DocumentBuilder()
    .setTitle('Process payments')
    .setDescription('api the process payments')
    .setVersion('1.0')
    .build();
  const document = SwaggerModule.createDocument(app, config);
  SwaggerModule.setup('payments/docs', app, document);
  await app.listen(configService.get('PORT'));
}
bootstrap();
