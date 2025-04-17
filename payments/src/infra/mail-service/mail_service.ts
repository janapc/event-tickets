import { MailerService } from '@nestjs-modules/mailer';
import { Injectable, Logger } from '@nestjs/common';
import { IMail } from './mail';

@Injectable()
export class MailService implements IMail {
  private readonly logger = new Logger(MailService.name);

  constructor(private mailerService: MailerService) {}
  async sendMail(to: string, message: string, subject: string): Promise<void> {
    await this.mailerService.sendMail({
      to,
      from: process.env.MAIL_FROM,
      subject,
      text: message,
    });
    this.logger.log(`message sent to mail: ${to}`);
  }
}
