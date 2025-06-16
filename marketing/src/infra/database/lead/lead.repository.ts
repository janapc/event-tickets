import { Injectable } from '@nestjs/common';
import { Lead } from '@domain/lead';
import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { LeadDocument, LeadModel } from './lead.schema';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { LeadDuplicatedException } from '@domain/exceptions/lead-duplicated.exception';

@Injectable()
export class LeadRepository implements LeadAbstractRepository {
  constructor(
    @InjectModel(LeadModel.name) private leadModel: Model<LeadDocument>,
  ) {}

  async save(lead: Lead): Promise<Lead> {
    try {
      const createdLead = await this.leadModel.create(lead);
      return new Lead({
        id: createdLead._id.toString(),
        email: createdLead.email,
        converted: createdLead.converted,
        language: createdLead.language,
        createdAt: createdLead.createdAt,
        updatedAt: createdLead.updatedAt,
      });
    } catch (error: any) {
      const mongooseError = error as {
        code: number;
        keyPattern?: Record<string, any>;
      };
      if (mongooseError.code === 11000) {
        throw new LeadDuplicatedException(lead.email);
      }
      throw error;
    }
  }
}
