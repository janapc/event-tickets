import { Lead } from '@domain/lead';
import { Controller, Post, Body, UseFilters } from '@nestjs/common';
import { CommandBus } from '@nestjs/cqrs';
import { CreateLeadCommand } from '@commands/create-lead/create-lead.command';
import { CreateLeadDto } from './dtos/create-lead.dto';
import {
  ApiConflictResponse,
  ApiCreatedResponse,
  ApiInternalServerErrorResponse,
  ApiOperation,
  ApiTags,
} from '@nestjs/swagger';
import { HttpExceptionFilter } from './exceptions/http.exception';

@ApiTags('leads')
@Controller('leads')
@UseFilters(HttpExceptionFilter)
export class LeadController {
  constructor(private readonly commandBus: CommandBus) {}

  @ApiOperation({ summary: 'Register a new lead' })
  @ApiCreatedResponse({ description: 'Lead registered successfully' })
  @ApiConflictResponse({ description: 'Lead already exists' })
  @ApiInternalServerErrorResponse({ description: 'Internal server error' })
  @Post()
  async create(@Body() body: CreateLeadDto): Promise<Lead> {
    return this.commandBus.execute(
      new CreateLeadCommand(body.email, body.converted, body.language),
    );
  }
}
