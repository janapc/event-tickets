import { User, USER_ROLES } from './user.entity';
import * as bcrypt from 'bcrypt';

describe('User Entity', () => {
  describe('User Creation', () => {
    it('should create a valid user with hashed password', () => {
      const email = 'test@example.com';
      const password = 'password123';
      const role = USER_ROLES.PUBLIC;
      const user = User.create({ email, password, role });
      expect(user.email).toBe(email);
      expect(user.role).toBe(role);
      expect(user.password).not.toBe(password);
      expect(bcrypt.compareSync(password, user.password)).toBeTruthy();
    });
  });

  describe('Password Validation', () => {
    it('should return true for valid password comparison', () => {
      const password = 'password123';
      const user = User.create({
        email: 'test@example.com',
        password,
        role: USER_ROLES.PUBLIC,
      });
      const isValid = User.isValidPassword(password, user.password);
      expect(isValid).toBeTruthy();
    });

    it('should return false for invalid password comparison', () => {
      const password = 'password123';
      const wrongPassword = 'wrongpassword';
      const user = User.create({
        email: 'test@example.com',
        password,
        role: USER_ROLES.PUBLIC,
      });
      const isValid = User.isValidPassword(wrongPassword, user.password);
      expect(isValid).toBeFalsy();
    });
  });
});
