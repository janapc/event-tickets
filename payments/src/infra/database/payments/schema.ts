import { Prop, Schema, SchemaFactory, raw } from '@nestjs/mongoose';
import mongoose, { HydratedDocument } from 'mongoose';
import { Transaction } from '../transactions/schema';

export type PaymentDocument = HydratedDocument<Payment>;

@Schema()
export class Payment {
  @Prop({ required: true })
  eventId: string;

  @Prop(
    raw({
      name: { type: String },
      email: { type: String },
    }),
  )
  user: Record<string, string>;

  @Prop({
    type: mongoose.Schema.Types.ObjectId,
    ref: 'transactions',
  })
  transactionId: Transaction;
}

export const PaymentSchema = SchemaFactory.createForClass(Payment);
