// src/orders/orders.module.ts
import { ClientsModule, Transport } from '@nestjs/microservices';
import { OrdersService } from './orders.service';
import { OrderItem } from './entities/order-item.entity';
import { Order } from './entities/order.entity';
import { TypeOrmModule } from '@nestjs/typeorm';
import { Module } from '@nestjs/common';
import { OrdersController } from './order.controller';
import { OrdersGrpcService } from './order.service.grpc';
import * as path from 'path';
import { OrderItemRepository } from './order-item.repository';
import { DataSource } from 'typeorm';
import { OrderRepository } from './order.repository';
import { ProductRemote } from '../remote';

@Module({
  imports: [
    ProductRemote,
    TypeOrmModule.forFeature([Order, OrderItem]),
  ],
  controllers: [OrdersController,OrdersGrpcService],
  providers: [
    OrdersService,
    OrdersGrpcService,
  
    {
      provide: OrderRepository,
      useFactory: (dataSource: DataSource) => {
        return new OrderRepository(dataSource);
      },
      inject: [DataSource],
    },
    {
      provide: OrderItemRepository,
      useFactory: (dataSource: DataSource) => {
        return new OrderItemRepository(dataSource);
      },
      inject: [DataSource],
    },
  ],
})
export class OrdersModule {}
