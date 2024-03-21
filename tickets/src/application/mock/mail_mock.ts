import { type IMail } from '@infra/mail/mail'

export class MailMock implements IMail {
  public mails: Array<{ to: string; subject: string; message: string }> = []
  async sendMail(to: string, subject: string, message: string): Promise<void> {
    this.mails.push({
      to,
      subject,
      message,
    })
  }
}
