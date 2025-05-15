import { context, trace } from '@opentelemetry/api';
import { WinstonModule } from 'nest-winston';
import * as winston from 'winston';

export const traceContextFormat = winston.format((info) => {
  const span = trace.getSpan(context.active());
  const spanContext = span?.spanContext();
  if (spanContext) {
    info.traceId = spanContext.traceId;
    info.spanId = spanContext.spanId;
  }
  return info;
});

export const logger = WinstonModule.createLogger({
  format: winston.format.combine(
    winston.format.timestamp(),
    traceContextFormat(),
    winston.format.json(),
  ),
  transports: [
    new winston.transports.File({
      filename: 'logs/app.log',
      maxsize: 5_000_000,
      maxFiles: 5,
    }),
    new winston.transports.Console(),
  ],
});
