import { InjectModel } from '@nestjs/mongoose';
import { User } from '@domain/user.entity';
import { UserAbstractRepository } from '@domain/user-abstract.repository';
import { UserModel, UserDocument } from './user.schema';
import { Injectable } from '@nestjs/common';
import { Model } from 'mongoose';
import { UserNotFoundException } from '@domain/exceptions/user-not-found.exception';

@Injectable()
export class UserRepository implements UserAbstractRepository {
  constructor(
    @InjectModel(UserModel.name) private userModel: Model<UserDocument>,
  ) {}
  async create(user: User): Promise<User> {
    const createdUser = await this.userModel.create(user);
    const result = new User({
      id: createdUser._id.toString(),
      email: createdUser.email,
      password: createdUser.password,
      role: createdUser.role,
    });
    return result;
  }
  async findByEmail(email: string): Promise<User | null> {
    const foundUser = await this.userModel.findOne({ email });
    if (!foundUser) return null;
    return new User({
      id: foundUser._id.toString(),
      email: foundUser.email,
      password: foundUser.password,
      role: foundUser.role,
    });
  }
  async delete(id: string): Promise<void> {
    const deletedUser = await this.userModel.deleteOne({ _id: id });
    if (deletedUser.deletedCount === 0) throw new UserNotFoundException(id);
  }
}
