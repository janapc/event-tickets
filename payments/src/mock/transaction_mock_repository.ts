import { Transaction } from '@domain/transaction';
import { ITransactionRepository } from '@domain/repository';

export class TransactionMockRepository implements ITransactionRepository {
  public transactions: Transaction[] = [];
  async save(transaction: Transaction): Promise<string> {
    const id = new Date().valueOf().toString();
    transaction.id = id;
    this.transactions.push(transaction);
    return transaction.id;
  }
}
