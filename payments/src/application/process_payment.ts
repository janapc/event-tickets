import { IPaymentRepository, ITransactionRepository } from '@domain/repository';
import { Transaction } from '@domain/transaction';
import { IMail } from '@infra/mail-service/mail';
import { IProducer } from '@infra/message-queue/message_queue';
import { Injectable } from '@nestjs/common';

type InputProcessPaymentDto = {
  name: string;
  email: string;
  eventId: string;
  eventAmount: number;
  cardNumber: string;
  securityCode: string;
  paymentId: string;
  eventName: string;
  eventDescription: string;
  eventImageUrl: string;
};

@Injectable()
export class ProcessPayment {
  constructor(
    private transactionRepository: ITransactionRepository,
    private paymentRepository: IPaymentRepository,
    private producer: IProducer,
    private mailService: IMail,
  ) {}
  async execute(input: InputProcessPaymentDto): Promise<void> {
    const transaction = new Transaction(
      input.cardNumber,
      input.securityCode,
      input.eventAmount,
    );
    const transactionId = await this.transactionRepository.save(transaction);
    await this.paymentRepository.updateTransaction(
      transactionId,
      input.paymentId,
    );
    if (transaction.status === 'SUCCESS') {
      await this.producer.addToQueue(
        JSON.stringify({
          name: input.name,
          email: input.email,
          eventId: input.eventId,
          eventName: input.eventName,
          eventDescription: input.eventDescription,
          eventImageUrl: input.eventImageUrl,
        }),
        process.env.QUEUE_SUCCESS_PAYMENT,
      );
      await this.mailService.sendMail(
        input.email,
        `Hello ${input.name}.\n Your payment has been approved.`,
        'Payment Approved =)',
      );
    } else {
      await this.mailService.sendMail(
        input.email,
        `Hello ${input.name}.\n Your payment has been rejected, try again in us site`,
        'Payment Rejected =(',
      );
    }
  }
}
