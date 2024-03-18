import { DatabaseMockRepository } from './mock/database_mock_repostory'
import { CreateLead } from './create_lead'

describe('Create Lead', () => {
  it('should create a new lead', async () => {
    const repository = new DatabaseMockRepository()
    const createLead = new CreateLead(repository)
    const result = await createLead.execute({
      email: 'test@test.com',
      converted: true,
    })
    expect(result.id).not.toBeNull()
    expect(repository.leads).toHaveLength(1)
  })
})
