/* eslint-disable @typescript-eslint/no-floating-promises */
import 'dotenv/config'
import swaggerAutogen from 'swagger-autogen'

const doc = {
  host: 'localhost:' + process.env.PORT,
  info: {
    version: 'v1.0.0',
    title: 'Marketing API',
    description: 'api to manager marketing',
  },
  servers: [
    {
      url: 'http://localhost:' + process.env.PORT,
      description: 'api to create and list leads',
    },
  ],
  definitions: {
    createLeadRequest: {
      $email: 'test@test.com',
      $converted: true,
      $language: 'pt',
    },
    lead: {
      $id: 'asd-123-asd',
      $email: 'test@test.com',
      $converted: true,
      language: 'pt',
    },
    getLeadsResponse: [
      {
        $ref: '#/definitions/lead',
      },
    ],
    getByEmailResponse: [
      {
        $ref: '#/definitions/lead',
      },
    ],
  },
}

const outputFile = './swagger_output.json'
const routes = ['../api/index.ts']

swaggerAutogen({ openapi: '3.0.0' })(outputFile, routes, doc)
