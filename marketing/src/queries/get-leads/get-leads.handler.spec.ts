import { Test, TestingModule } from '@nestjs/testing';
import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { Lead } from '@domain/lead';
import { GetLeadsHandler } from './get-leads.handler';
import { GetLeadsQuery } from './get-leads.query';

describe('GetLeadsHandler', () => {
  let handler: GetLeadsHandler;
  let leadRepository: LeadAbstractRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        GetLeadsHandler,
        {
          provide: LeadAbstractRepository,
          useValue: {
            getAll: jest.fn().mockImplementation((): Lead[] => {
              return [
                new Lead({
                  email: 'test@test.com',
                  converted: false,
                  language: 'en',
                  createdAt: new Date(),
                  updatedAt: new Date(),
                  id: '6850496ae64e494cfaa8cf58',
                }),
              ];
            }),
          },
        },
      ],
    }).compile();

    handler = module.get<GetLeadsHandler>(GetLeadsHandler);
    leadRepository = module.get<LeadAbstractRepository>(LeadAbstractRepository);
  });

  it('should get all leads successfully', async () => {
    const query = new GetLeadsQuery();
    const getAllSpy = jest.spyOn(leadRepository, 'getAll');
    const result = await handler.execute(query);
    expect(result).toBeInstanceOf(Array);
    expect(result.length).toBeGreaterThan(0);
    expect(result[0].email).toEqual('test@test.com');
    expect(result[0].converted).toBeFalsy();
    expect(result[0].language).toEqual('en');
    expect(result[0].id).toEqual('6850496ae64e494cfaa8cf58');
    expect(result[0].createdAt).toBeDefined();
    expect(result[0].updatedAt).toBeDefined();
    expect(getAllSpy).toHaveBeenCalledWith();
  });
});
