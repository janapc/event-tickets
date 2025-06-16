import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument } from 'mongoose';

export type LeadDocument = HydratedDocument<LeadModel>;

@Schema({ collection: 'leads' })
export class LeadModel {
  @Prop({ required: true, type: String, unique: true, index: true })
  email: string;

  @Prop({ required: false, type: Boolean, default: false })
  converted: boolean;

  @Prop({ type: String, required: false, default: 'pt-br' })
  language: string;

  @Prop({ default: Date.now })
  createdAt: Date;

  @Prop({ default: Date.now })
  updatedAt: Date;
}

export const LeadSchema = SchemaFactory.createForClass(LeadModel);
