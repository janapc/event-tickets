import { Payment } from './payment';

describe('Payment', () => {
  it('should create a payment', () => {
    const payment = new Payment(
      { name: 'test', email: 'test@test.com' },
      '9b1deb4d-3b7d-4bad-9bdd-2b0d7b3dcb6d',
      'transaction_1',
    );
    expect(payment.transactionId).toEqual('transaction_1');
  });
});
