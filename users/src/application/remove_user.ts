import { type IUserRepository } from '@domain/repository'

export class RemoveUser {
  constructor(private readonly repository: IUserRepository) {}

  async execute(id: string): Promise<void> {
    const response = await this.repository.remove(id)
    if (!response) {
      throw new Error('we can not find the user.Please try again')
    }
  }
}
