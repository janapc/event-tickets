export class LeadDuplicatedException extends Error {
  field: string;

  constructor(field: string) {
    super(`Duplicate lead found for field: ${field}`);
    this.name = 'LeadDuplicatedException';
  }
}
