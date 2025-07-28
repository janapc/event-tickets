import { Test, TestingModule } from '@nestjs/testing';
import { CreateTicketCommand } from '@commands/create-ticket/create-ticket.command';
import { EventController } from './event.controller';
import { CommandBus } from '@nestjs/cqrs';

describe('EventController', () => {
  let controller: EventController;
  let bus: CommandBus;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        EventController,
        {
          provide: CommandBus,
          useValue: {
            execute: jest.fn(),
          },
        },
      ],
    }).compile();

    controller = module.get<EventController>(EventController);
    bus = module.get<CommandBus>(CommandBus);
  });

  it('should be defined', () => {
    expect(controller).toBeDefined();
    expect(bus).toBeDefined();
  });
  it('should call CommandBus.execute when sendTicket is called', async () => {
    const command = new CreateTicketCommand({
      name: 'Test User',
      email: 'user@test.com',
      eventId: '123',
      eventName: 'Event 1',
      eventDescription: 'description of event 1',
      eventImageUrl: 'http://example.com/image.jpg',
      language: 'en',
    });
    const busSpy = jest.spyOn(bus, 'execute').mockResolvedValue(undefined);

    await expect(controller.sendTicket(command)).resolves.toBeUndefined();

    expect(busSpy).toHaveBeenCalledWith(command);
  });

  it('should throw if CommandBus.execute throws', async () => {
    const command = new CreateTicketCommand({
      name: 'Test User',
      email: 'user@test.com',
      eventId: '123',
      eventName: 'Event 1',
      eventDescription: 'description of event 1',
      eventImageUrl: 'http://example.com/image.jpg',
      language: 'en',
    });
    const busSpy = jest
      .spyOn(bus, 'execute')
      .mockRejectedValueOnce(new Error('Failed'));

    await expect(controller.sendTicket(command)).rejects.toThrow('Failed');
    expect(busSpy).toHaveBeenCalled();
  });
});
