export {}

declare global {
  namespace NodeJS {
    interface ProcessEnv {
      [key: string]: string | undefined
      MONGO_URI: string
      JWT_SECRET: string
      JWT_EXPIRES_IN: string
      PORT: number
      NODE_ENV: string
    }
  }
}
