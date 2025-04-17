export class Transaction {
  id?: string;
  cardNumber: string;
  amount: number;
  securityCode: string;
  errorMessage?: string;
  status: 'SUCCESS' | 'FAILED';

  constructor(cardNumber: string, securityCode: string, amount: number) {
    this.cardNumber = cardNumber;
    this.securityCode = securityCode;
    this.amount = amount;
    if (this.isFraud()) {
      this.errorMessage = 'the credit card is invalid';
      this.status = 'FAILED';
    } else {
      this.errorMessage = '';
      this.status = 'SUCCESS';
    }
  }

  isFraud(): boolean {
    if (this.securityCode === '111') {
      return true;
    }
    if (this.cardNumber === '11111111') {
      return true;
    }
    return false;
  }
}
