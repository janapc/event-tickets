import Fastify, { type FastifyError } from 'fastify'
import cors from '@fastify/cors'
import fastifySwagger from '@fastify/swagger'
import fastifySwaggerUi from '@fastify/swagger-ui'
import { swaggerConfig } from '@infra/docs/swagger'
import { requestTotalCounter } from '@infra/metrics/prometheus_metrics'
import { routes } from './routes'
import { errorsMap } from './error_handler'

const fastify = Fastify({
  logger: true,
})
fastify.register(cors)

fastify.setErrorHandler(async (error: FastifyError, request, reply) => {
  console.error(
    `${new Date().toISOString()} [users] error api - ${error.message}`,
  )
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
  routePrefix: '/users/docs',
})
fastify.register(routes)

export async function server(): Promise<void> {
  try {
    await fastify.listen({ port: process.env.PORT, host: '0.0.0.0' })
  } catch (err) {
    fastify.log.error(err)
    process.exit(1)
  }
}

process.once('SIGINT', () => {
  fastify.close().catch((error) => {
    console.error(
      `${new Date().toISOString()} [users] error api - ${String(error)}`,
    )
  })
})
