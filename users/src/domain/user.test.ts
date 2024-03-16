import { User, type USERROLES } from '@domain/user'

describe('User', () => {
  it('should create a user', () => {
    const user = new User('email@email.com', '123123123', 'ADMIN')
    expect(user).not.toBeNull()
    expect(user.password).not.toEqual('123123123')
    expect(user.email).toEqual('email@email.com')
  })
  it('should error if fields are invalid', () => {
    const data = [
      {
        input: {
          role: '' as USERROLES,
          email: 'asd@asd.com',
          password: 'asdasd123',
        },
        error: 'the email, role and the password fields are mandatory',
      },
      {
        input: {
          role: 'PUBLIC' as USERROLES,
          email: 'asd@asd',
          password: 'asdasd123',
        },
        error: 'the email format is invalid',
      },
      {
        input: {
          role: 'PUBLIC' as USERROLES,
          email: 'asd@asd.com',
          password: 'asda',
        },
        error: 'the password field must be more than 8 characters',
      },
      {
        input: {
          role: 'PUBLICS' as USERROLES,
          email: 'asd@asd.com',
          password: 'asda123a',
        },
        error: 'the role field must be PUBLIC or ADMIN',
      },
    ]
    for (const d of data) {
      expect(
        () => new User(d.input.email, d.input.password, d.input.role),
      ).toThrow(d.error)
    }
  })
})
