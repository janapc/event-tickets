import { DatabaseMockRepository } from './mock/database_mock_repostory'
import { Lead } from '@domain/lead'
import { ConsumeMessagesQueue } from './consume_messages_queue'

describe('Consume Messages Queue', () => {
  it('should consume message from queue and update lead', async () => {
    const repository = new DatabaseMockRepository()
    await repository.save(new Lead('test.test@test.com', false))
    const application = new ConsumeMessagesQueue(repository)
    await application.execute(
      JSON.stringify({
        email: 'test.test@test.com',
        hasClient: true,
      }),
    )
    expect(repository.leads[0].converted).toBeTruthy()
  })
})
