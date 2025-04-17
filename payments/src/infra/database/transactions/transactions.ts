import { Model } from 'mongoose';
import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { ITransactionRepository } from 'src/domain/repository';
import { Transaction } from 'src/domain/transaction';

@Injectable()
export class TransactionsRepository implements ITransactionRepository {
  constructor(
    @InjectModel(Transaction.name) private transactionModel: Model<Transaction>,
  ) {}
  async save(transaction: Transaction): Promise<string> {
    const data = {
      cardNumber: transaction.cardNumber.substring(
        transaction.cardNumber.length - 4,
      ),
      amount: transaction.amount,
      errorMessage: transaction.errorMessage,
      status: transaction.status,
    };
    const createSave = new this.transactionModel(data);
    const response = await createSave.save();
    return String(response._id);
  }
}
