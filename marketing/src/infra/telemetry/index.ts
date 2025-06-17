import { NodeSDK } from '@opentelemetry/sdk-node';
import * as process from 'process';
import { HttpInstrumentation } from '@opentelemetry/instrumentation-http';
import { ExpressInstrumentation } from '@opentelemetry/instrumentation-express';
import { NestInstrumentation } from '@opentelemetry/instrumentation-nestjs-core';
import { MongooseInstrumentation } from '@opentelemetry/instrumentation-mongoose';
import { Logger } from '@nestjs/common';
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-proto';
import { ATTR_SERVICE_NAME } from '@opentelemetry/semantic-conventions';
import { resourceFromAttributes } from '@opentelemetry/resources';
import { KafkaJsInstrumentation } from '@opentelemetry/instrumentation-kafkajs';
import { PeriodicExportingMetricReader } from '@opentelemetry/sdk-metrics';
import { OTLPMetricExporter } from '@opentelemetry/exporter-metrics-otlp-proto';

const logger = new Logger('Telemetry');

const resource = resourceFromAttributes({
  [ATTR_SERVICE_NAME]: 'marketing-service',
});

const otelSDK = new NodeSDK({
  traceExporter: new OTLPTraceExporter(),
  metricReader: new PeriodicExportingMetricReader({
    exporter: new OTLPMetricExporter(),
    exportIntervalMillis: 1000,
  }),
  instrumentations: [
    new NestInstrumentation(),
    new HttpInstrumentation(),
    new ExpressInstrumentation(),
    new MongooseInstrumentation(),
    new KafkaJsInstrumentation(),
  ],
  resource,
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

otelSDK.start();
