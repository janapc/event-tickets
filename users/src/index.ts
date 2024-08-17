import 'dotenv/config'
import { server } from '@infra/api'
import {
  closeConnectionDatabase,
  connectDatabase,
} from '@infra/database/database'

async function init(): Promise<void> {
  try {
    await connectDatabase()
    await server()
  } catch (error: any) {
    console.error(
      `${new Date().toISOString()} [users] error init - ${String(error.message)}`,
    )
  }
}
void init()

process.once('SIGINT', () => {
  closeConnectionDatabase().catch((error) => {
    console.error(
      `${new Date().toISOString()} [users] error database - ${String(error.message)}`,
    )
  })
})
