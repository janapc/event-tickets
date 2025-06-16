import { getModelToken } from '@nestjs/mongoose';
import { Test, TestingModule } from '@nestjs/testing';
import { Model } from 'mongoose';
import { LeadRepository } from './lead.repository';
import { LeadDocument, LeadModel } from './lead.schema';
import { Lead } from '@domain/lead';

describe('LeadRepository', () => {
  let repository: LeadRepository;
  let leadModel: Model<LeadDocument>;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        LeadRepository,
        {
          provide: getModelToken(LeadModel.name),
          useValue: {
            create: jest.fn().mockImplementation((lead): LeadDocument => {
              return {
                ...lead,
                createdAt: Date.now(),
                updatedAt: Date.now(),
                _id: 'mockedId123',
              };
            }),
          },
        },
      ],
    }).compile();

    repository = module.get<LeadRepository>(LeadRepository);
    leadModel = module.get<Model<LeadDocument>>(getModelToken(LeadModel.name));
  });

  it('should be defined', () => {
    expect(repository).toBeDefined();
    expect(leadModel).toBeDefined();
  });

  describe('save', () => {
    it('should save a lead and return in', async () => {
      const inputLead = new Lead({
        email: 'email@email.com',
        converted: false,
        language: 'en',
      });
      const createSpy = jest.spyOn(leadModel, 'create');
      const result = await repository.save(inputLead);
      expect(result).toBeInstanceOf(Lead);
      expect(result.email).toEqual(inputLead.email);
      expect(result.converted).toEqual(inputLead.converted);
      expect(result.language).toEqual(inputLead.language);
      expect(result.createdAt).toBeDefined();
      expect(result.updatedAt).toBeDefined();
      expect(result.id).toBeDefined();
      expect(createSpy).toHaveBeenCalledTimes(1);
    });
    it('should throw an error when something goes wrong', async () => {
      const expectedError = new Error('Database save failed');
      const createSpy = jest
        .spyOn(leadModel, 'create')
        .mockRejectedValue(expectedError);
      const inputLead = new Lead({
        email: 'email@email.com',
        converted: false,
        language: 'en',
      });
      await expect(repository.save(inputLead)).rejects.toThrow(expectedError);
      expect(createSpy).toHaveBeenCalledTimes(1);
    });
  });
});
