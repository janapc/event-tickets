import { PrismaClient } from '@prisma/client'
import { logger } from '@infra/logger/logger'

export default class Database {
  private static instance: Database | null = null
  readonly connection: PrismaClient

  private constructor() {
    this.connection = new PrismaClient()
  }

  public static getInstance(): Database {
    if (Database.instance === null) {
      Database.instance = new Database()
    }

    return Database.instance
  }

  close(): void {
    this.connection.$disconnect().catch((error) => {
      logger.error(`database close error ${String(error)}`)
    })
  }
}
