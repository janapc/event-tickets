import { type IUserRepository } from '@domain/repository'

export class RemoveUser {
  constructor(private readonly repository: IUserRepository) {}

  async execute(id: string): Promise<void> {
    await this.repository.remove(id)
  }
}
