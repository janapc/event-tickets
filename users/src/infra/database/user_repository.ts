import { type IUserRepository } from '@domain/repository'
import { type IUser, type User } from '@domain/user'
import { type Model } from 'mongoose'

export class UserRepository implements IUserRepository {
  constructor(private readonly UserModel: Model<User>) {}

  async save(user: User): Promise<void> {
    const data = new this.UserModel(user)
    await data.save()
  }

  async findByEmail(email: string): Promise<IUser | null> {
    const result = await this.UserModel.findOne({
      email,
    })
    if (!result) return null
    return {
      id: String(result._id) ?? '',
      email: result.email,
      password: result.password,
      role: result.role,
    }
  }

  async remove(email: string): Promise<IUser | null> {
    return await this.UserModel.findOneAndDelete({ email })
  }
}
