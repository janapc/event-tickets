/** @type {import('ts-jest').JestConfigWithTsJest} */

module.exports = {
	preset: 'ts-jest',
	testEnvironment: 'node',
	moduleNameMapper: {
		'^@domain/(.*)$': '<rootDir>/src/domain/$1',
		'^@infra/(.*)$': '<rootDir>/src/infra/$1',
		'^@application/(.*)$': '<rootDir>/src/application/$1',
	}
};