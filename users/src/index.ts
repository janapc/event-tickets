import 'dotenv/config'
import mongoose from 'mongoose'
import { server } from '@infra/api'
import { connectDatabase } from '@infra/database/connect_database'

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
  mongoose.disconnect().catch((error) => {
    console.error(
      `${new Date().toISOString()} [users] error - ${getErrorMessage(error)}`,
    )
  })
})

function getErrorMessage(error: unknown): string {
  if (error instanceof Error) return error.message
  return String(error)
}
