import { type Lead } from './lead'

export interface ILeadRepository {
  save: (lead: Lead) => Promise<Lead>
  update: (email: string, converted: boolean) => Promise<void>
  getAll: () => Promise<Lead[]>
  getByEmail: (email: string) => Promise<Lead | null>
}
