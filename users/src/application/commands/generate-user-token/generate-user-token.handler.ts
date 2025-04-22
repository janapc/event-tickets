import { CommandHandler, ICommandHandler } from '@nestjs/cqrs';
import { GenerateUserTokenCommand } from './generate-user-token.command';
import { UserAbstractRepository } from '@domain/user-abstract.repository';
import { UserNotFoundException } from '@domain/exceptions/user-not-found.exception';
import { JwtService } from '@nestjs/jwt';
import { User } from '@domain/user.entity';
import { Logger } from '@nestjs/common';

@CommandHandler(GenerateUserTokenCommand)
export class GenerateUserTokenHandler
  implements ICommandHandler<GenerateUserTokenCommand>
{
  private readonly logger = new Logger(GenerateUserTokenHandler.name);
  constructor(
    private readonly userRepository: UserAbstractRepository,
    private readonly jwtService: JwtService,
  ) {}
  async execute(
    query: GenerateUserTokenCommand,
  ): Promise<{ token: string; expiresIn: number }> {
    const user = await this.userRepository.findByEmail(query.email);
    if (!user) {
      throw new UserNotFoundException(query.email);
    }
    if (!User.isValidPassword(query.password, user.password)) {
      throw new UserNotFoundException(query.email);
    }
    const token = await this.generateToken(user.id!, user.role, user.email);
    this.logger.log(`Token generated for user ${user.id}`);
    return { token, expiresIn: Number(process.env.JWT_EXPIRES_IN) };
  }

  private async generateToken(
    id: string,
    role: string,
    email: string,
  ): Promise<string> {
    return this.jwtService.signAsync({ role, id, email });
  }
}
