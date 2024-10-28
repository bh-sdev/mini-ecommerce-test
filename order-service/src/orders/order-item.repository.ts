// src/orders/order.repository.ts
import { DataSource, Repository } from 'typeorm';
import { OrderItem } from './entities/order-item.entity';

// @EntityRepository(OrderItem)
export class OrderItemRepository extends Repository<OrderItem> {
    constructor(private dataSource: DataSource) {
        super(OrderItem, dataSource.createEntityManager());
      }
}
