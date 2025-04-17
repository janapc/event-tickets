import { DatabaseMockRepository } from './mock/database_mock_repostory'
import { CreateLead } from './create_lead'
import { Lead } from '@domain/lead'

describe('Create Lead', () => {
  it('should create a new lead', async () => {
    const repository = new DatabaseMockRepository()
    const spy = jest.spyOn(repository, 'save')
    const createLead = new CreateLead(repository)
    const result = await createLead.execute({
      email: 'test@test.com',
      converted: true,
      language: 'pt',
    })
    expect(result.id).not.toBeNull()
    expect(repository.leads).toHaveLength(1)
    expect(spy).toHaveBeenCalledTimes(1)
  })

  it('should return a lead if already exists', async () => {
    const repository = new DatabaseMockRepository()
    const spy = jest.spyOn(repository, 'getByEmail')
    const createLead = new CreateLead(repository)
    await repository.save(new Lead('test@test.com', true, 'pt'))
    const result = await createLead.execute({
      email: 'test@test.com',
      converted: true,
      language: 'pt',
    })
    expect(result.id).not.toBeNull()
    expect(repository.leads).toHaveLength(1)
    expect(spy).toHaveBeenCalledTimes(1)
  })
})
