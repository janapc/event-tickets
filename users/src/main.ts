import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { ValidationPipe } from '@nestjs/common';
import { DocumentBuilder, SwaggerModule } from '@nestjs/swagger';
import { ConfigService } from '@nestjs/config';

async function bootstrap() {
  const app = await NestFactory.create(AppModule);
  const configService = app.get(ConfigService);
  const baseApiUrl = configService.get<string>('BASE_API_URL');
  app.useGlobalPipes(new ValidationPipe());
  const config = new DocumentBuilder()
    .setTitle('Users')
    .setDescription('Users API')
    .setVersion('1.0')
    .build();
  const documentFactory = () => SwaggerModule.createDocument(app, config);
  SwaggerModule.setup(baseApiUrl + '/api', app, documentFactory);
  app.enableCors();
  app.setGlobalPrefix(baseApiUrl ?? '');
  await app.listen(process.env.PORT ?? 3000);
}
bootstrap();
