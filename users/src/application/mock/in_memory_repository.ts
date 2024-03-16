import { type IUserRepository } from '@domain/repository'
import { type User } from '@domain/user'

export class InMemoryRepository implements IUserRepository {
  public users: User[] = []
  async save(user: User): Promise<void> {
    this.users.push(user)
  }

  async findByEmail(email: string): Promise<User | null> {
    const user = this.users.find((u) => u.email === email)
    return user ?? null
  }

  async remove(email: string): Promise<void> {
    const userIndex = this.users.findIndex((item) => item.email === email)
    this.users.splice(userIndex, 1)
  }
}
