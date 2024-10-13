import { logger } from '@infra/logger'
import mongoose from 'mongoose'

export async function init(): Promise<void> {
  await mongoose.connect(process.env.MONGO_URI)
}

export function close(): void {
  mongoose.connection.close().catch((error) => {
    logger.error(`database close connection error ${String(error)}`)
  })
}
