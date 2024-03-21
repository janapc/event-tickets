import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument } from 'mongoose';

export type TransactionDocument = HydratedDocument<Transaction>;

@Schema()
export class Transaction {
  @Prop({ required: true, length: 4 })
  cardNumber: number;

  @Prop({ required: true })
  amount: number;

  @Prop({ default: '' })
  errorMessage: string;

  @Prop({ required: true })
  status: string;
}

export const TransactionSchema = SchemaFactory.createForClass(Transaction);
