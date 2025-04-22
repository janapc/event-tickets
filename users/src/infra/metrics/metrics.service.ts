import { Injectable } from '@nestjs/common';
import { Counter, metrics } from '@opentelemetry/api';

@Injectable()
export class MetricsService {
  private counter: Counter;

  constructor() {
    const meter = metrics.getMeter('users-services');
    this.counter = meter.createCounter('user_created_total', {
      description: 'Total number of users created',
    });
  }

  incrementUserCreated(): void {
    this.counter.add(1);
  }
}
