import 'dotenv/config'
import { server } from '@infra/api'
import {
  closeDatabase,
  initDatabase,
} from '@infra/database/database_connection'

process.once('SIGINT', () => {
  closeDatabase().catch((error) => {
    console.error(
      `${new Date().toISOString()} [users] close database connection - ${String(error.message)}`,
    )
  })
})

async function init(): Promise<void> {
  try {
    await initDatabase()
    await server()
  } catch (error: any) {
    console.error(
      `${new Date().toISOString()} [users] error init - ${String(error.message)}`,
    )
  }
}
void init()
