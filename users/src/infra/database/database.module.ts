import { Module } from '@nestjs/common';
import { UserRepository } from './user/user.repository';
import { UserAbstractRepository } from '@domain/user-abstract.repository';

@Module({
  imports: [],
  providers: [
    {
      provide: UserAbstractRepository,
      useClass: UserRepository,
    },
  ],
  exports: [
    {
      provide: UserAbstractRepository,
      useClass: UserRepository,
    },
  ],
})
export class DatabaseModule {}
