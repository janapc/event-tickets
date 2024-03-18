import { type PrismaClient } from '@prisma/client'
import { type Lead } from '@domain/lead'
import { type ILeadRepository } from '@domain/repository'

export class LeadPrismaRepository implements ILeadRepository {
  constructor(private readonly prisma: PrismaClient) {}
  async getByEmail(email: string): Promise<Lead | null> {
    return await this.prisma.lead.findUnique({
      where: { email },
    })
  }

  async save(lead: Lead): Promise<Lead> {
    return await this.prisma.lead.create({
      data: {
        email: lead.email,
        converted: lead.converted,
      },
    })
  }

  async getAll(): Promise<Lead[]> {
    return await this.prisma.lead.findMany()
  }

  async update(email: string, converted: boolean): Promise<void> {
    await this.prisma.lead.update({
      where: {
        email,
      },
      data: {
        converted,
      },
    })
  }
}
