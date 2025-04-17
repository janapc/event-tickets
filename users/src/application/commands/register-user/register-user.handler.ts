import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { RegisterUserCommand } from './register-user.command';
import { User } from '@domain/user.entity';
import { UserAbstractRepository } from '@domain/user-abstract.repository';
import { Logger } from '@nestjs/common';

@CommandHandler(RegisterUserCommand)
export class RegisterUserHandler
  implements ICommandHandler<RegisterUserCommand>
{
  private readonly logger = new Logger(RegisterUserHandler.name);
  constructor(private readonly userRepository: UserAbstractRepository) {}

  async execute(command: RegisterUserCommand): Promise<Omit<User, 'password'>> {
    const user = User.create({
      email: command.email,
      password: command.password,
      role: command.role,
    });
    const result = await this.userRepository.create(user);
    this.logger.log(`User ${result.email} registered successfully`);
    return {
      id: result.id,
      email: result.email,
      role: result.role,
    };
  }
}
