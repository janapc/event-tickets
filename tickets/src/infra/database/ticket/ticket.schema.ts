import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument } from 'mongoose';

export type TicketDocument = HydratedDocument<TicketModel>;

@Schema({ collection: 'tickets' })
export class TicketModel {
  @Prop({ type: String, required: true })
  email: string;

  @Prop({ type: String, required: true })
  eventId: string;

  @Prop({ type: String, required: true })
  passport: string;

  @Prop({ default: Date.now })
  createdAt: Date;

  @Prop({ default: Date.now })
  updatedAt: Date;
}

export const TicketSchema = SchemaFactory.createForClass(TicketModel);
