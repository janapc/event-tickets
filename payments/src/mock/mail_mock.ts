import { IMail } from '@infra/mail-service/mail';

export class MailMock implements IMail {
  public mails: { to: string; message: string; subject: string }[] = [];
  async sendMail(to: string, message: string, subject: string): Promise<void> {
    this.mails.push({
      to,
      message,
      subject,
    });
  }
}
