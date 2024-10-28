import { Injectable, Inject, NotFoundException } from '@nestjs/common';
import { CreateOrderDto } from './dto/create-order.dto';
import { Order } from './entities/order.entity';
import { OrderItem } from './entities/order-item.entity';
import { Repository } from 'typeorm';
import { ClientGrpc, GrpcService } from '@nestjs/microservices';
import { lastValueFrom } from 'rxjs';
import { InjectRepository } from '@nestjs/typeorm';
import { ProductService } from './product-grpc.interface';

@GrpcService("OrderService")
@Injectable()
export class OrdersService {
  private productService: ProductService;

  constructor(
    @Inject('PRODUCT_PACKAGE') private client: ClientGrpc,
    @InjectRepository(Order) private orderRepository: Repository<Order>,
    @InjectRepository(OrderItem) private orderItemRepository: Repository<OrderItem>,
  ) {
    this.productService = this.client.getService<ProductService>('ProductService');
  }


  public async checkAvailability(
    productId: number, quantity: number
  ){
    const response = await lastValueFrom(
      this.productService.CheckProductAvailability({ productId, quantity })
    );
    return response;
  }

  public async getProduct(
    productId: number
  ){
    const response = await lastValueFrom(
      this.productService.GetProduct({ id: productId })
    );
    return response;
  }

  async createOrder(createOrderDto: CreateOrderDto): Promise<Order> {
    const order = new Order();
    order.customerName = createOrderDto.customerName;
    order.items = [];

    for (const itemDto of createOrderDto.items) {
      const response = await lastValueFrom(
        this.productService.CheckProductAvailability({ productId: itemDto.productId, quantity: itemDto.quantity })
      );

      if (!response.available) {
        throw new NotFoundException(`Product with ID ${itemDto.productId} is not available in required quantity.`);
      }

      const orderItem = new OrderItem();
      orderItem.productId = itemDto.productId;
      orderItem.quantity = itemDto.quantity;
      order.items.push(orderItem);
    }

    return this.orderRepository.save(order);
  }


  async findOne(id: number): Promise<Order> {
    const order = await this.orderRepository.findOne({ where: { id } })//findOne(id);
    if (!order) throw new NotFoundException(`Order with ID ${id} not found`);
    return order;
  }

  // Existing code...

  async getOrders(): Promise<Order[]> {
    return this.orderRepository.find({ relations: ['items', 'items.product'] });
  }
}
