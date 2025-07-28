import { Injectable, Logger } from '@nestjs/common';
import { MailerService } from '@nestjs-modules/mailer';
import { MailAbstract, SendMailProps } from '@domain/mail';

@Injectable()
export class MailService implements MailAbstract {
  private logger = new Logger(MailService.name);
  constructor(private readonly mailerService: MailerService) {}

  async sendMail(props: SendMailProps) {
    await this.mailerService.sendMail({
      to: props.to,
      from: props.from,
      subject: props.subject,
      text: props.text,
      html: props.html,
    });
    this.logger.log('Mail sent successfully');
  }
}
