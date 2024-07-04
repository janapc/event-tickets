import { User } from '@domain/user'
import { InMemoryRepository } from './mock/in_memory_repository'
import { RemoveUser } from './remove_user'

describe('Remove User', () => {
  it('should remove a user', async () => {
    const user = new User('test@test.com', '123123123', 'ADMIN')
    user.id = '123'
    const repository = new InMemoryRepository()
    await repository.save(user)
    const removeUser = new RemoveUser(repository)
    await expect(removeUser.execute('test@test.com')).resolves.toBeUndefined()
    expect(repository.users.length).toEqual(0)
  })

  it('should not remove a user', async () => {
    const repository = new InMemoryRepository()
    const removeUser = new RemoveUser(repository)
    await expect(removeUser.execute('test@test.com')).rejects.toThrow()
    expect(repository.users.length).toEqual(0)
  })
})
