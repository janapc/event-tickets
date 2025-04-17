export interface IMail {
  sendMail: (to: string, subject: string, message: string) => Promise<void>
}
