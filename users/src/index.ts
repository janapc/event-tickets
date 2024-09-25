import 'dotenv/config'
import * as metrics from '@infra/metrics'
import * as trace from '@infra/trace'
import { server } from '@infra/api'
import {
  closeDatabase,
  initDatabase,
} from '@infra/database/database_connection'

if (process.env.NODE_ENV !== 'development') {
  console.log('initialize metrics')
  trace.init()
  metrics.init()
}

process.once('SIGINT', (): void => {
  closeDatabase().catch((error) => {
    console.error(
      `${new Date().toISOString()} [users] close database connection - ${String(error.message)}`,
    )
  })
  trace.closeConnection().catch((error) => {
    console.error(
      `${new Date().toISOString()} [users] close trace connection - ${String(error.message)}`,
    )
  })
})

async function init(): Promise<void> {
  try {
    await initDatabase()
    await server()
  } catch (error: unknown) {
    if (error instanceof Error) {
      console.error(
        `${new Date().toISOString()} [users] error init - ${String(error.message)}`,
      )
    }
  }
}
void init()
