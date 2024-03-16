import { type User } from '@domain/user'
import { Schema, model } from 'mongoose'

export const userSchema = new Schema<User>({
  email: { type: String, required: true },
  password: { type: String, required: true },
  role: { type: String, required: true },
})

export const userModel = model<User>('Users', userSchema)
