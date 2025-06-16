export class CreateLeadCommand {
  constructor(
    public readonly email: string,
    public readonly converted: boolean,
    public readonly language: string,
  ) {}
}
