import { NodeSDK } from '@opentelemetry/sdk-node'
import { diag, DiagConsoleLogger, DiagLogLevel } from '@opentelemetry/api'
import { HttpInstrumentation } from '@opentelemetry/instrumentation-http'
import { FastifyInstrumentation } from '@opentelemetry/instrumentation-fastify'
import { MongooseInstrumentation } from '@opentelemetry/instrumentation-mongoose'
import { MongoDBInstrumentation } from '@opentelemetry/instrumentation-mongodb'
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-proto'
diag.setLogger(new DiagConsoleLogger(), DiagLogLevel.INFO)

const exporter = new OTLPTraceExporter()
const sdk = new NodeSDK({
  traceExporter: exporter,
  instrumentations: [
    new HttpInstrumentation(),
    new FastifyInstrumentation(),
    new MongooseInstrumentation(),
    new MongoDBInstrumentation(),
  ],
})

export function init() {
  sdk.start()
}

export async function closeConnection() {
  if (sdk instanceof NodeSDK) {
    await sdk.shutdown()
  }
}
