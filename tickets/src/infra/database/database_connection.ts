import mongoose from 'mongoose'

export async function initDatabase(): Promise<void> {
  await mongoose.connect(process.env.MONGO_URI)
}

export async function closeDatabase(): Promise<void> {
  await mongoose.connection.close()
}
