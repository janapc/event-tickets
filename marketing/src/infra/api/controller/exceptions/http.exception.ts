import { LeadDuplicatedException } from '@domain/exceptions/lead-duplicated.exception';
import { LeadNotFoundException } from '@domain/exceptions/lead-not-found.exception';
import {
  ExceptionFilter,
  Catch,
  ArgumentsHost,
  HttpException,
  HttpStatus,
  Logger,
} from '@nestjs/common';
import { Response, Request } from 'express';

export const ErrorHttpStatusMap = new Map<string, HttpStatus>([
  [LeadDuplicatedException.name, HttpStatus.CONFLICT],
  [LeadNotFoundException.name, HttpStatus.NOT_FOUND],
]);

@Catch()
export class HttpExceptionFilter implements ExceptionFilter {
  private readonly logger = new Logger(HttpExceptionFilter.name);
  catch(exception: HttpException, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();
    let status = HttpStatus.INTERNAL_SERVER_ERROR;
    let error: string = 'Internal server error';
    if (exception instanceof HttpException) {
      status = exception.getStatus();
      const responseBody = exception.getResponse();
      if (typeof responseBody === 'string') {
        error = responseBody;
      } else if (typeof responseBody === 'object' && responseBody !== null) {
        error =
          (responseBody['message'] as string) ||
          (responseBody['error'] as string) ||
          'Internal server error';
      }
    }
    if (ErrorHttpStatusMap.has(exception.name)) {
      status = ErrorHttpStatusMap.get(exception.name)!;
      error = exception.message;
    }
    this.logger.error(
      `${request.method} - ${request.url} - ${status} - ${exception.message}`,
    );
    response.status(status).json({
      statusCode: status,
      error,
      timestamp: new Date().toISOString(),
    });
  }
}
