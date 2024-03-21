import { Transaction } from './transaction';

describe('Transaction', () => {
  it('should create a transaction', () => {
    const transaction = new Transaction('192837429102345', '1234', 128.5);
    expect(transaction.id).not.toBeNull();
    expect(transaction.status).toEqual('SUCCESS');
    expect(transaction.errorMessage).toEqual('');
  });

  it('should error if credit card is invalid', () => {
    const data = [
      {
        cardNumber: '1929292831',
        securityCode: '111',
        amount: 60.98,
        expectedMessage: 'the credit card is invalid',
      },
      {
        cardNumber: '11111111',
        securityCode: '12344',
        amount: 60.98,
        expectedMessage: 'the credit card is invalid',
      },
    ];
    for (const d of data) {
      const transaction = new Transaction(
        d.cardNumber,
        d.securityCode,
        d.amount,
      );
      expect(transaction.errorMessage).toEqual(d.expectedMessage);
      expect(transaction.status).toEqual('FAILED');
    }
  });
});
