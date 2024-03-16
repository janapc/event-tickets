import { type User, type IUser } from './user'

export interface IUserRepository {
  save: (user: User) => Promise<void>
  findByEmail: (email: string) => Promise<IUser | null>
  remove: (email: string) => Promise<void>
}
