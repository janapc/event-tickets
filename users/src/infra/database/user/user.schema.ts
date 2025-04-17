import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { HydratedDocument } from 'mongoose';
import { USER_ROLES } from '@domain/user.entity';

export type UserDocument = HydratedDocument<UserModel>;

@Schema({ collection: 'users' })
export class UserModel {
  @Prop({ required: true, type: String, unique: true })
  email: string;

  @Prop({ required: true, type: String })
  password: string;

  @Prop({ type: String, enum: USER_ROLES, required: true })
  role: USER_ROLES;
}

export const UserSchema = SchemaFactory.createForClass(UserModel);
