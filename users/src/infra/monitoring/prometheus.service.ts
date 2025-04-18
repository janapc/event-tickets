import { Injectable } from '@nestjs/common';
import * as client from 'prom-client';

@Injectable()
export class PrometheusService {
  private readonly httpRequestDurationMicroseconds: client.Histogram<string>;
  private readonly httpRequestCountByStatus: client.Counter<string>;
  private readonly userCreatedCount: client.Counter<string>;

  constructor() {
    client.collectDefaultMetrics();

    this.httpRequestDurationMicroseconds = new client.Histogram({
      name: 'http_request_duration_ms',
      help: 'Duration of HTTP requests in ms',
      labelNames: ['method', 'route', 'status_code'],
    });

    this.httpRequestCountByStatus = new client.Counter({
      name: 'http_requests_by_status',
      help: 'Total number of HTTP requests by status code',
      labelNames: ['status_code'],
    });

    this.userCreatedCount = new client.Counter({
      name: 'user_created_count',
      help: 'Total number of users created',
      labelNames: ['status_code'],
    });
  }

  startTimer() {
    return this.httpRequestDurationMicroseconds.startTimer();
  }

  incrementRequestCountByStatus(labels: Record<string, string>) {
    this.httpRequestCountByStatus.inc(labels);
  }

  incrementUserCreatedCount(labels: Record<string, string>) {
    this.userCreatedCount.inc(labels);
  }

  async getMetrics(): Promise<string> {
    return await client.register.metrics();
  }
}
