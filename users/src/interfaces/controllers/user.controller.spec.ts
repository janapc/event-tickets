import { CommandBus } from '@nestjs/cqrs';
import { UserController } from './user.controller';
import { Test, TestingModule } from '@nestjs/testing';
import { CreateUserDto } from '@interfaces/dto/create-user.dto';
import { USER_ROLES } from '@domain/user.entity';
import { UserAlreadyExistsException } from '@domain/exceptions/user-already-exists.exception';
import {
  ConflictException,
  InternalServerErrorException,
  NotFoundException,
} from '@nestjs/common';
import { Types } from 'mongoose';
import { RemoveUserCommand, RegisterUserCommand } from '@application/commands';
import { UserNotFoundException } from '@domain/exceptions/user-not-found.exception';
import { MetricsService } from '@infra/metrics/metrics.service';

describe('UserController', () => {
  let controller: UserController;
  let commandBus: CommandBus;
  let service: MetricsService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        UserController,
        {
          provide: CommandBus,
          useValue: {
            execute: jest.fn(),
          },
        },
        {
          provide: MetricsService,
          useValue: {
            incrementUserCreated: jest.fn(),
          },
        },
      ],
    }).compile();

    controller = module.get<UserController>(UserController);
    commandBus = module.get<CommandBus>(CommandBus);
    service = module.get<MetricsService>(MetricsService);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
    expect(commandBus).toBeDefined();
    expect(service).toBeDefined();
  });

  describe('POST /', () => {
    it('should register a user', async () => {
      const executeSpy = jest
        .spyOn(commandBus, 'execute')
        .mockResolvedValueOnce({
          id: new Types.ObjectId().toString(),
          email: 'test@test.com',
          role: USER_ROLES.ADMIN,
        });

      const incrementUserCreatedSpy = jest.spyOn(
        service,
        'incrementUserCreated',
      );

      const createUserDto: CreateUserDto = {
        email: 'test@test.com',
        password: 'password',
        role: USER_ROLES.ADMIN,
      };

      const result = await controller.register(createUserDto);

      expect(executeSpy).toHaveBeenCalledWith(
        new RegisterUserCommand(
          createUserDto.email,
          createUserDto.password,
          createUserDto.role,
        ),
      );

      expect(incrementUserCreatedSpy).toHaveBeenCalled();

      expect(result).toEqual({
        id: expect.any(String) as string,
        email: 'test@test.com',
        role: USER_ROLES.ADMIN,
      });
    });

    it('should throw a conflict error when user already exists', async () => {
      const createUserDto: CreateUserDto = {
        email: 'test@test.com',
        password: 'password',
        role: USER_ROLES.ADMIN,
      };

      jest
        .spyOn(commandBus, 'execute')
        .mockRejectedValue(new UserAlreadyExistsException(createUserDto.email));

      await expect(controller.register(createUserDto)).rejects.toThrow(
        ConflictException,
      );
    });

    it('should throw an internal server error when something goes wrong', async () => {
      const createUserDto: CreateUserDto = {
        email: 'test@test.com',
        password: 'password',
        role: USER_ROLES.ADMIN,
      };

      jest.spyOn(commandBus, 'execute').mockRejectedValue(new Error());

      await expect(controller.register(createUserDto)).rejects.toThrow(
        InternalServerErrorException,
      );
    });
  });

  describe('DELETE /:id', () => {
    it('should remove a user', async () => {
      const executeSpy = jest.spyOn(commandBus, 'execute');
      const id = new Types.ObjectId().toString();
      await expect(controller.remove(id)).resolves.not.toThrow();
      expect(executeSpy).toHaveBeenCalledWith(new RemoveUserCommand(id));
    });

    it('should throw a not found error when user is not found', async () => {
      const id = new Types.ObjectId().toString();
      jest
        .spyOn(commandBus, 'execute')
        .mockRejectedValue(new UserNotFoundException(id));
      await expect(controller.remove(id)).rejects.toThrow(NotFoundException);
    });

    it('should throw an internal server error when something goes wrong', async () => {
      const id = new Types.ObjectId().toString();
      jest.spyOn(commandBus, 'execute').mockRejectedValue(new Error());
      await expect(controller.remove(id)).rejects.toThrow(
        InternalServerErrorException,
      );
    });
  });

  describe('POST /token', () => {
    it('should generate a token', async () => {
      jest.spyOn(commandBus, 'execute').mockResolvedValueOnce({
        token: 'token',
        expiresIn: 3600,
      });
      process.env.JWT_EXPIRES_IN = '3600';
      const result = await controller.generateUserToken({
        email: 'test@test.com',
        password: 'password',
      });
      delete process.env.JWT_EXPIRES_IN;
      expect(result).toEqual({
        token: expect.any(String) as string,
        expiresIn: expect.any(Number) as number,
      });
    });

    it('should throw a not found error when user is not found', async () => {
      jest
        .spyOn(commandBus, 'execute')
        .mockRejectedValue(new UserNotFoundException('test@test.com'));
      await expect(
        controller.generateUserToken({
          email: 'test@test.com',
          password: 'password',
        }),
      ).rejects.toThrow(NotFoundException);
    });

    it('should throw an internal server error when something goes wrong', async () => {
      jest.spyOn(commandBus, 'execute').mockRejectedValueOnce(new Error());
      await expect(
        controller.generateUserToken({
          email: 'test@test.com',
          password: 'password',
        }),
      ).rejects.toThrow(InternalServerErrorException);
    });
  });
});
