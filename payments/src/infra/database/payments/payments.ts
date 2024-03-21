import { Model } from 'mongoose';
import { Injectable } from '@nestjs/common';
import { InjectModel } from '@nestjs/mongoose';
import { IPaymentRepository } from 'src/domain/repository';
import { Payment } from 'src/domain/payment';

@Injectable()
export class PaymentsRepository implements IPaymentRepository {
  constructor(
    @InjectModel(Payment.name) private paymentModel: Model<Payment>,
  ) {}

  async updateTransaction(
    transactionId: string,
    paymentId: string,
  ): Promise<void> {
    await this.paymentModel.findByIdAndUpdate(paymentId, { transactionId });
  }

  async save(payment: Payment): Promise<string> {
    const createSave = new this.paymentModel(payment);
    const response = await createSave.save();
    return response._id.toString();
  }
}
