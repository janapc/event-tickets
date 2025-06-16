import { Injectable } from '@nestjs/common';
import { Lead } from '@domain/lead';
import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { LeadDocument, LeadModel } from './lead.schema';
import { InjectModel } from '@nestjs/mongoose';
import { Model } from 'mongoose';
import { LeadDuplicatedException } from '@domain/exceptions/lead-duplicated.exception';
import { LeadNotFoundException } from '@domain/exceptions/lead-not-found.exception';

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
  async getByEmail(email: string): Promise<Lead> {
    const lead = await this.leadModel.findOne({ email });
    if (!lead) {
      throw new LeadNotFoundException(email);
    }
    return new Lead({
      id: lead._id.toString(),
      email: lead.email,
      converted: lead.converted,
      language: lead.language,
      createdAt: lead.createdAt,
      updatedAt: lead.updatedAt,
    });
  }

  async getAll(): Promise<Lead[]> {
    const leads = await this.leadModel.find();
    return leads.map(
      (lead) =>
        new Lead({
          id: lead._id.toString(),
          email: lead.email,
          converted: lead.converted,
          language: lead.language,
          createdAt: lead.createdAt,
          updatedAt: lead.updatedAt,
        }),
    );
  }
}
