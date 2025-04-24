import { start } from './infra/telemetry/tracing';
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { RequestMethod, ValidationPipe } from '@nestjs/common';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { ConfigService } from '@nestjs/config';

async function bootstrap() {
  start();
  const app = await NestFactory.create(AppModule);
  const configService = app.get(ConfigService);
  const prefix = configService.get<string>('PREFIX');
  const port = configService.get<number>('PORT');
  app.useGlobalPipes(new ValidationPipe());
  const config = new DocumentBuilder()
    .setTitle('Users')
    .setDescription('Users API')
    .setVersion('1.0')
    .build();
  const documentFactory = () => SwaggerModule.createDocument(app, config);
  SwaggerModule.setup(prefix + '/api', app, documentFactory);
  app.enableCors();
  app.setGlobalPrefix(prefix ?? '', {
    exclude: [{ path: 'health', method: RequestMethod.GET }],
  });
  await app.listen(port ?? 3000);
}
bootstrap();
