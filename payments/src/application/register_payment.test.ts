import { PaymentMockRepository } from '@mock/payment_mock_repository';
import { ProducerMockMessageQueue } from '@mock/producer_mock_message_queue';
import { RegisterPayment } from './register_payment';

describe('Register Payment', () => {
  it('should register a payment', async () => {
    const mockRepository = new PaymentMockRepository();
    const mockProducer = new ProducerMockMessageQueue();
    const application = new RegisterPayment(mockRepository, mockProducer);
    await application.execute({
      email: 'banana@banana.com',
      name: 'banana banana',
      eventId: '9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d',
      eventAmount: 120,
      cardNumber: '123012391823',
      securityCode: '1234',
      eventName: 'show banana',
      eventDescription: 'teste ',
      eventImageUrl: 'http://image.png',
    });
    expect(mockRepository.payments).toHaveLength(1);
    expect(mockProducer.messages).toHaveLength(1);
  });
});
