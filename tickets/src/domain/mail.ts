export interface SendMailProps {
  to: string;
  from: string;
  subject: string;
  text?: string;
  html?: string;
}

export abstract class MailAbstract {
  abstract sendMail: (props: SendMailProps) => Promise<void>;
}
