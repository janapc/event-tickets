import { type Lead } from './lead';

export abstract class LeadAbstractRepository {
  abstract save: (lead: Lead) => Promise<Lead>;
  abstract getByEmail: (email: string) => Promise<Lead>;
  abstract getAll: () => Promise<Lead[]>;
  abstract converted: (email: string) => Promise<void>;
}
