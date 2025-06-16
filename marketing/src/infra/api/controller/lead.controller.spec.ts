import { Test, TestingModule } from '@nestjs/testing';
import { LeadController } from './lead.controller';
import { CommandBus } from '@nestjs/cqrs';
import { CreateLeadCommand } from '@commands/create-lead/create-lead.command';
import { CreateLeadDto } from './dtos/create-lead.dto';
import { Lead } from '@domain/lead';

describe('LeadController', () => {
  let controller: LeadController;
  let commandBus: CommandBus;

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
      ],
    }).compile();

    controller = module.get<LeadController>(LeadController);
    commandBus = module.get<CommandBus>(CommandBus);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
  });

  describe('create', () => {
    it('should execute CreateLeadCommand with correct parameters', async () => {
      const createLeadDto: CreateLeadDto = {
        email: 'test@example.com',
        converted: true,
        language: 'en',
      };
      const lead: Lead = {
        id: '1',
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

      jest
        .spyOn(commandBus, 'execute')
        .mockRejectedValue(new Error('Command failed'));

      await expect(controller.create(createLeadDto)).rejects.toThrow(
        'Command failed',
      );
    });
  });
});
