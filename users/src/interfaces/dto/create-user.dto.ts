import { USER_ROLES } from '@domain/user.entity';
import { ApiProperty } from '@nestjs/swagger';
import {
  IsEmail,
  IsEnum,
  IsNotEmpty,
  IsString,
  MinLength,
} from 'class-validator';

export class CreateUserDto {
  @ApiProperty({
    description: 'The email of the user',
    example: 'test@test.com',
    required: true,
  })
  @IsEmail()
  @IsNotEmpty()
  email: string;

  @ApiProperty({
    description: 'The password of the user',
    example: 'password',
    required: true,
  })
  @IsString()
  @IsNotEmpty()
  @MinLength(8)
  password: string;

  @ApiProperty({
    description: 'The role of the user',
    example: USER_ROLES.ADMIN,
    required: true,
  })
  @IsEnum(USER_ROLES)
  @IsNotEmpty()
  role: USER_ROLES;
}
