import { type PrismaClient } from '@prisma/client'
import { LeadPrismaRepository } from './lead_prisma_repository'
import { prismaMock } from './mock/prisma_mock'
import { Lead } from '../../domain/lead'

describe('Lead Prisma Repository', () => {
  it('Should create a lead', async () => {
    const lead = new Lead('test@test.com', false, 'en')
    prismaMock.lead.create.mockResolvedValue({ ...lead, id: 'asd-asd' })
    const repository = new LeadPrismaRepository(prismaMock as PrismaClient)
    const result = await repository.save(lead)
    expect(result.id).not.toBeNull()
    expect(result.email).toEqual('test@test.com')
    expect(result.converted).toBeFalsy()
  })

  it('Should get all leads', async () => {
    prismaMock.lead.findMany.mockResolvedValue([
      {
        email: 'test@test.com',
        converted: false,
        id: 'asd-asd',
        language: 'pt',
      },
      {
        email: 'test2@test.com',
        converted: true,
        id: 'fgh-asd',
        language: 'en',
      },
    ])
    const repository = new LeadPrismaRepository(prismaMock as PrismaClient)
    await expect(repository.getAll()).resolves.toHaveLength(2)
  })

  it('Should update a lead', async () => {
    const lead = new Lead('test@test.com', false, 'en')
    prismaMock.lead.create.mockResolvedValue({ ...lead, id: 'asd-asd' })
    prismaMock.lead.update.mockResolvedValue({
      ...lead,
      id: 'asd-asd',
      converted: true,
    })
    const repository = new LeadPrismaRepository(prismaMock as PrismaClient)
    await expect(
      repository.update('test@test.com', true),
    ).resolves.toBeUndefined()
  })

  it('Should get a lead by email', async () => {
    const lead = new Lead('test@test.com', false, 'pt')
    prismaMock.lead.findUnique.mockResolvedValue({ ...lead, id: 'asd-asd' })
    const repository = new LeadPrismaRepository(prismaMock as PrismaClient)
    const result = await repository.getByEmail('test@test.com')
    expect(result?.id).toEqual('asd-asd')
  })
})
