import { Lead } from '@domain/lead';
import { Controller, Post, Body, UseFilters, Get, Param } from '@nestjs/common';
import { CommandBus, QueryBus } from '@nestjs/cqrs';
import { CreateLeadCommand } from '@commands/create-lead/create-lead.command';
import { CreateLeadDto } from './dtos/create-lead.dto';
import {
  ApiConflictResponse,
  ApiCreatedResponse,
  ApiInternalServerErrorResponse,
  ApiNotFoundResponse,
  ApiOkResponse,
  ApiOperation,
  ApiTags,
} from '@nestjs/swagger';
import { HttpExceptionFilter } from './exceptions/http.exception';
import { GetLeadByEmailQuery } from '@queries/get-lead-by-email/get-lead-by-email.query';
import { GetLeadsQuery } from '@queries/get-leads/get-leads.query';
import { MessagePattern, Payload } from '@nestjs/microservices';
import { ProcessCreatedClientCommand } from '@commands/process-created-client/process-created-client.query';

@ApiTags('leads')
@Controller('leads')
@UseFilters(HttpExceptionFilter)
export class LeadController {
  constructor(
    private readonly commandBus: CommandBus,
    private readonly queryBus: QueryBus,
  ) {}

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

  @ApiOperation({ summary: 'Get lead by email' })
  @ApiOkResponse({ description: 'Lead found successfully' })
  @ApiNotFoundResponse({ description: 'Lead not found' })
  @ApiInternalServerErrorResponse({ description: 'Internal server error' })
  @Get(':email')
  async getByEmail(@Param('email') email: string): Promise<Lead> {
    return this.queryBus.execute(new GetLeadByEmailQuery(email));
  }

  @ApiOperation({ summary: 'Get all leads' })
  @ApiOkResponse({ description: 'Leads retrieved successfully' })
  @ApiInternalServerErrorResponse({ description: 'Internal server error' })
  @Get()
  async getAll(): Promise<Lead[]> {
    return this.queryBus.execute(new GetLeadsQuery());
  }

  @MessagePattern('CLIENT_CREATED_TOPIC')
  async handleClientCreated(
    @Payload()
    message: {
      email: string;
    },
  ): Promise<void> {
    return this.commandBus.execute(
      new ProcessCreatedClientCommand(message.email),
    );
  }
}
