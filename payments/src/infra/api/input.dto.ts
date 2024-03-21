import {
  IsString,
  IsNotEmpty,
  IsEmail,
  MinLength,
  IsNumber,
  MaxLength,
  Min,
} from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

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
  @MinLength(36)
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
}
