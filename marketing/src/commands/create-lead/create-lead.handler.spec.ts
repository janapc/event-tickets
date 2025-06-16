import { Test, TestingModule } from '@nestjs/testing';
import { CreateLeadHandler } from './create-lead.handler';
import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { CreateLeadCommand } from './create-lead.command';
import { Lead } from '@domain/lead';

describe('CreateLeadHandler', () => {
  let handler: CreateLeadHandler;
  let leadRepository: LeadAbstractRepository;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        CreateLeadHandler,
        {
          provide: LeadAbstractRepository,
          useValue: {
            save: jest.fn().mockImplementation((lead: Lead): Lead => {
              return new Lead({
                email: lead.email,
                converted: lead.converted,
                language: lead.language,
                createdAt: new Date(),
                updatedAt: new Date(),
                id: 'mockedId123',
              });
            }),
          },
        },
      ],
    }).compile();

    handler = module.get<CreateLeadHandler>(CreateLeadHandler);
    leadRepository = module.get<LeadAbstractRepository>(LeadAbstractRepository);
  });

  it('should create a lead successfully', async () => {
    const saveSpy = jest.spyOn(leadRepository, 'save');
    const command = new CreateLeadCommand('test@example.com', false, 'en');

    const result = await handler.execute(command);
    expect(result).toBeInstanceOf(Lead);
    expect(result.email).toBe(command.email);
    expect(result.converted).toBe(command.converted);
    expect(result.language).toBe(command.language);
    expect(result.id).toBeDefined();
    expect(result.createdAt).toBeDefined();
    expect(result.updatedAt).toBeDefined();
    expect(saveSpy).toHaveBeenCalledTimes(1);
  });

  it('should rethrow other errors', async () => {
    const command = new CreateLeadCommand('error@example.com', false, 'en');
    const unexpectedError = new Error('Unexpected error');
    const saveSpy = jest
      .spyOn(leadRepository, 'save')
      .mockRejectedValueOnce(unexpectedError);

    await expect(handler.execute(command)).rejects.toThrow(unexpectedError);
    expect(saveSpy).toHaveBeenCalled();
  });
});
