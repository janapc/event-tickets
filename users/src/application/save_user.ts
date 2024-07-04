import { type IUserRepository } from '@domain/repository'
import { type USERROLES, User } from '@domain/user'

export interface InputSaveUserDTO {
  email: string
  password: string
  role: USERROLES
}

export class SaveUser {
  constructor(private readonly repository: IUserRepository) {}

  async execute(input: InputSaveUserDTO): Promise<void> {
    const userExists = await this.repository.findByEmail(input.email)
    if (userExists) {
      throw new Error('email already registered')
    }
    const user = new User(input.email, input.password, input.role)
    await this.repository.save(user)
  }
}
