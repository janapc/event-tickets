import { HttpStatus, Injectable, NestMiddleware } from '@nestjs/common';
import { Request, Response, NextFunction } from 'express';
import { PrometheusService } from './prometheus.service';

@Injectable()
export class PrometheusMiddleware implements NestMiddleware {
  constructor(private readonly prometheusService: PrometheusService) {}

  use(req: Request, res: Response, next: NextFunction) {
    const path: string = req.baseUrl;
    const end = this.prometheusService.startTimer();

    res.on('finish', () => {
      if (
        req.method === 'POST' &&
        res.statusCode === Number(HttpStatus.CREATED) &&
        req.originalUrl === 'v1/users'
      ) {
        this.prometheusService.incrementUserCreatedCount({
          status_code: res.statusCode.toString(),
        });
      }
      this.prometheusService.incrementRequestCountByStatus({
        status_code: res.statusCode.toString(),
      });
      end({
        method: req.method,
        route: path,
        status_code: res.statusCode.toString(),
      });
    });

    next();
  }
}
