export class UserAlreadyExistsException extends Error {
  constructor(email: string) {
    super(`User email ${email} already exists`);
    this.name = 'UserAlreadyExistsException';
  }
}
