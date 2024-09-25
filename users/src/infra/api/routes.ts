import { GetTokenUser } from '@application/get_token_user'
import { register } from 'prom-client'
import { RemoveUser } from '@application/remove_user'
import { type InputSaveUserDTO, SaveUser } from '@application/save_user'
import { userModel } from '@infra/database/schema'
import { requestTotalCounter } from '@infra/metrics/prometheus_metrics'
import { UserRepository } from '@infra/database/user_repository'
import {
  type FastifyPluginOptions,
  type FastifyInstance,
  type FastifyReply,
  type FastifyRequest,
  type FastifyError,
} from 'fastify'

export function routes(
  fastify: FastifyInstance,
  opts: FastifyPluginOptions,
  done: (err?: FastifyError) => void,
): void {
  fastify.post(
    '/users/get-token',
    {
      schema: {
        description: 'Generate access token to user',
        body: {
          type: 'object',
          properties: {
            email: { type: 'string' },
            password: { type: 'string' },
          },
        },
        response: {
          200: {
            type: 'object',
            properties: {
              token: { type: 'string' },
            },
          },
          404: {
            type: 'object',
            properties: {
              message: { type: 'string' },
            },
          },
          500: {
            type: 'object',
            properties: {
              message: { type: 'string' },
            },
          },
        },
      },
    },
    async (request, reply) => {
      const body = request.body as { email: string; password: string }
      const repository = new UserRepository(userModel)
      const getTokenUser = new GetTokenUser(repository)
      const token = await getTokenUser.execute(body.email, body.password)
      requestTotalCounter
        .labels({
          route: request.routeOptions.url,
          method: request.method,
          statusCode: 200,
        })
        .inc()
      await reply.code(200).send({ token })
    },
  )

  fastify.post(
    '/users',
    {
      schema: {
        description: 'Create a new user',
        body: {
          type: 'object',
          properties: {
            email: { type: 'string' },
            password: { type: 'string' },
            role: { type: 'string' },
          },
        },
        response: {
          201: {},
          409: {
            type: 'object',
            properties: {
              message: { type: 'string' },
            },
          },
          400: {
            type: 'object',
            properties: {
              message: { type: 'string' },
            },
          },
          500: {
            type: 'object',
            properties: {
              message: { type: 'string' },
            },
          },
        },
      },
    },
    async (request, reply): Promise<void> => {
      const body = request.body as InputSaveUserDTO
      const repository = new UserRepository(userModel)
      const saveUser = new SaveUser(repository)
      await saveUser.execute(body)
      reply.statusCode = 201
      requestTotalCounter
        .labels({
          route: request.routeOptions.url,
          method: request.method,
          statusCode: 201,
        })
        .inc()
    },
  )

  fastify.delete(
    '/users',
    {
      schema: {
        description: 'Delete a user',
        querystring: {
          type: 'object',
          properties: {
            email: { type: 'string' },
          },
        },
        response: {
          204: {},
          404: {
            type: 'object',
            properties: {
              message: { type: 'string' },
            },
          },
          500: {
            type: 'object',
            properties: {
              message: { type: 'string' },
            },
          },
        },
      },
    },
    async (request, reply) => {
      const { email } = request.query as { email: string }
      const repository = new UserRepository(userModel)
      const removeUser = new RemoveUser(repository)
      await removeUser.execute(email)
      reply.statusCode = 204
      requestTotalCounter
        .labels({
          route: request.routeOptions.url,
          method: request.method,
          statusCode: 204,
        })
        .inc()
    },
  )

  fastify.get(
    '/users/healthcheck',
    async (request: FastifyRequest, reply: FastifyReply): Promise<void> => {
      const healthcheck = {
        uptime: process.uptime(),
        message: 'OK',
        timestamp: Date.now(),
      }
      try {
        await reply.send(healthcheck)
      } catch (error) {
        healthcheck.message = error as string
        reply.statusCode = 503
      }
    },
  )

  fastify.get(
    '/metrics',
    async (request: FastifyRequest, reply: FastifyReply): Promise<void> => {
      const metrics = await register.metrics()
      await reply.header('Content-Type', register.contentType).send(metrics)
    },
  )
  done()
}
