import { type Lead } from './lead';

export abstract class LeadAbstractRepository {
  abstract save: (lead: Lead) => Promise<Lead>;
}
