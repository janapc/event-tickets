import { MailService } from './mail_service'

const sendMailMock = jest.fn().mockResolvedValue({
  messageId: new Date().valueOf(),
})

jest.mock('nodemailer', () => ({
  createTransport: jest.fn().mockImplementation(() => ({
    sendMail: sendMailMock,
  })),
}))

describe('Mail Service', () => {
  it('should send to email', async () => {
    const mail = new MailService()
    await expect(
      mail.sendMail('test.test@test.com', 'test', 'description'),
    ).resolves.toBeUndefined()
    expect(sendMailMock).toHaveBeenCalledTimes(1)
  })
})
