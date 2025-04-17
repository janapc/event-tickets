import { Lead } from '@domain/lead'
import { DatabaseMockRepository } from './mock/database_mock_repostory'
import { GetByEmail } from './get_by_email'

describe('Get By Email', () => {
  it('should return a lead if found', async () => {
    const repository = new DatabaseMockRepository()
    const email = 'test.test@test.com'
    await repository.save(new Lead(email, false, 'pt'))
    const application = new GetByEmail(repository)
    const result = await application.execute(email)
    expect(result.id).not.toBeNull()
    expect(result.email).toEqual(email)
    expect(result.language).toEqual('pt')
  })

  it('should error if not found', async () => {
    const repository = new DatabaseMockRepository()
    const email = 'test.test@test.com'
    const application = new GetByEmail(repository)
    await expect(application.execute(email)).rejects.toThrow(
      'lead is not found',
    )
  })
})
