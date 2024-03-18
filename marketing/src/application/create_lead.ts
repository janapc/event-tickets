import { Lead } from '@domain/lead'
import { type ILeadRepository } from '@domain/repository'

export interface InputCreateLead {
  email: string
  converted: boolean
}

interface OuputCreateLead {
  id: number
  email: string
  converted: boolean
}

export class CreateLead {
  constructor(private readonly repository: ILeadRepository) {}

  async execute(input: InputCreateLead): Promise<OuputCreateLead> {
    const lead = new Lead(input.email, input.converted)
    const result = await this.repository.save(lead)
    return {
      id: Number(result.id),
      email: result.email,
      converted: result.converted,
    }
  }
}
