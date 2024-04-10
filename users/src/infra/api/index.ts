/* eslint-disable @typescript-eslint/no-floating-promises */
import Fastify, { type FastifyError } from 'fastify'
import cors from '@fastify/cors'
import fastifySwagger from '@fastify/swagger'
import fastifySwaggerUi from '@fastify/swagger-ui'
import { swaggerConfig } from '@infra/docs/swagger'
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
  const err = errorsMap.get(error.message)
  const message = err ? error.message : 'internal server error'
  const statusCode = err ?? 500
  await reply.status(statusCode).send({ message })
})

fastify.register(fastifySwagger, swaggerConfig)

fastify.register(fastifySwaggerUi, {
  routePrefix: '/users/docs',
  uiConfig: {
    docExpansion: 'full',
    deepLinking: false,
  },
  uiHooks: {
    onRequest: function (_request: any, _reply: any, next: any) {
      next()
    },
    preHandler: function (_request: any, _reply: any, next: any) {
      next()
    },
  },
  staticCSP: true,
  transformStaticCSP: (header: any) => header,
  transformSpecification: (swaggerObject: any) => {
    return swaggerObject
  },
  transformSpecificationClone: true,
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
