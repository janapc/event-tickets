import { Test, TestingModule } from '@nestjs/testing';
import { Types } from 'mongoose';
import { USER_ROLES } from '@domain/user.entity';
import { RegisterUserHandler } from './register-user.handler';
import { UserAbstractRepository } from '@domain/user-abstract.repository';
import { UserAlreadyExistsException } from '@domain/exceptions/user-already-exists.exception';

const mockUser = {
  id: new Types.ObjectId().toString(),
  email: 'test@example.com',
  password: '12345678',
  role: USER_ROLES.ADMIN,
};

describe('RegisterUserHandler', () => {
  let handler: RegisterUserHandler;
  let repository: UserAbstractRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        RegisterUserHandler,
        {
          provide: UserAbstractRepository,
          useValue: {
            create: jest.fn().mockResolvedValue(mockUser),
          },
        },
      ],
    }).compile();

    handler = module.get<RegisterUserHandler>(RegisterUserHandler);
    repository = module.get<UserAbstractRepository>(UserAbstractRepository);
  });

  it('should be defined', () => {
    expect(repository).toBeDefined();
    expect(handler).toBeDefined();
  });

  it('should create a new user', async () => {
    await expect(
      handler.execute({
        email: 'test@example.com',
        password: '12345678',
        role: USER_ROLES.ADMIN,
      }),
    ).resolves.toEqual({
      id: expect.any(String) as string,
      email: 'test@example.com',
      role: USER_ROLES.ADMIN,
    });
    expect(repository.create).toHaveBeenCalledTimes(1);
  });

  it('should throw an error if email is already taken', async () => {
    jest
      .spyOn(repository, 'create')
      .mockRejectedValueOnce(
        new UserAlreadyExistsException('test@example.com'),
      );
    await expect(
      handler.execute({
        email: 'test@example.com',
        password: '12345678',
        role: USER_ROLES.ADMIN,
      }),
    ).rejects.toThrow(UserAlreadyExistsException);
  });
});
