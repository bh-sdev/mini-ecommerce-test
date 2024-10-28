import { Entity, PrimaryGeneratedColumn, Column, ManyToOne, JoinColumn } from 'typeorm';
import { Order } from './order.entity';
// import { Product } from '../products/entities/product.entity';

@Entity()
export class OrderItem {
  @PrimaryGeneratedColumn()
  id: number;

  @ManyToOne(() => Order, (order) => order.items)
  order: Order;
  
  @Column()
  productId: number;


  @Column()
  name?: string;


  @Column()
  description?: string;

//   @ManyToOne(() => Product)
//   @JoinColumn()
//   product: Product;
  @Column()
  price: number;

  @Column()
  quantity: number;
}
