import { Body, Controller, Logger, Post } from '@nestjs/common';
import { RegisterPayment } from '@application/register_payment';
import { InputRegisterPaymentDto } from './input.dto';
import {
  Ctx,
  MessagePattern,
  Payload,
  RmqContext,
} from '@nestjs/microservices';
import { ProcessPayment } from '@application/process_payment';

@Controller('payments')
export class ApiController {
  private readonly logger = new Logger(ApiController.name);
  constructor(
    private readonly registerPayment: RegisterPayment,
    private readonly processPayment: ProcessPayment,
  ) {}

  @Post()
  async register(@Body() body: InputRegisterPaymentDto) {
    const input = {
      name: body.name,
      email: body.email,
      eventId: body.event_id,
      eventAmount: body.event_amount,
      cardNumber: body.card_number,
      securityCode: body.security_code,
      eventName: body.event_name,
      eventDescription: body.event_description,
      eventImageUrl: body.event_image_url,
      language: body.language,
    };
    await this.registerPayment.execute(input);
  }

  @MessagePattern('process_payment')
  async receiveMessages(@Payload() data: string, @Ctx() context: RmqContext) {
    try {
      const message = JSON.parse(data);
      this.logger.log(`receive payment id: ${message.paymentId}`);
      const channel = context.getChannelRef();
      const originalMsg = context.getMessage();

      await this.processPayment.execute(message);

      channel.ack(originalMsg);
      this.logger.log(`message processed successfully`);
      return data;
    } catch (error) {
      this.logger.error(`error receive message: ${error}`);
    }
  }
}
