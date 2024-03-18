import 'dotenv/config'
import { server } from '@infra/api'

async function init(): Promise<void> {
  try {
    server()
  } catch (error) {
    let message = 'internal server error'
    if (error instanceof Error) message = error.message
    console.error(
      `${new Date().toISOString()} [marketing] error init - ${message}`,
    )
  }
}
void init()
