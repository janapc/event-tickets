import { PaymentMockRepository } from '@mock/payment_mock_repository';
import { ProducerMockMessageQueue } from '@mock/producer_mock_message_queue';
import { ProcessPayment } from './process_payment';
import { TransactionMockRepository } from '@mock/transaction_mock_repository';
import { MailMock } from '@mock/mail_mock';
import { Payment } from '@domain/payment';

describe('Process Payment', () => {
  let transactionMock: TransactionMockRepository;
  let paymentMock: PaymentMockRepository;
  let producerMock: ProducerMockMessageQueue;
  let mailMock: MailMock;

  beforeEach(() => {
    transactionMock = new TransactionMockRepository();
    paymentMock = new PaymentMockRepository();
    producerMock = new ProducerMockMessageQueue();
    mailMock = new MailMock();
  });

  afterEach(() => {
    jest.restoreAllMocks();
  });

  it('should process a payment message', async () => {
    const spyTransactionSave = jest.spyOn(transactionMock, 'save');
    const spyPaymentUpdate = jest.spyOn(paymentMock, 'updateTransaction');
    const spyAddToQueue = jest.spyOn(producerMock, 'addToQueue');
    const spySendMail = jest.spyOn(mailMock, 'sendMail');
    const paymentId = await paymentMock.save(
      new Payment(
        { email: 'banana@banana.com', name: 'banana banana' },
        '9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d',
      ),
    );
    const application = new ProcessPayment(
      transactionMock,
      paymentMock,
      producerMock,
      mailMock,
    );
    await application.execute({
      email: 'banana@banana.com',
      name: 'banana banana',
      eventId: '9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d',
      eventAmount: 120,
      cardNumber: '123012391823',
      securityCode: '1234',
      paymentId,
      eventName: 'test',
      eventDescription: 'test description',
      eventImageUrl: 'http://image.png',
    });
    expect(transactionMock.transactions).toHaveLength(1);
    expect(paymentMock.payments).toHaveLength(1);
    expect(paymentMock.payments[0].transactionId).not.toBeNull();
    expect(producerMock.messages).toHaveLength(1);
    expect(mailMock.mails).toHaveLength(1);
    expect(mailMock.mails[0].subject).toEqual('Payment Approved =)');
    expect(spyTransactionSave).toHaveBeenCalledTimes(1);
    expect(spyAddToQueue).toHaveBeenCalledTimes(1);
    expect(spyPaymentUpdate).toHaveBeenCalledTimes(1);
    expect(spySendMail).toHaveBeenCalledTimes(1);
  });

  it('should process a payment with error', async () => {
    const spyTransactionSave = jest.spyOn(transactionMock, 'save');
    const spyPaymentUpdate = jest.spyOn(paymentMock, 'updateTransaction');
    const spyAddToQueue = jest.spyOn(producerMock, 'addToQueue');
    const spySendMail = jest.spyOn(mailMock, 'sendMail');
    const paymentId = await paymentMock.save(
      new Payment(
        { email: 'banana@banana.com', name: 'banana banana' },
        '9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d',
      ),
    );
    const application = new ProcessPayment(
      transactionMock,
      paymentMock,
      producerMock,
      mailMock,
    );
    await application.execute({
      email: 'banana@banana.com',
      name: 'banana banana',
      eventId: '9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d',
      eventAmount: 120,
      cardNumber: '123012391823',
      securityCode: '111',
      paymentId,
      eventName: 'test',
      eventDescription: 'test description',
      eventImageUrl: 'http://image.png',
    });
    expect(transactionMock.transactions).toHaveLength(1);
    expect(paymentMock.payments).toHaveLength(1);
    expect(paymentMock.payments[0].transactionId).not.toBeNull();
    expect(producerMock.messages).toHaveLength(0);
    expect(mailMock.mails).toHaveLength(1);
    expect(mailMock.mails[0].subject).toEqual('Payment Rejected =(');
    expect(spyTransactionSave).toHaveBeenCalledTimes(1);
    expect(spyAddToQueue).toHaveBeenCalledTimes(0);
    expect(spyPaymentUpdate).toHaveBeenCalledTimes(1);
    expect(spySendMail).toHaveBeenCalledTimes(1);
  });
});
