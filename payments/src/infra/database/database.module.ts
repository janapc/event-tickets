import { Module } from '@nestjs/common';
import { ConfigModule, ConfigService } from '@nestjs/config';
import { MongooseModule } from '@nestjs/mongoose';
import { PaymentsRepository } from './payments/payments';
import { PaymentSchema, Payment } from './payments/schema';
import {
  IPaymentRepository,
  ITransactionRepository,
} from 'src/domain/repository';
import { Transaction, TransactionSchema } from './transactions/schema';
import { TransactionsRepository } from './transactions/transactions';

@Module({
  imports: [
    ConfigModule.forRoot(),
    MongooseModule.forRootAsync({
      imports: [ConfigModule],
      useFactory: async (configService: ConfigService) => ({
        uri: configService.get<string>('MONGODB_URI'),
      }),
      inject: [ConfigService],
    }),
    MongooseModule.forFeature([
      { name: Payment.name, schema: PaymentSchema },
      { name: Transaction.name, schema: TransactionSchema },
    ]),
  ],
  providers: [
    {
      provide: IPaymentRepository,
      useClass: PaymentsRepository,
    },
    {
      provide: ITransactionRepository,
      useClass: TransactionsRepository,
    },
  ],
  exports: [IPaymentRepository, ITransactionRepository],
})
export class DatabaseModule {}
