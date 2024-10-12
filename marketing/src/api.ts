import 'dotenv/config'
import { logger } from '@infra/logger/logger'
import * as api from '@infra/api'
import * as metrics from '@infra/metrics'
import Database from '@infra/dabatase'

process.on('SIGINT', () => {
  api.close()
  Database.getInstance().close()
})

if (process.env.NODE_ENV !== 'development') {
  import('@infra/trace')
  metrics.init()
}

function init(): void {
  try {
    api.init()
  } catch (error) {
    logger.error(`error init api ${String(error)}`)
  }
}

init()
