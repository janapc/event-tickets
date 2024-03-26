import { Lead } from '@domain/lead'
import { type ILeadRepository } from '@domain/repository'

export interface InputCreateLead {
  email: string
  language: string
  converted: boolean
}

interface OuputCreateLead {
  id: number
  email: string
  language: string
  converted: boolean
}

export class CreateLead {
  constructor(private readonly repository: ILeadRepository) {}

  async execute(input: InputCreateLead): Promise<OuputCreateLead> {
    const lead = new Lead(input.email, input.converted, input.language)
    const getByEmail = await this.repository.getByEmail(input.email)
    if (!getByEmail) {
      const save = await this.repository.save(lead)
      return {
        id: Number(save.id),
        email: save.email,
        language: save.language,
        converted: save.converted,
      }
    }
    return {
      id: Number(getByEmail.id),
      email: getByEmail.email,
      language: getByEmail.language,
      converted: getByEmail.converted,
    }
  }
}
