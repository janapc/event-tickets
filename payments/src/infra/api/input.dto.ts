import {
  IsString,
  IsNotEmpty,
  IsEmail,
  MinLength,
  IsNumber,
  MaxLength,
  Min,
  IsEnum,
} from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

enum language {
  pt = 'pt',
  en = 'en',
}

export class InputRegisterPaymentDto {
  @IsString()
  @IsNotEmpty()
  @ApiProperty()
  name: string;

  @IsEmail()
  @IsNotEmpty()
  @ApiProperty()
  email: string;

  @IsString()
  @IsNotEmpty()
  @ApiProperty()
  event_id: string;

  @IsNumber()
  @Min(1)
  @ApiProperty()
  event_amount: number;

  @IsString()
  @IsNotEmpty()
  @MinLength(8)
  @MaxLength(19)
  @ApiProperty()
  card_number: string;

  @IsString()
  @IsNotEmpty()
  @MaxLength(4)
  @ApiProperty()
  security_code: string;

  @IsString()
  @IsNotEmpty()
  @ApiProperty()
  event_name: string;

  @IsString()
  @IsNotEmpty()
  @ApiProperty()
  event_description: string;

  @IsString()
  @IsNotEmpty()
  @ApiProperty()
  event_image_url: string;

  @IsEnum(language)
  @IsNotEmpty()
  @ApiProperty()
  language: string;
}
