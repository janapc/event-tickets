import { Test, TestingModule } from '@nestjs/testing';
import { CreateTicketHandler } from './create-ticket.handler';
import { TicketAbstractRepository } from '@domain/ticket-abstract.repository';
import { MailAbstract } from '@domain/mail';
import { Ticket } from '@domain/ticket.entity';
import { CreateTicketCommand } from './create-ticket.command';

describe('CreateTicketHandler', () => {
  let handler: CreateTicketHandler;
  let repository: TicketAbstractRepository;
  let mail: MailAbstract;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [
        CreateTicketHandler,
        {
          provide: TicketAbstractRepository,
          useValue: {
            save: jest.fn().mockImplementation((ticket: Ticket): Ticket => {
              return new Ticket({
                email: ticket.email,
                eventId: ticket.eventId,
                passport: ticket.passport,
                id: 'generated-id',
              });
            }),
          },
        },
        {
          provide: MailAbstract,
          useValue: {
            sendMail: jest.fn().mockResolvedValue(undefined),
          },
        },
      ],
    }).compile();

    handler = module.get<CreateTicketHandler>(CreateTicketHandler);
    repository = module.get<TicketAbstractRepository>(TicketAbstractRepository);
    mail = module.get<MailAbstract>(MailAbstract);
  });

  it('should be defined', () => {
    expect(handler).toBeDefined();
    expect(repository).toBeDefined();
    expect(mail).toBeDefined();
  });

  beforeAll(() => {
    process.env.MAIL_FROM = 'testing@email.com';
  });

  it('should create a ticket', async () => {
    const command = new CreateTicketCommand({
      name: 'Test User',
      email: 'user@test.com',
      eventId: '123',
      eventName: 'Event 1',
      eventDescription: 'description of event 1',
      eventImageUrl: 'http://example.com/image.jpg',
      language: 'en',
    });
    const saveSpy = jest.spyOn(repository, 'save');
    const sendMailSpy = jest.spyOn(mail, 'sendMail');
    await expect(handler.execute(command)).resolves.toBeUndefined();
    expect(saveSpy).toHaveBeenCalledWith(
      expect.objectContaining({
        email: command.email,
        eventId: command.eventId,
        passport: expect.anything() as string,
      }),
    );
    expect(sendMailSpy).toHaveBeenCalledWith(
      expect.objectContaining({
        to: command.email,
        subject: expect.any(String) as string,
        html: expect.any(String) as string,
        from: process.env.MAIL_FROM,
      }),
    );
  });
  it('should throw an error if repository.save fails', async () => {
    const command = new CreateTicketCommand({
      name: 'Test User',
      email: 'user@test.com',
      eventId: '123',
      eventName: 'Event 1',
      eventDescription: 'description of event 1',
      eventImageUrl: 'http://example.com/image.jpg',
      language: 'en',
    });
    jest.spyOn(repository, 'save').mockImplementation(() => {
      throw new Error('Save failed');
    });
    await expect(handler.execute(command)).rejects.toThrow('Save failed');
  });

  it('should throw an error if mail.sendMail fails', async () => {
    const command = new CreateTicketCommand({
      name: 'Test User',
      email: 'user@test.com',
      eventId: '123',
      eventName: 'Event 1',
      eventDescription: 'description of event 1',
      eventImageUrl: 'http://example.com/image.jpg',
      language: 'en',
    });
    jest.spyOn(mail, 'sendMail').mockRejectedValue(new Error('Mail failed'));
    await expect(handler.execute(command)).rejects.toThrow('Mail failed');
  });
});
