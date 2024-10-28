import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Order } from './orders/entities/order.entity';
import { OrderItem } from './orders/entities/order-item.entity';
import { ClientsModule, Transport } from '@nestjs/microservices';
import * as path from 'path';
import { OrdersModule } from './orders/orders.module';
import { ProductRemote } from './remote';

@Module({
  imports: [

    // ClientsModule.register([
    //   {
    //     name: 'ORDER_PACKAGE',
    //     transport: Transport.GRPC,
    //     options: {
    //       package: 'order',
    //       protoPath: path.join(__dirname, '../../proto/order.proto'),
    //       url: `0.0.0.0:${process.env.PORT??'50052'}`,
    //     },
    //   },
    // ]),

    ProductRemote,
    TypeOrmModule.forRoot({
      type: 'postgres',
      url: process.env.DATABASE_URL,
      synchronize: true, // for development only
      retryAttempts: 10, // Number of retries
      retryDelay: 5000,  // Delay between retries in ms
      entities: [Order,OrderItem],
    }),
    OrdersModule,
  ],
  controllers: [AppController],
  providers: [AppService],
})
export class AppModule {}
