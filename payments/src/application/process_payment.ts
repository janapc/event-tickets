import { IPaymentRepository, ITransactionRepository } from '@domain/repository';
import { Transaction } from '@domain/transaction';
import {
  IntlApprovedPayment,
  IntlRejectedPayment,
} from '@infra/mail-service/intl';
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
  language: string;
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
          language: input.language,
        }),
        process.env.QUEUE_SUCCESS_PAYMENT,
      );
      const email = IntlApprovedPayment(input.language, input.name);
      await this.mailService.sendMail(
        input.email,
        email.message,
        email.subject,
      );
    } else {
      const email = IntlRejectedPayment(input.language, input.name);
      await this.mailService.sendMail(
        input.email,
        email.message,
        email.subject,
      );
    }
  }
}
