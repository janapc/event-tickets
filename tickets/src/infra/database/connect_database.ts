import mongoose from 'mongoose'

export async function connectDatabase(): Promise<void> {
  await mongoose.connect(process.env.MONGO_URI)
}
