import { v4 as uuid } from 'uuid';

export interface TicketProps {
  id?: string;
  email: string;
  eventId: string;
  passport?: string;
}

export class Ticket {
  id?: string;
  email: string;
  eventId: string;
  passport: string;

  constructor(props: TicketProps) {
    this.email = props.email;
    this.eventId = props.eventId;
    this.passport = props.passport || uuid();
    this.id = props.id;
  }
}
