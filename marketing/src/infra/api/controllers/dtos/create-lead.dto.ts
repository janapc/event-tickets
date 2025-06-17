import { ApiProperty } from '@nestjs/swagger';
import { IsBoolean, IsEmail, IsNotEmpty, IsString } from 'class-validator';

export class CreateLeadDto {
  @ApiProperty({
    description: 'email of lead',
    required: true,
    example: 'email@email.com',
  })
  @IsEmail()
  email: string;

  @ApiProperty({
    description: 'Language of the lead',
    required: true,
    example: 'pt-br',
  })
  @IsNotEmpty()
  @IsString()
  language: string;

  @ApiProperty({
    description: 'Indicates if the lead has been converted',
    required: true,
    example: false,
  })
  @IsNotEmpty()
  @IsBoolean()
  converted: boolean;
}
