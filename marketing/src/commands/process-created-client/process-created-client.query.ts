export class ProcessCreatedClientCommand {
  constructor(
    public readonly messageId: string,
    public readonly email: string,
  ) {}
}
