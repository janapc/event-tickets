import bcrypt from 'bcrypt'

export type USERROLES = 'ADMIN' | 'PUBLIC'

export interface IUser {
  id?: string
  email: string
  password: string
  role: USERROLES
}

export class User {
  id?: string
  email: string
  password: string
  role: USERROLES

  constructor(email: string, password: string, role: string) {
    this.#valid(email, password, role)
    const hashPassword = this.#generatePassword(password)
    this.email = email
    this.password = hashPassword
    this.role = role as USERROLES
  }

  #valid(email: string, password: string, role: string): void {
    if (!email || !password || !role) {
      throw new Error(`the email, role and the password fields are mandatory`)
    }
    if (!/^[\w-.]+@([\w-]+\.)+[\w-]{2,4}$/g.test(email)) {
      throw new Error(`the email format is invalid`)
    }
    if (password.length < 8) {
      throw new Error(`the password field must be more than 8 characters`)
    }
    if (!['ADMIN', 'PUBLIC'].includes(role)) {
      throw new Error(`the role field must be PUBLIC or ADMIN`)
    }
  }

  #generatePassword(password: string): string {
    return bcrypt.hashSync(password, 12)
  }
}
