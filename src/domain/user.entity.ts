import * as bcrypt from 'bcrypt';

export enum USER_ROLES {
  'ADMIN' = 'ADMIN',
  'PUBLIC' = 'PUBLIC',
}

type userParams = {
  id?: string;
  email: string;
  password: string;
  role: USER_ROLES;
};

export class User {
  id?: string;
  email: string;
  password: string;
  role: USER_ROLES;

  constructor(params: userParams) {
    this.email = params.email;
    this.password = params.password;
    this.role = params.role;
    this.id = params.id;
  }
  static create(params: userParams): User {
    const hashPassword = User.generatePassword(params.password);
    return new User({
      email: params.email,
      password: hashPassword,
      role: params.role,
    });
  }

  static generatePassword(password: string): string {
    const salt = bcrypt.genSaltSync(12);
    return bcrypt.hashSync(password, salt);
  }

  static isValidPassword(password: string, userPassword: string): boolean {
    return bcrypt.compareSync(password, userPassword);
  }
}
