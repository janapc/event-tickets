import { Test, TestingModule } from '@nestjs/testing';
import { LeadController } from './lead.controller';
import { CommandBus, QueryBus } from '@nestjs/cqrs';
import { CreateLeadCommand } from '@commands/create-lead/create-lead.command';
import { CreateLeadDto } from './dtos/create-lead.dto';
import { Lead } from '@domain/lead';
import { GetLeadByEmailQuery } from '@queries/get-lead-by-email/get-lead-by-email.query';
import { GetLeadsQuery } from '@queries/get-leads/get-leads.query';
import { ProcessCreatedClientCommand } from '@commands/process-created-client/process-created-client.query';

describe('LeadController', () => {
  let controller: LeadController;
  let commandBus: CommandBus;
  let queryBus: QueryBus;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      controllers: [LeadController],
      providers: [
        {
          provide: CommandBus,
          useValue: {
            execute: jest.fn(),
          },
        },
        {
          provide: QueryBus,
          useValue: {
            execute: jest.fn(),
          },
        },
      ],
    }).compile();

    controller = module.get<LeadController>(LeadController);
    commandBus = module.get<CommandBus>(CommandBus);
    queryBus = module.get<QueryBus>(QueryBus);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
    expect(commandBus).toBeDefined();
    expect(queryBus).toBeDefined();
  });

  describe('create', () => {
    it('should execute CreateLeadCommand with correct parameters', async () => {
      const createLeadDto: CreateLeadDto = {
        email: 'test@example.com',
        converted: true,
        language: 'en',
      };
      const lead: Lead = {
        id: '6850496ae64e494cfaa8cf58',
        email: createLeadDto.email,
        converted: createLeadDto.converted,
        language: createLeadDto.language,
        createdAt: new Date(),
        updatedAt: new Date(),
      };

      const executeSpy = jest
        .spyOn(commandBus, 'execute')
        .mockResolvedValue(lead);

      const result = await controller.create(createLeadDto);

      expect(executeSpy).toHaveBeenCalledWith(
        new CreateLeadCommand(
          createLeadDto.email,
          createLeadDto.converted,
          createLeadDto.language,
        ),
      );
      expect(result).toEqual(lead);
    });

    it('should throw an error if command execution fails', async () => {
      const createLeadDto: CreateLeadDto = {
        email: 'test@example.com',
        converted: true,
        language: 'en',
      };
      const expectedError = new Error('Command failed');

      jest.spyOn(commandBus, 'execute').mockRejectedValue(expectedError);

      await expect(controller.create(createLeadDto)).rejects.toThrow(
        expectedError,
      );
    });
  });

  describe('getByEmail', () => {
    it('should return a lead by email', async () => {
      const email = 'test@test.com';
      const lead: Lead = {
        id: '6850496ae64e494cfaa8cf58',
        email,
        converted: false,
        language: 'en',
        createdAt: new Date(),
        updatedAt: new Date(),
      };

      const executeSpy = jest
        .spyOn(queryBus, 'execute')
        .mockResolvedValue(lead);

      const result = await controller.getByEmail(email);

      expect(executeSpy).toHaveBeenCalledWith(new GetLeadByEmailQuery(email));
      expect(result).toEqual(lead);
    });

    it('should throw an error if lead not found', async () => {
      const email = 'error@error.com';
      const expectedError = new Error('Lead not found');

      const executeSpy = jest
        .spyOn(queryBus, 'execute')
        .mockRejectedValueOnce(expectedError);

      await expect(controller.getByEmail(email)).rejects.toThrow(expectedError);
      expect(executeSpy).toHaveBeenCalledWith(new GetLeadByEmailQuery(email));
    });
  });

  describe('getAll', () => {
    it('should return all leads', async () => {
      const lead: Lead = {
        id: '6850496ae64e494cfaa8cf58',
        email: 'test@test.com',
        converted: false,
        language: 'en',
        createdAt: new Date(),
        updatedAt: new Date(),
      };

      const executeSpy = jest
        .spyOn(queryBus, 'execute')
        .mockResolvedValue([lead]);

      const result = await controller.getAll();

      expect(executeSpy).toHaveBeenCalledWith(new GetLeadsQuery());
      expect(result).toHaveLength(1);
      expect(result[0]).toEqual(lead);
    });
  });

  describe('handleClientCreated', () => {
    it('should execute handleClientCreated with correct parameters', async () => {
      const message = {
        email: 'test@example.com',
        messageId: 'message123',
      };

      const executeSpy = jest.spyOn(commandBus, 'execute');

      await expect(
        controller.handleClientCreated(message),
      ).resolves.toBeUndefined();

      expect(executeSpy).toHaveBeenCalledWith(
        new ProcessCreatedClientCommand(message.messageId, message.email),
      );
    });
  });
});
