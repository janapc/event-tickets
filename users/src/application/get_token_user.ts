import bcrypt from 'bcrypt'
import { sign } from 'jsonwebtoken'
import { type IUserRepository } from '@domain/repository'

export class GetTokenUser {
  constructor(private readonly repository: IUserRepository) {}

  async execute(email: string, password: string): Promise<string> {
    const user = await this.repository.findByEmail(email)
    if (!user) {
      throw new Error('we can not find the user.Please try again')
    }
    const isValidPassword = bcrypt.compareSync(password, String(user.password))
    if (!isValidPassword) {
      throw new Error('your email or password is incorrect.Please try again')
    }
    const token = this.#generateToken(String(user.id), user.role)
    return token
  }

  #generateToken(id: string, role: string): string {
    return sign({ role, id }, process.env.JWT_SECRET, {
      expiresIn: process.env.JWT_EXPIRES_IN,
    })
  }
}
