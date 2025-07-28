import { Module } from '@nestjs/common';
import { MailAbstract } from '@domain/mail';
import { MailService } from './service.mail';

@Module({
  imports: [],
  providers: [
    {
      provide: MailAbstract,
      useClass: MailService,
    },
  ],
  exports: [
    {
      provide: MailAbstract,
      useClass: MailService,
    },
  ],
})
export class MailModule {}
