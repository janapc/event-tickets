import 'dotenv/config'
import * as metrics from '@infra/metrics'
import * as api from '@infra/api'
import * as database from '@infra/database/database'
import { logger } from '@infra/logger'

if (process.env.NODE_ENV !== 'development') {
  import('@infra/trace')
  metrics.init()
}

process.once('SIGINT', (): void => {
  database.close()
  api.close()
})

async function init(): Promise<void> {
  try {
    await database.init()
    await api.init()
  } catch (error: unknown) {
    if (error instanceof Error) {
      logger.error(`init error ${String(error.message)}`)
    }
  }
}
void init()
