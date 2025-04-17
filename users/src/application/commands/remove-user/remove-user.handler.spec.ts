import { Test, TestingModule } from '@nestjs/testing';
import { Types } from 'mongoose';
import { UserAbstractRepository } from '@domain/user-abstract.repository';
import { RemoveUserHandler } from './remove-user.handler';
import { UserNotFoundException } from '@domain/exceptions/user-not-found.exception';

describe('RemoveUserHandler', () => {
  let handler: RemoveUserHandler;
  let repository: UserAbstractRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        RemoveUserHandler,
        {
          provide: UserAbstractRepository,
          useValue: {
            delete: jest.fn(),
          },
        },
      ],
    }).compile();

    handler = module.get<RemoveUserHandler>(RemoveUserHandler);
    repository = module.get<UserAbstractRepository>(UserAbstractRepository);
  });

  it('should be defined', () => {
    expect(repository).toBeDefined();
    expect(handler).toBeDefined();
  });

  it('should remove a user', async () => {
    const id = new Types.ObjectId().toString();
    await expect(
      handler.execute({
        id,
      }),
    ).resolves.not.toThrow();
    expect(repository.delete).toHaveBeenCalledTimes(1);
    expect(repository.delete).toHaveBeenCalledWith(id);
  });

  it('should throw an error if user is not found', async () => {
    const id = new Types.ObjectId().toString();
    jest
      .spyOn(repository, 'delete')
      .mockRejectedValueOnce(new UserNotFoundException(id));
    await expect(
      handler.execute({
        id,
      }),
    ).rejects.toThrow(UserNotFoundException);
  });

  it('should throw an error if something goes wrong', async () => {
    const id = new Types.ObjectId().toString();
    jest.spyOn(repository, 'delete').mockRejectedValueOnce(new Error());
    await expect(
      handler.execute({
        id,
      }),
    ).rejects.toThrow(Error);
  });
});
