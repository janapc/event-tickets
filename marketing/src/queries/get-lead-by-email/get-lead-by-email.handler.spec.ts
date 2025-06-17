import { Test, TestingModule } from '@nestjs/testing';
import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { Lead } from '@domain/lead';
import { GetLeadByEmailHandler } from './get-lead-by-email.handler';
import { GetLeadByEmailQuery } from './get-lead-by-email.query';
import { LeadNotFoundException } from '@domain/exceptions/lead-not-found.exception';

describe('GetLeadByEmailHandler', () => {
  let handler: GetLeadByEmailHandler;
  let leadRepository: LeadAbstractRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        GetLeadByEmailHandler,
        {
          provide: LeadAbstractRepository,
          useValue: {
            getByEmail: jest.fn().mockImplementation((email: string): Lead => {
              return new Lead({
                email,
                converted: false,
                language: 'en',
                createdAt: new Date(),
                updatedAt: new Date(),
                id: '6850496ae64e494cfaa8cf58',
              });
            }),
          },
        },
      ],
    }).compile();

    handler = module.get<GetLeadByEmailHandler>(GetLeadByEmailHandler);
    leadRepository = module.get<LeadAbstractRepository>(LeadAbstractRepository);
  });

  it('should get a lead by email successfully', async () => {
    const query = new GetLeadByEmailQuery('test@test.com');
    const getByEmailSpy = jest.spyOn(leadRepository, 'getByEmail');
    const result = await handler.execute(query);
    expect(result.email).toBe(query.email);
    expect(result.converted).toBe(false);
    expect(result.language).toBe('en');
    expect(result.id).toBeDefined();
    expect(result.createdAt).toBeDefined();
    expect(result.updatedAt).toBeDefined();
    expect(getByEmailSpy).toHaveBeenCalledWith(query.email);
  });

  it('should erro if lead not found', async () => {
    const query = new GetLeadByEmailQuery('error@example.com');
    const findByEmailSpy = jest
      .spyOn(leadRepository, 'getByEmail')
      .mockResolvedValueOnce(null);

    await expect(handler.execute(query)).rejects.toThrow(LeadNotFoundException);
    expect(findByEmailSpy).toHaveBeenCalled();
  });
});
