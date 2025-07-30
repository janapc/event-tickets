import { WinstonModule } from 'nest-winston';
import * as winston from 'winston';

export const logger = WinstonModule.createLogger({
  format: winston.format.combine(
    winston.format.timestamp(),
    winston.format.json(),
  ),
  transports: [new winston.transports.Console()],
});
