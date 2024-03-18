import { type Lead } from '@domain/lead'
import { type ILeadRepository } from '@domain/repository'

export class DatabaseMockRepository implements ILeadRepository {
  public leads: Lead[] = []

  async save(lead: Lead): Promise<Lead> {
    const id = new Date().valueOf()
    this.leads.push({ id, ...lead })
    return { ...lead, id }
  }

  async getAll(): Promise<Lead[]> {
    return this.leads
  }

  async getByEmail(email: string): Promise<Lead | null> {
    return this.leads.find((lead) => lead.email === email) ?? null
  }

  async update(email: string, converted: boolean): Promise<void> {
    const indexLead = this.leads.findIndex((l) => l.email === email)
    if (indexLead === -1) return
    this.leads[indexLead].converted = converted
  }
}
