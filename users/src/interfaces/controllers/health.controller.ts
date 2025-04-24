import { Controller, Get } from '@nestjs/common';
import { ConfigService } from '@nestjs/config';
import {
  HealthCheckService,
  HttpHealthIndicator,
  HealthCheck,
  MongooseHealthIndicator,
} from '@nestjs/terminus';

@Controller('health')
export class HealthController {
  constructor(
    private health: HealthCheckService,
    private http: HttpHealthIndicator,
    private mongoose: MongooseHealthIndicator,
    private configService: ConfigService,
  ) {}

  @Get()
  @HealthCheck()
  check() {
    const baseURl = this.configService.get<string>('BASE_API_URL');
    const prefix = this.configService.get<string>('PREFIX');
    return this.health.check([
      () => this.http.pingCheck('api', `${baseURl}/${prefix}/api`),
      () => this.mongoose.pingCheck('mongoose'),
    ]);
  }
}
