const swaggerConfig = {
  openapi: {
    info: {
      title: 'User API',
      description: 'User API Documentation',
      version: '1.0.0',
    },
    servers: [
      {
        url: 'http://localhost:' + process.env.PORT,
      },
    ],
  },
}

const swaggerUiConfig = {
  routePrefix: '/users/docs',
  uiConfig: {
    docExpansion: 'full',
    deepLinking: false,
  },
  uiHooks: {
    onRequest: function (_request, _reply, next) {
      next()
    },
    preHandler: function (_request, _reply, next) {
      next()
    },
  },
  staticCSP: true,
  transformStaticCSP: (header) => header,
  transformSpecification: (swaggerObject) => {
    return swaggerObject
  },
  transformSpecificationClone: true,
}
export { swaggerUiConfig, swaggerConfig }
