import { type ILeadRepository } from '@domain/repository'

interface OuputGetByEmail {
  id: string
  email: string
  language: string
  converted: boolean
}

export class GetByEmail {
  constructor(private readonly repository: ILeadRepository) {}

  async execute(email: string): Promise<OuputGetByEmail> {
    const lead = await this.repository.getByEmail(email)
    if (!lead) {
      throw new Error('lead is not found')
    }
    return {
      id: String(lead.id),
      email: lead.email,
      language: lead.language,
      converted: lead.converted,
    }
  }
}
