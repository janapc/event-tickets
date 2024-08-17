import mongoose from 'mongoose'

export async function connectDatabase(): Promise<void> {
  await mongoose.connect(process.env.MONGO_URI)
}

export async function closeConnectionDatabase(): Promise<void> {
  await mongoose.connection.close()
}
