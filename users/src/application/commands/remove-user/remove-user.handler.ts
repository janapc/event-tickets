import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { UserAbstractRepository } from '@domain/user-abstract.repository';
import { Logger } from '@nestjs/common';
import { RemoveUserCommand } from '@application/commands/remove-user/remove-user.command';

@CommandHandler(RemoveUserCommand)
export class RemoveUserHandler implements ICommandHandler<RemoveUserCommand> {
  private readonly logger = new Logger(RemoveUserHandler.name);
  constructor(private readonly userRepository: UserAbstractRepository) {}

  async execute(command: RemoveUserCommand): Promise<void> {
    await this.userRepository.delete(command.id);
    this.logger.log(`User ${command.id} removed successfully`);
  }
}
