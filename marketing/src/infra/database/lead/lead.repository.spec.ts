import { getModelToken } from '@nestjs/mongoose';
import { Test, TestingModule } from '@nestjs/testing';
import mongoose, { Model } from 'mongoose';
import { LeadRepository } from './lead.repository';
import { LeadDocument, LeadModel } from './lead.schema';
import { Lead } from '@domain/lead';
import { LeadNotFoundException } from '@domain/exceptions/lead-not-found.exception';

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
            findOne: jest.fn().mockImplementation((): any => {
              return {
                email: 'test@test.com',
                converted: false,
                language: 'en',
                createdAt: new Date(),
                updatedAt: new Date(),
                _id: new mongoose.Types.ObjectId('6850496ae64e494cfaa8cf58'),
              };
            }),
            create: jest.fn().mockImplementation((lead): any => {
              return {
                ...lead,
                createdAt: new Date(),
                updatedAt: new Date(),
                _id: new mongoose.Types.ObjectId('6850496ae64e494cfaa8cf58'),
              };
            }),
            updateOne: jest.fn(),
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
    it('should save a lead and return it', async () => {
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

  describe('getByEmail', () => {
    it('should return a lead by email', async () => {
      const findOneSpy = jest.spyOn(leadModel, 'findOne');
      const email = 'test@test.com';
      const result = await repository.getByEmail(email);
      expect(result).toBeInstanceOf(Lead);
      expect(result.converted).toBe(false);
      expect(result.language).toBe('en');
      expect(result.email).toBe(email);
      expect(result.createdAt).toBeDefined();
      expect(result.updatedAt).toBeDefined();
      expect(result.id).toBeDefined();
      expect(findOneSpy).toHaveBeenCalledWith({ email });
    });

    it('should throw an error when lead is not found', async () => {
      const findOneSpy = jest
        .spyOn(leadModel, 'findOne')
        .mockResolvedValue(null);
      const email = 'test2@test.com';
      await expect(repository.getByEmail(email)).rejects.toThrow(
        LeadNotFoundException,
      );
      expect(findOneSpy).toHaveBeenCalledWith({ email });
    });
  });

  describe('converted', () => {
    it('should converted a lead', async () => {
      const updateOneSpy = jest.spyOn(leadModel, 'updateOne');
      await expect(
        repository.converted('test@test.com'),
      ).resolves.toBeUndefined();
      expect(updateOneSpy).toHaveBeenCalledTimes(1);
    });
    it('should throw an error when something goes wrong', async () => {
      const expectedError = new Error('Database save failed');
      const updateOneSpy = jest
        .spyOn(leadModel, 'updateOne')
        .mockRejectedValue(expectedError);
      await expect(repository.converted('erro@error.com')).rejects.toThrow(
        expectedError,
      );
      expect(updateOneSpy).toHaveBeenCalledTimes(1);
    });
  });
});
