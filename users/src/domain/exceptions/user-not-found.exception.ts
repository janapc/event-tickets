export class UserNotFoundException extends Error {
  constructor(emailOrId: string) {
    super(`User with identifier "${emailOrId}" not found`);
    this.name = 'UserNotFoundException';
  }
}
