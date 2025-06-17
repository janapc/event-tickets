import { Test, TestingModule } from '@nestjs/testing';
import { LeadAbstractRepository } from '@domain/lead-abstract.repository';
import { ProcessCreatedClientHandler } from './process-created-client.handler';
import { ProcessCreatedClientCommand } from './process-created-client.query';
import { Lead } from '@domain/lead';
import { MetricsService } from '@infra/telemetry/metrics';

describe('ProcessCreatedClientHandler', () => {
  let handler: ProcessCreatedClientHandler;
  let leadRepository: LeadAbstractRepository;
  let metricsService: MetricsService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        ProcessCreatedClientHandler,
        {
          provide: LeadAbstractRepository,
          useValue: {
            converted: jest.fn(),
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
            save: jest.fn().mockImplementation((email: string): Lead => {
              return new Lead({
                email,
                converted: true,
                language: 'en',
                createdAt: new Date(),
                updatedAt: new Date(),
                id: '6850496ae64e494cfaa8cf51',
              });
            }),
          },
        },
        {
          provide: MetricsService,
          useValue: {
            incrementLeadCreated: jest.fn(),
          },
        },
      ],
    }).compile();

    handler = module.get<ProcessCreatedClientHandler>(
      ProcessCreatedClientHandler,
    );
    leadRepository = module.get<LeadAbstractRepository>(LeadAbstractRepository);
    metricsService = module.get<MetricsService>(MetricsService);
  });

  it('should convert a lead successfully', async () => {
    const convertedSpy = jest.spyOn(leadRepository, 'converted');
    const command = new ProcessCreatedClientCommand(
      'message123',
      'test@example.com',
    );

    await expect(handler.execute(command)).resolves.toBeUndefined();
    expect(convertedSpy).toHaveBeenCalledWith(command.email);
  });

  it('should create an already converted lead', async () => {
    const getByEmailSpy = jest
      .spyOn(leadRepository, 'getByEmail')
      .mockResolvedValueOnce(null);
    const saveSpy = jest.spyOn(leadRepository, 'save');
    const incrementLeadCreatedSpy = jest.spyOn(
      metricsService,
      'incrementLeadCreated',
    );
    const command = new ProcessCreatedClientCommand(
      'message123',
      'test@example.com',
    );

    await expect(handler.execute(command)).resolves.toBeUndefined();
    expect(getByEmailSpy).toHaveBeenCalledWith(command.email);
    expect(saveSpy).toHaveBeenCalled();
    expect(incrementLeadCreatedSpy).toHaveBeenCalled();
  });
});
