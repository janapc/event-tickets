import { PaymentMockRepository } from '@mock/payment_mock_repository';
import { ProducerMockMessageQueue } from '@mock/producer_mock_message_queue';
import { RegisterPayment } from '@application/register_payment';
import { ApiController } from './api.controller';
import { ProcessPayment } from '@application/process_payment';
import { TransactionMockRepository } from '@mock/transaction_mock_repository';
import { MailMock } from '@mock/mail_mock';
import { Payment } from '@domain/payment';

const mockContext = {
  getChannelRef: jest.fn().mockImplementation(() => ({
    ack: jest.fn().mockResolvedValue(true),
  })),
  getMessage: jest.fn().mockResolvedValue(true),
};

describe('ApiController', () => {
  let mockRepositoryPayment: PaymentMockRepository;
  let mockRepositoryTransaction: TransactionMockRepository;
  let mockProducer: ProducerMockMessageQueue;
  let mockMail: MailMock;
  beforeEach(() => {
    jest.restoreAllMocks();
    mockRepositoryPayment = new PaymentMockRepository();
    mockRepositoryTransaction = new TransactionMockRepository();
    mockProducer = new ProducerMockMessageQueue();
    mockMail = new MailMock();
  });
  it('should register a payment', async () => {
    const controller = new ApiController(
      new RegisterPayment(mockRepositoryPayment, mockProducer),
      new ProcessPayment(
        mockRepositoryTransaction,
        mockRepositoryPayment,
        mockProducer,
        mockMail,
      ),
    );
    await controller.register({
      email: 'banana@banana.com',
      name: 'banana banana',
      event_id: '9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d',
      event_amount: 120,
      card_number: '123012391823',
      security_code: '1234',
      event_description: 'test description',
      event_image_url: 'http://test.png',
      event_name: 'test',
      language: 'pt',
    });
    expect(mockRepositoryPayment.payments).toHaveLength(1);
    expect(mockProducer.messages).toHaveLength(1);
  });

  it('should process a payment', async () => {
    const paymentId = await mockRepositoryPayment.save(
      new Payment(
        { email: 'banana@banana.com', name: 'banana banana' },
        '123-123',
      ),
    );
    const controller = new ApiController(
      new RegisterPayment(mockRepositoryPayment, mockProducer),
      new ProcessPayment(
        mockRepositoryTransaction,
        mockRepositoryPayment,
        mockProducer,
        mockMail,
      ),
    );

    await controller.receiveMessages(
      JSON.stringify({
        eventId: '9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d',
        eventAmount: 120,
        cardNumber: '123012391823',
        securityCode: '1234',
        user: {
          email: 'banana@banana.com',
          name: 'banana banana',
        },
        paymentId,
        language: 'en',
      }),
      mockContext as any,
    );
    expect(mockRepositoryTransaction.transactions).toHaveLength(1);
    expect(mockRepositoryPayment.payments).toHaveLength(1);
    expect(mockRepositoryPayment.payments[0].transactionId).not.toBeNull();
    expect(mockProducer.messages).toHaveLength(1);
    expect(mockMail.mails).toHaveLength(1);
    expect(mockMail.mails[0].subject).toEqual('Payment Approved =)');
  });

  it('should process a payment with erro', async () => {
    const paymentId = await mockRepositoryPayment.save(
      new Payment(
        { email: 'banana@banana.com', name: 'banana banana' },
        '123-123',
      ),
    );
    const controller = new ApiController(
      new RegisterPayment(mockRepositoryPayment, mockProducer),
      new ProcessPayment(
        mockRepositoryTransaction,
        mockRepositoryPayment,
        mockProducer,
        mockMail,
      ),
    );
    await controller.receiveMessages(
      JSON.stringify({
        eventId: '9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d',
        eventAmount: 120,
        cardCumber: '123012391823',
        securityCode: '111',
        user: {
          email: 'banana@banana.com',
          name: 'banana banana',
        },
        paymentId,
        language: 'pt',
      }),
      mockContext as any,
    );
    expect(mockRepositoryTransaction.transactions).toHaveLength(1);
    expect(mockRepositoryPayment.payments).toHaveLength(1);
    expect(mockRepositoryPayment.payments[0].transactionId).not.toBeNull();
    expect(mockProducer.messages).toHaveLength(0);
    expect(mockMail.mails).toHaveLength(1);
    expect(mockMail.mails[0].subject).toEqual('Pagamento Rejeitado =(');
  });
});
