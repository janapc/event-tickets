import { Injectable } from '@nestjs/common';
import { Counter, Histogram, metrics } from '@opentelemetry/api';

@Injectable()
export class MetricsService {
  private leadCreatedCounter: Counter;
  private httpRequestCounter: Counter;
  private httpRequestHistogram: Histogram;

  constructor() {
    const meter = metrics.getMeter('marketing-service');
    this.leadCreatedCounter = meter.createCounter('lead_created_total', {
      description: 'Total number of leads created',
    });

    this.httpRequestCounter = meter.createCounter('http_request_total', {
      description: 'Total number of HTTP requests',
    });

    this.httpRequestHistogram = meter.createHistogram(
      'http_request_duration_seconds',
      {
        description: 'Duration of HTTP requests',
      },
    );
  }

  incrementLeadCreated(): void {
    this.leadCreatedCounter.add(1);
  }

  incrementHttpRequest(method: string, path: string, statusCode: number): void {
    this.httpRequestCounter.add(1, { method, path, status_code: statusCode });
  }

  recordHttpRequestDuration(
    duration: number,
    method: string,
    statusCode: number,
    path: string,
  ): void {
    this.httpRequestHistogram.record(duration, {
      method,
      status_code: statusCode,
      path,
    });
  }
}
