import { type ILeadRepository } from '@domain/repository'

interface OuputGetByEmail {
  id: number
  email: string
  language: string
  converted: boolean
}

export class GetByEmail {
  constructor(private readonly repository: ILeadRepository) {}

  async execute(email: string): Promise<OuputGetByEmail> {
    const result = await this.repository.getByEmail(email)
    if (!result) {
      throw new Error('lead is not found')
    }
    return {
      id: Number(result.id),
      email: result.email,
      language: result.language,
      converted: result.converted,
    }
  }
}
