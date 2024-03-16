import { User } from '@domain/user'
import { UserRepository } from './user_repository'
import { userModel } from './schema'

describe('User Database', () => {
  afterEach(() => {
    jest.resetAllMocks()
  })

  it('should save a user on database', async () => {
    const user = new User('email@email.com', '12345678', 'ADMIN')
    const spySave = jest
      .spyOn(userModel.prototype, 'save')
      .mockResolvedValue(true)
    const repository = new UserRepository(userModel)
    await expect(repository.save(user)).resolves.toBeUndefined()
    expect(spySave).toHaveBeenCalledTimes(1)
  })

  it('should error if the user cannot be saved', async () => {
    const user = new User('email@email.com', '12345678', 'PUBLIC')
    const spySave = jest
      .spyOn(userModel.prototype, 'save')
      .mockRejectedValue(new Error('database error'))
    const repository = new UserRepository(userModel)
    await expect(repository.save(user)).rejects.toThrow('database error')
    expect(spySave).toHaveBeenCalledTimes(1)
  })

  it('should find user by email', async () => {
    const userMock = {
      email: 'email@email.com',
      password: '03a851ff97d17ca7638850a47d4e0146d3b1adb6',
      role: 'PUBLIC',
      _id: 'Asd123',
    }
    const spyFindOne = jest
      .spyOn(userModel, 'findOne')
      .mockResolvedValue(userMock)
    const repository = new UserRepository(userModel)
    await expect(
      repository.findByEmail('email@email.com'),
    ).resolves.toMatchObject({
      email: userMock.email,
      id: userMock._id.toString(),
      role: userMock.role,
    })
    expect(spyFindOne).toHaveBeenCalledTimes(1)
  })

  it('should error if the user cannot be found', async () => {
    const spyFindOne = jest.spyOn(userModel, 'findOne').mockResolvedValue(null)
    const repository = new UserRepository(userModel)
    await expect(repository.findByEmail('email@email.com')).resolves.toBeNull()
    expect(spyFindOne).toHaveBeenCalledTimes(1)
  })

  it('should remove a user', async () => {
    const spyFindOneAndDelete = jest
      .spyOn(userModel, 'findOneAndDelete')
      .mockResolvedValue({})
    const repository = new UserRepository(userModel)
    await expect(repository.remove('test@test.com')).resolves.toBeUndefined()
    expect(spyFindOneAndDelete).toHaveBeenCalledTimes(1)
  })

  it('should error remove a user', async () => {
    const spyFindOneAndDelete = jest
      .spyOn(userModel, 'findOneAndDelete')
      .mockRejectedValue(new Error('database error'))
    const repository = new UserRepository(userModel)
    await expect(repository.remove('test@test.com')).rejects.toThrow(
      'database error',
    )
    expect(spyFindOneAndDelete).toHaveBeenCalledTimes(1)
  })
})
