import { type ILeadRepository } from '@domain/repository'

interface OuputGetLeads {
  id: string
  email: string
  language: string
  converted: boolean
}

export class GetLeads {
  constructor(private readonly repository: ILeadRepository) {}

  async execute(): Promise<OuputGetLeads[]> {
    const leads = await this.repository.getAll()
    return leads.map((lead) => ({
      id: String(lead.id),
      email: lead.email,
      language: lead.language,
      converted: lead.converted,
    }))
  }
}
