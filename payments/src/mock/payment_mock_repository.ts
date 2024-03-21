import { Payment } from '@domain/payment';
import { IPaymentRepository } from '@domain/repository';

export class PaymentMockRepository implements IPaymentRepository {
  public payments: Payment[] = [];

  async save(payment: Payment): Promise<string> {
    const id = new Date().valueOf().toString();
    payment.id = id;
    this.payments.push(payment);
    return payment.id;
  }

  async updateTransaction(
    transactionId: string,
    paymentId: string,
  ): Promise<void> {
    const index = this.payments.findIndex((payment) => payment.id == paymentId);
    this.payments[index] = {
      ...this.payments[index],
      transactionId: transactionId,
    };
  }
}
