import 'dotenv/config'
import { server } from '@infra/api'
import { logger } from '@infra/logger/logger'

async function init(): Promise<void> {
  try {
    server()
  } catch (error) {
    let message = 'internal server error'
    if (error instanceof Error) message = error.message
    logger.error(`error init api : ${message}`)
  }
}
void init()
