import { Lead } from '@domain/lead'
import { type ILeadRepository } from '@domain/repository'

export interface InputCreateLead {
  email: string
  language: string
  converted: boolean
}

interface OuputCreateLead {
  id: string
  email: string
  language: string
  converted: boolean
}

export class CreateLead {
  constructor(private readonly repository: ILeadRepository) {}

  async execute(input: InputCreateLead): Promise<OuputCreateLead> {
    const existsLead = await this.repository.getByEmail(input.email)
    if (existsLead) {
      return {
        id: String(existsLead.id),
        email: existsLead.email,
        language: existsLead.language,
        converted: existsLead.converted,
      }
    }
    const lead = new Lead(input.email, input.converted, input.language)
    const newLead = await this.repository.save(lead)
    return {
      id: String(newLead.id),
      email: newLead.email,
      language: newLead.language,
      converted: newLead.converted,
    }
  }
}
