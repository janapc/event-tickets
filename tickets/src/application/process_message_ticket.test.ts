import { DatabaseMock } from './mock/database_mock'
import { MailMock } from './mock/mail_mock'
import { ProcessMessageTicket } from './process_message_ticket'

describe('Process Message Ticket', () => {
  it('should save ticket on db and send email ticket', async () => {
    const mail = new MailMock()
    const database = new DatabaseMock()
    const application = new ProcessMessageTicket(database, mail)
    const message = JSON.stringify({
      name: 'teste',
      email: 'teste@test.com',
      eventId: '12312-123',
      eventName: 'show banana',
      eventDescription: 'show banana',
      eventImageUrl: 'http://image.png',
      language: 'pt',
    })
    await expect(application.execute(message)).resolves.toBeUndefined()
    expect(mail.mails).toHaveLength(1)
    expect(mail.mails[0].subject).toEqual('Seu ticket chegou =)')
    expect(database.tickets).toHaveLength(1)
  })
})
