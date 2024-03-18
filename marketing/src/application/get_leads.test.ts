import { Lead } from '@domain/lead'
import { GetLeads } from './get_leads'
import { DatabaseMockRepository } from './mock/database_mock_repostory'

describe('Get Leads', () => {
  it('should get all leads', async () => {
    const repository = new DatabaseMockRepository()
    await repository.save(new Lead('test@test.com', false))
    const getLeads = new GetLeads(repository)
    const result = await getLeads.execute()
    expect(result).toHaveLength(1)
  })
})
