import Fastify, { type FastifyError } from 'fastify'
import cors from '@fastify/cors'
import fastifySwagger from '@fastify/swagger'
import fastifySwaggerUi from '@fastify/swagger-ui'
import autoload from '@fastify/autoload'
import path from 'node:path'
import { swaggerConfig } from '@infra/docs/swagger'
import { requestTotalCounter } from '@infra/metrics/prometheus_metrics'
import { errorsMap } from './error_handler'
import { logger } from '@infra/logger'

const fastify = Fastify({
  logger: true,
})

fastify.register(cors)

fastify.setErrorHandler(async (error: FastifyError, request, reply) => {
  logger.error(`handler error - ${error.message}`)
  const mappedError = errorsMap.get(error.message)
  const message = mappedError ? error.message : 'internal server error'
  const statusCode = mappedError ?? 500
  requestTotalCounter
    .labels({
      route: request.routeOptions.url,
      method: request.method,
      statusCode,
    })
    .inc()
  await reply.status(statusCode).send({ message })
})

fastify.register(fastifySwagger, swaggerConfig)
fastify.register(fastifySwaggerUi, {
  routePrefix: '/docs',
})
fastify.register(autoload, {
  dir: path.join(__dirname, 'routes'),
  options: Object.assign({}),
})

export async function init() {
  fastify.listen(
    { port: process.env.PORT, host: '0.0.0.0' },
    (err, address) => {
      if (err) {
        fastify.log.error(err)
        process.exit(1)
      }
      logger.info(`server listening at ${address}`)
    },
  )
}

export function close() {
  fastify.close().catch((error) => {
    logger.info(`api error ${String(error)}`)
  })
}
