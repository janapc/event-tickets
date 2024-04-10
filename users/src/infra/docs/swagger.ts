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

export { swaggerConfig }
