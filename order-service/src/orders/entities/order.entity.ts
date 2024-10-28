import { Entity, PrimaryGeneratedColumn, Column, ManyToOne, OneToMany, JoinColumn, CreateDateColumn } from 'typeorm';
// import { Product } from '../../products/entities/product.entity';
import { OrderItem } from './order-item.entity';

@Entity()
export class Order {
  @PrimaryGeneratedColumn()
  id: number;


  @Column()
  customerId?: number;

  @Column()
  customerName: string;

  @CreateDateColumn()
  createdAt: Date;

  @OneToMany(() => OrderItem, (orderItem) => orderItem.order, { cascade: true })
  items: OrderItem[];

//   @OneToMany(() => OrderItem, (orderItem) => orderItem.order, { cascade: true })
//   items: OrderItem[];
}
