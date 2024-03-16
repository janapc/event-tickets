export const errorsMap = new Map([
  ['the email, role and the password fields are mandatory', 400],
  ['the email format is invalid', 400],
  ['the password field must be more than 8 characters', 400],
  ['the role field must be PUBLIC or ADMIN', 400],
  ['we can not find the user.Please try again', 404],
  ['your email or password is incorrect.Please try again', 400],
  ['email already registered', 409],
])
