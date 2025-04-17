type User = {
  name: string;
  email: string;
};

export class Payment {
  id?: string;
  user: User;
  eventId: string;
  transactionId?: string;

  constructor(user: User, eventId: string, transactionId?: string) {
    this.user = user;
    this.eventId = eventId;
    if (transactionId) {
      this.transactionId = transactionId;
    }
  }
}
