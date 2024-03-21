export abstract class IMail {
  abstract sendMail(
    to: string,
    message: string,
    subject: string,
  ): Promise<void>;
}
