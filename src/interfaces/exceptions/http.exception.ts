import {
  ExceptionFilter,
  Catch,
  ArgumentsHost,
  HttpException,
  HttpStatus,
  Logger,
} from '@nestjs/common';
import { Request, Response } from 'express';

export interface HttpExceptionResponse {
  statusCode: number;
  message: any;
  error: string;
}

@Catch(HttpException)
export class HttpExceptionFilter implements ExceptionFilter {
  private readonly logger = new Logger(HttpExceptionFilter.name);
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();
    const status = exception.getStatus() || HttpStatus.INTERNAL_SERVER_ERROR;
    const [message, originalMessage] = this.getMessage(exception, status);
    this.logger.error(
      `${request.method} - ${request.url} - ${status} - ${originalMessage}`,
    );
    response.status(status).json({
      timestamp: new Date().toISOString(),
      path: request.url,
      message: message,
    });
  }

  private getMessage(
    exception: HttpException,
    status: number,
  ): [message: string, originalMessage: string] {
    const response = exception.getResponse() as HttpExceptionResponse;
    return status === Number(HttpStatus.INTERNAL_SERVER_ERROR)
      ? ['Internal server error', String(response.message)]
      : [String(response.message), String(response.message)];
  }
}
