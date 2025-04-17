export {}

declare global {
  namespace NodeJS {
    interface ProcessEnv {
      [key: string]: string | undefined
      DATABASE_URL: string
      QUEUE_CLIENT_CREATED: string
      PORT: number
      RABBITMQ_URL: string
      NODE_ENV: string
    }
  }
}
