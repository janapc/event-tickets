export class LeadNotFoundException extends Error {
  field: string;

  constructor(field: string) {
    super(`Lead not found for field: ${field}`);
    this.name = 'LeadNotFoundException';
  }
}
