/* eslint-disable @typescript-eslint/no-floating-promises */
import Fastify, { type FastifyError } from 'fastify'
import { errorsMap } from './error_handler'
import fastifySwagger from '@fastify/swagger'
import fastifySwaggerUi from '@fastify/swagger-ui'
import { swaggerConfig, swaggerUiConfig } from '@infra/docs/swagger'
import { routes } from './routes'
const fastify = Fastify({
  logger: true,
})

fastify.setErrorHandler(async (error: FastifyError, request, reply) => {
  console.error(
    `${new Date().toISOString()} [users] error api - ${error.message}`,
  )
  const err = errorsMap.get(error.message)
  const message = err ? error.message : 'internal server error'
  const statusCode = err ?? 500
  await reply.status(statusCode).send({ message })
})

fastify.register(fastifySwagger, swaggerConfig)

fastify.register(fastifySwaggerUi, swaggerUiConfig)
fastify.register(routes)

export async function server(): Promise<void> {
  try {
    await fastify.listen({ port: process.env.PORT })
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
