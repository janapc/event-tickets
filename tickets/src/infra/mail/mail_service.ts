import type SMTPTransport from 'nodemailer/lib/smtp-transport'
import nodemailer, { type Transporter } from 'nodemailer'
import { type IMail } from './mail'
import { logger } from '@infra/logger/logger'

export class MailService implements IMail {
  private readonly transporter: Transporter<SMTPTransport.SentMessageInfo>
  constructor() {
    this.transporter = nodemailer.createTransport({
      host: process.env.MAIL_HOST,
      port: process.env.MAIL_PORT,
      secure: false,
      auth: {
        user: process.env.MAIL_AUTH_USER,
        pass: process.env.MAIL_AUTH_PASS,
      },
    })
  }

  async sendMail(to: string, subject: string, message: string): Promise<void> {
    const result = await this.transporter.sendMail({
      from: process.env.MAIL_FROM,
      to,
      subject,
      html: message,
    })
    logger.info(`send new mail - ${result.messageId}`)
  }
}
