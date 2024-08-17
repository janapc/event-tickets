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
  } catch (error) {
    console.error(
      `${new Date().toISOString()} [users] error init - ${getErrorMessage(error)}`,
    )
  }
}
void init()

process.once('SIGINT', () => {
  closeConnectionDatabase().catch((error) => {
    console.error(
      `${new Date().toISOString()} [users] error database - ${getErrorMessage(error)}`,
    )
  })
})

function getErrorMessage(error: unknown): string {
  if (error instanceof Error) return error.message
  return String(error)
}
