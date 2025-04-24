import { Test, TestingModule } from '@nestjs/testing';
import { UserAbstractRepository } from '@domain/user-abstract.repository';
import { GenerateUserTokenHandler } from './generate-user-token.handler';
import { UserNotFoundException } from '@domain/exceptions/user-not-found.exception';
import { User, USER_ROLES } from '@domain/user.entity';
import { JwtService } from '@nestjs/jwt';
import { ConfigService } from '@nestjs/config';

const mockUser = User.create({
  email: 'test@test.com',
  password: 'password',
  role: USER_ROLES.ADMIN,
});

describe('GenerateUserTokenHandler', () => {
  let handler: GenerateUserTokenHandler;
  let repository: UserAbstractRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        GenerateUserTokenHandler,
        {
          provide: UserAbstractRepository,
          useValue: {
            findByEmail: jest.fn().mockResolvedValue(mockUser),
          },
        },
        {
          provide: JwtService,
          useValue: {
            signAsync: jest.fn().mockResolvedValue('token'),
          },
        },
        {
          provide: ConfigService,
          useValue: {
            get: jest.fn().mockReturnValue('3600'),
          },
        },
      ],
    }).compile();

    handler = module.get<GenerateUserTokenHandler>(GenerateUserTokenHandler);
    repository = module.get<UserAbstractRepository>(UserAbstractRepository);
  });

  it('should be defined', () => {
    expect(repository).toBeDefined();
    expect(handler).toBeDefined();
  });

  it('should generate a token', async () => {
    const result: { token: string; expiresIn: number } = await handler.execute({
      email: 'test@test.com',
      password: 'password',
    });
    expect(result.token).toEqual(expect.any(String));
    expect(result.expiresIn).toEqual(3600);
  });

  it('should throw an error if user is not found', async () => {
    jest
      .spyOn(repository, 'findByEmail')
      .mockRejectedValueOnce(new UserNotFoundException('test@test.com'));
    await expect(
      handler.execute({
        email: 'test@test.com',
        password: 'password',
      }),
    ).rejects.toThrow(UserNotFoundException);
  });

  it('should throw an error if password is invalid', async () => {
    jest.spyOn(repository, 'findByEmail').mockResolvedValueOnce(
      new User({
        id: '1',
        email: 'test@test.com',
        password: 'password',
        role: USER_ROLES.ADMIN,
      }),
    );
    await expect(
      handler.execute({
        email: 'test@test.com',
        password: 'invalid',
      }),
    ).rejects.toThrow(UserNotFoundException);
  });
});
