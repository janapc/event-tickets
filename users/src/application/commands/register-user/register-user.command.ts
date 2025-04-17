import { USER_ROLES } from '@domain/user.entity';

export class RegisterUserCommand {
  constructor(
    public readonly email: string,
    public readonly password: string,
    public readonly role: USER_ROLES,
  ) {}
}
