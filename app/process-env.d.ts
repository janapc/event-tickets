export {}

declare global {
  namespace NodeJS {
    interface ProcessEnv {
      [key: string]: string | undefined
      NEXT_PUBLIC_API_LEAD: string
      NEXT_PUBLIC_API_USERS: string
      NEXT_PUBLIC_EMAIL_USERS: string
      NEXT_PUBLIC_PASSWORD_USERS: string
      NEXT_PUBLIC_API_EVENTS: string
      NEXT_PUBLIC_API_PAYMENT: string
    }
  }
}
