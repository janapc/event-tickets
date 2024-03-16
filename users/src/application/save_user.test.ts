import { User } from '@domain/user'
import { InMemoryRepository } from './mock/in_memory_repository'
import { SaveUser } from './save_user'

describe('Save User', () => {
  it('should save a user', async () => {
    const repository = new InMemoryRepository()
    const saveUser = new SaveUser(repository)
    await expect(
      saveUser.execute({
        role: 'PUBLIC',
        email: 'asd@asd.com',
        password: 'asdasd123',
      }),
    ).resolves.toBeUndefined()
    expect(repository.users).toHaveLength(1)
  })

  it('should error if the user exists', async () => {
    const user = new User('test@test.com', '123123123', 'ADMIN')
    user.id = '123'
    const repository = new InMemoryRepository()
    await repository.save(user)
    const saveUser = new SaveUser(repository)
    await expect(
      saveUser.execute({
        role: 'PUBLIC',
        email: 'test@test.com',
        password: 'asdasd123',
      }),
    ).rejects.toThrow('email already registered')
    expect(repository.users).toHaveLength(1)
  })
})
