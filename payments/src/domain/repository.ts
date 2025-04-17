import { Payment } from './payment';
import { Transaction } from './transaction';

export abstract class ITransactionRepository {
  abstract save(transaction: Transaction): Promise<string>;
}

export abstract class IPaymentRepository {
  abstract save(payment: Payment): Promise<string>;
  abstract updateTransaction(
    transactionId: string,
    paymentId: string,
  ): Promise<void>;
}
