import { type ILeadRepository } from '@domain/repository'

interface OuputGetLeads {
  id: number
  email: string
  converted: boolean
}

export class GetLeads {
  constructor(private readonly repository: ILeadRepository) {}

  async execute(): Promise<OuputGetLeads[]> {
    const result = await this.repository.getAll()
    return result.map((r) => ({
      id: Number(r.id),
      email: r.email,
      converted: r.converted,
    }))
  }
}
