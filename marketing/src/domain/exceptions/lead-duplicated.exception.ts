export class LeadDuplicatedException extends Error {
  constructor() {
    super(`Duplicate lead`);
    this.name = 'LeadDuplicatedException';
  }
}
