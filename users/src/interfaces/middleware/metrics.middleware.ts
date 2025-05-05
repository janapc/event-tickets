import { MetricsService } from '@infra/metrics/metrics.service';
import { Injectable, NestMiddleware } from '@nestjs/common';
import { NextFunction, Request, Response } from 'express';

@Injectable()
export class MetricsMiddleware implements NestMiddleware {
  constructor(private readonly metricsService: MetricsService) {}

  use(req: Request, res: Response, next: NextFunction) {
    const path: string = req.originalUrl.replace('v1/', '');
    const start = Date.now();

    res.on('finish', () => {
      const duration = (Date.now() - start) / 1000;
      const method = req.method;
      const statusCode = res.statusCode;
      this.metricsService.incrementHttpRequest(method, path, statusCode);
      this.metricsService.recordHttpRequestDuration(
        duration,
        method,
        statusCode,
        path,
      );
    });
    next();
  }
}
