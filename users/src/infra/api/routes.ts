import { GetTokenUser } from '@application/get_token_user'
import { RemoveUser } from '@application/remove_user'
import { type InputSaveUserDTO, SaveUser } from '@application/save_user'
import { userModel } from '@infra/database/schema'
import { UserRepository } from '@infra/database/user_repository'
import { type FastifyPluginOptions, type FastifyInstance } from 'fastify'

export function routes(
  fastify: FastifyInstance,
  opts: FastifyPluginOptions,
  done: any,
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
    async (req, reply) => {
      const body = req.body as { email: string; password: string }
      const repository = new UserRepository(userModel)
      const getTokenUser = new GetTokenUser(repository)
      const token = await getTokenUser.execute(body.email, body.password)
      reply.statusCode = 200
      await reply.send({ token })
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
    },
  )
  done()
}
