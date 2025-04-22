import {
  BatchSpanProcessor,
  SimpleSpanProcessor,
} from '@opentelemetry/sdk-trace-base';
import { NodeSDK } from '@opentelemetry/sdk-node';
import * as process from 'process';
import { HttpInstrumentation } from '@opentelemetry/instrumentation-http';
import { ExpressInstrumentation } from '@opentelemetry/instrumentation-express';
import { NestInstrumentation } from '@opentelemetry/instrumentation-nestjs-core';
import { MongooseInstrumentation } from '@opentelemetry/instrumentation-mongoose';
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-http';
import { Logger } from '@nestjs/common';

const logger = new Logger('Tracing');

const traceExporter = new OTLPTraceExporter();

const spanProcessor =
  process.env.NODE_ENV === `development`
    ? new SimpleSpanProcessor(traceExporter)
    : new BatchSpanProcessor(traceExporter);

const otelSDK = new NodeSDK({
  spanProcessor,
  instrumentations: [
    new HttpInstrumentation(),
    new ExpressInstrumentation(),
    new NestInstrumentation(),
    new MongooseInstrumentation(),
  ],
});

process.on('SIGTERM', () => {
  otelSDK
    .shutdown()
    .then(
      () => logger.log('SDK shut down successfully'),
      (err) => logger.log('Error shutting down SDK', err),
    )
    .finally(() => process.exit(0));
});

export function start() {
  try {
    otelSDK.start();
    logger.log('OpenTelemetry initialized');
  } catch (err) {
    logger.error('OpenTelemetry init failed', err);
  }
}
