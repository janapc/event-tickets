import { Injectable } from '@nestjs/common';
import { MailerService } from '@nestjs-modules/mailer';
import { MailAbstract, SendMailProps } from '@domain/mail';

@Injectable()
export class MailService implements MailAbstract {
  constructor(private readonly mailerService: MailerService) {}

  async sendMail(props: SendMailProps) {
    await this.mailerService.sendMail({
      to: props.to,
      from: props.from,
      subject: props.subject,
      text: props.text,
      html: props.html,
    });
  }
}
