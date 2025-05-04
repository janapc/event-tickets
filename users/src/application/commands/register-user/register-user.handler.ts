import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { RegisterUserCommand } from './register-user.command';
import { User } from '@domain/user.entity';
import { UserAbstractRepository } from '@domain/user-abstract.repository';
import { Logger } from '@nestjs/common';
import { UserAlreadyExistsException } from '@domain/exceptions/user-already-exists.exception';

@CommandHandler(RegisterUserCommand)
export class RegisterUserHandler
  implements ICommandHandler<RegisterUserCommand>
{
  private readonly logger = new Logger(RegisterUserHandler.name);
  constructor(private readonly userRepository: UserAbstractRepository) {}

  async execute(command: RegisterUserCommand): Promise<Omit<User, 'password'>> {
    const hasUser = await this.userRepository.findByEmail(command.email);
    if (hasUser) {
      throw new UserAlreadyExistsException(command.email);
    }
    const newUser = User.create({
      email: command.email,
      password: command.password,
      role: command.role,
    });
    const user = await this.userRepository.create(newUser);
    this.logger.log(`User ${user.id} registered successfully`);
    return {
      id: user.id,
      email: user.email,
      role: user.role,
    };
  }
}
