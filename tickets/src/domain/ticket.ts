import { v4 as uuid } from 'uuid'

export class Ticket {
  id?: string
  email: string
  eventId: string
  passport: string

  constructor(email: string, eventId: string) {
    this.email = email
    this.eventId = eventId
    this.passport = uuid()
  }
}
