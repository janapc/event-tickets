export class GenerateUserTokenCommand {
  constructor(
    public readonly email: string,
    public readonly password: string,
  ) {}
}
