import { Injectable } from '@nestjs/common';
import { Payment } from '@domain/payment';
import { IPaymentRepository } from '@domain/repository';
import { IProducer } from '@infra/message-queue/message_queue';

type InputRegisterPaymentDto = {
  name: string;
  email: string;
  eventId: string;
  eventAmount: number;
  cardNumber: string;
  securityCode: string;
  eventName: string;
  eventDescription: string;
  eventImageUrl: string;
};

@Injectable()
export class RegisterPayment {
  constructor(
    private paymentRepository: IPaymentRepository,
    private producer: IProducer,
  ) {}

  async execute(input: InputRegisterPaymentDto) {
    const payment = new Payment(
      { email: input.email, name: input.name },
      input.eventId,
    );
    const paymentId = await this.paymentRepository.save(payment);
    const message = JSON.stringify({
      ...input,
      paymentId,
    });
    await this.producer.addToQueue(message, process.env.QUEUE_PROCESS_PAYMENT);
  }
}
