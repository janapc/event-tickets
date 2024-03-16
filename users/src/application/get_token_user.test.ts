import { User } from '@domain/user'
import { InMemoryRepository } from './mock/in_memory_repository'
import { GetTokenUser } from './get_token_user'

describe('Get Token User', () => {
  beforeAll(() => {
    process.env.JWT_SECRET = 'test'
    process.env.JWT_EXPIRES_IN = '24h'
  })
  it('should get token', async () => {
    const user = new User('test@test.com', '123123123', 'ADMIN')
    user.id = '123'
    const repository = new InMemoryRepository()
    await repository.save(user)
    expect(repository.users).toHaveLength(1)
    const getTokenUser = new GetTokenUser(repository)
    const result = await getTokenUser.execute('test@test.com', '123123123')
    expect(result).not.toBeNull()
  })

  it('should error if the user not exists', async () => {
    const repository = new InMemoryRepository()
    const getTokenUser = new GetTokenUser(repository)
    await expect(
      getTokenUser.execute('test@test.com', '123123123'),
    ).rejects.toEqual(new Error('we can not find the user.Please try again'))
  })

  it('should error if the password is wrong', async () => {
    const user = new User('test@test.com', '123123123', 'ADMIN')
    user.id = '123'
    const repository = new InMemoryRepository()
    await repository.save(user)
    expect(repository.users).toHaveLength(1)
    const getTokenUser = new GetTokenUser(repository)
    await expect(getTokenUser.execute('test@test.com', '123')).rejects.toEqual(
      new Error('your email or password is incorrect.Please try again'),
    )
  })
})
