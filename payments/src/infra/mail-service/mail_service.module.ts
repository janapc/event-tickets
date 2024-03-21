import { MailerModule } from '@nestjs-modules/mailer';
import { Module } from '@nestjs/common';
import { IMail } from './mail';
import { MailService } from './mail_service';

@Module({
  imports: [
    MailerModule.forRoot({
      transport: {
        host: process.env.MAIL_HOST,
        secure: false,
        port: process.env.MAIL_PORT,
        auth: {
          user: process.env.MAIL_AUTH_USER,
          pass: process.env.MAIL_AUTH_PASS,
        },
        ignoreTLS: true,
      },
    }),
  ],
  exports: [IMail],
  providers: [
    {
      provide: IMail,
      useClass: MailService,
    },
  ],
})
export class MailServiceModule {}
