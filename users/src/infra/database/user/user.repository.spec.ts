import { getModelToken } from '@nestjs/mongoose';
import { Test, TestingModule } from '@nestjs/testing';
import { Model, Types } from 'mongoose';
import { User, USER_ROLES } from '@domain/user.entity';
import { UserRepository } from './user.repository';
import { UserModel } from './user.schema';
import { UserAlreadyExistsException } from '@domain/exceptions/user-already-exists.exception';
import { UserNotFoundException } from '@domain/exceptions/user-not-found.exception';

const mockUser = {
  _id: new Types.ObjectId(),
  email: 'test@example.com',
  password: '12345678',
  role: USER_ROLES.ADMIN,
};

describe('UserRepository', () => {
  let repository: UserRepository;
  let model: jest.Mocked<Model<UserModel>>;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        UserRepository,
        {
          provide: getModelToken(UserModel.name),
          useValue: {
            create: jest.fn().mockResolvedValue(mockUser),
            findOne: jest.fn().mockResolvedValue(mockUser),
            deleteOne: jest.fn(),
          },
        },
      ],
    }).compile();

    repository = module.get<UserRepository>(UserRepository);
    model = module.get(getModelToken(UserModel.name));
  });

  it('should be defined', () => {
    expect(repository).toBeDefined();
    expect(model).toBeDefined();
  });

  describe('create', () => {
    it('should create a user successfully', async () => {
      const createSpy = jest.spyOn(model, 'create');
      const user = User.create({
        email: 'test@test.com',
        password: '123456',
        role: USER_ROLES.ADMIN,
      });
      const createdUser = await repository.create(user);
      expect(createdUser).toBeInstanceOf(User);
      expect(createSpy).toHaveBeenCalledWith(user);
      expect(createSpy).toHaveBeenCalledTimes(1);
    });
    it('should throw an error when user already exists', async () => {
      const error: any = new Error('User already exists');
      error.code = 11000;
      const createSpy = jest.spyOn(model, 'create').mockRejectedValue(error);
      const user = User.create({
        email: 'test@test.com',
        password: '123456',
        role: USER_ROLES.ADMIN,
      });
      await expect(repository.create(user)).rejects.toThrow(
        UserAlreadyExistsException,
      );
      expect(createSpy).toHaveBeenCalledTimes(1);
    });

    it('should throw an error when something goes wrong', async () => {
      const error = new Error('Internal server error');
      const createSpy = jest.spyOn(model, 'create').mockRejectedValue(error);
      const user = User.create({
        email: 'test@test.com',
        password: '123456',
        role: USER_ROLES.ADMIN,
      });
      await expect(repository.create(user)).rejects.toThrow(error);
      expect(createSpy).toHaveBeenCalledTimes(1);
    });
  });

  describe('findByEmail', () => {
    it('should find a user by email', async () => {
      const findOneSpy = jest.spyOn(model, 'findOne');
      const userEmail = 'test@test.com';
      const foundUser = await repository.findByEmail(userEmail);
      expect(findOneSpy).toHaveBeenCalledWith({ email: userEmail });
      expect(findOneSpy).toHaveBeenCalledTimes(1);
      expect(foundUser).toBeInstanceOf(User);
    });

    it('should throw an error when user is not found', async () => {
      const findOneSpy = jest
        .spyOn(model, 'findOne')
        .mockResolvedValueOnce(null);

      await expect(
        repository.findByEmail('nonexistent@example.com'),
      ).rejects.toThrow(UserNotFoundException);

      expect(findOneSpy).toHaveBeenCalledWith({
        email: 'nonexistent@example.com',
      });
    });
  });

  describe('delete', () => {
    it('should delete a user successfully', async () => {
      const deleteOne = jest
        .spyOn(model, 'deleteOne')
        .mockResolvedValueOnce({ acknowledged: true, deletedCount: 1 });
      const userId = new Types.ObjectId().toString();
      await expect(repository.delete(userId)).resolves.not.toThrow();
      expect(deleteOne).toHaveBeenCalledWith({
        _id: userId,
      });
    });

    it('should throw an error when trying to delete non-existent user', async () => {
      const deleteOne = jest
        .spyOn(model, 'deleteOne')
        .mockResolvedValueOnce({ acknowledged: false, deletedCount: 0 });
      const userId = new Types.ObjectId().toString();
      await expect(repository.delete(userId)).rejects.toThrow(
        UserNotFoundException,
      );

      expect(deleteOne).toHaveBeenCalledWith({
        _id: userId,
      });
    });
  });
});
