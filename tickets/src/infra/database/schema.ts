import { type Ticket } from '@domain/ticket'
import { Schema, model } from 'mongoose'

export const ticketSchema = new Schema<Ticket>({
  email: { type: String, required: true },
  eventId: { type: String, required: true },
  passport: { type: String, required: true },
})

export const ticketModel = model<Ticket>('Tickets', ticketSchema)
