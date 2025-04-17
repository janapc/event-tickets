import { User } from './user.entity';

export abstract class UserAbstractRepository {
  abstract create: (user: User) => Promise<User>;
  abstract findByEmail: (email: string) => Promise<User>;
  abstract delete: (id: string) => Promise<void>;
}
