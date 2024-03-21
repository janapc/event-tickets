export {}

declare global {
  namespace NodeJS {
    interface ProcessEnv {
      [key: string]: string | undefined
      MONGO_URI: string
      QUEUE_NAME: string
      RABBITMQ_URL: string
      MAIL_FROM: string
      MAIL_HOST: string
      MAIL_PORT: number
      MAIL_AUTH_USER: string
      MAIL_AUTH_PASS: string
      SERVICE: string
    }
  }
}
