// src/orders/order.repository.ts
import { DataSource, Repository } from 'typeorm';
import { Order } from './entities/order.entity';

// @EntityRepository(Order)
export class OrderRepository extends Repository<Order> {
    constructor(private dataSource: DataSource) {
        super(Order, dataSource.createEntityManager());
      }
}
