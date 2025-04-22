import { User } from '@domain/user.entity';
import {
  Body,
  ConflictException,
  Controller,
  Delete,
  HttpCode,
  HttpStatus,
  InternalServerErrorException,
  NotFoundException,
  Param,
  Post,
  UseFilters,
} from '@nestjs/common';
import { CommandBus } from '@nestjs/cqrs';
import {
  RegisterUserCommand,
  GenerateUserTokenCommand,
  RemoveUserCommand,
} from '@application/commands';
import { UserAlreadyExistsException } from '@domain/exceptions/user-already-exists.exception';
import { CreateUserDto } from '../dto/create-user.dto';
import {
  ApiConflictResponse,
  ApiInternalServerErrorResponse,
  ApiCreatedResponse,
  ApiOperation,
  ApiNotFoundResponse,
  ApiOkResponse,
  ApiNoContentResponse,
} from '@nestjs/swagger';
import { UserNotFoundException } from '@domain/exceptions/user-not-found.exception';
import { GenerateUserTokenDto } from '@interfaces/dto/generate-user-token.dto';
import { HttpExceptionFilter } from '@interfaces/exceptions/http.exception';
import { MetricsService } from '@infra/metrics/metrics.service';

@Controller('users')
@UseFilters(new HttpExceptionFilter())
export class UserController {
  constructor(
    private readonly commandBus: CommandBus,
    private readonly metricsService: MetricsService,
  ) {}

  @ApiOperation({ summary: 'Register a new user' })
  @ApiCreatedResponse({ description: 'User registered successfully' })
  @ApiConflictResponse({ description: 'User already exists' })
  @ApiInternalServerErrorResponse({ description: 'Internal server error' })
  @Post()
  async register(
    @Body() createUserDto: CreateUserDto,
  ): Promise<Omit<User, 'password'>> {
    try {
      const command = new RegisterUserCommand(
        createUserDto.email,
        createUserDto.password,
        createUserDto.role,
      );
      const result: Omit<User, 'password'> =
        await this.commandBus.execute(command);
      this.metricsService.incrementUserCreated();
      return result;
    } catch (err) {
      if (err instanceof UserAlreadyExistsException) {
        throw new ConflictException(err.message);
      }
      throw new InternalServerErrorException(err);
    }
  }

  @ApiOperation({ summary: 'Remove a user' })
  @HttpCode(Number(HttpStatus.NO_CONTENT))
  @ApiNoContentResponse({ description: 'user deleted successfully' })
  @ApiNotFoundResponse({ description: 'User not found' })
  @ApiInternalServerErrorResponse({ description: 'Internal server error' })
  @Delete(':id')
  async remove(@Param('id') id: string): Promise<void> {
    try {
      const command = new RemoveUserCommand(id);
      await this.commandBus.execute(command);
    } catch (err) {
      if (err instanceof UserNotFoundException) {
        throw new NotFoundException(err.message);
      }
      throw new InternalServerErrorException(err);
    }
  }

  @ApiOperation({ summary: 'Generate a token for a user' })
  @ApiOkResponse({ description: 'Token generated successfully' })
  @ApiNotFoundResponse({ description: 'User not found' })
  @ApiInternalServerErrorResponse({ description: 'Internal server error' })
  @Post('token')
  async generateUserToken(
    @Body() generateUserTokenDto: GenerateUserTokenDto,
  ): Promise<{ token: string; expiresIn: number }> {
    try {
      const command = new GenerateUserTokenCommand(
        generateUserTokenDto.email,
        generateUserTokenDto.password,
      );
      const result: { token: string; expiresIn: number } =
        await this.commandBus.execute(command);
      return result;
    } catch (err) {
      if (err instanceof UserNotFoundException) {
        throw new NotFoundException(err.message);
      }
      throw new InternalServerErrorException(err);
    }
  }
}
