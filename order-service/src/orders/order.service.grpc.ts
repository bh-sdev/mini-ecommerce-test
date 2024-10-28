// src/orders/orders.service.ts
import { Injectable, NotFoundException, UnprocessableEntityException } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { OrdersService } from './orders.service';
import { Empty } from '../generated/common';
import { CreateOrderRequest, Order as GrpcOrder, OrderItem as GrpcOrderItem, OrderListResponse, OrderRequest, } from '../generated/order';
import { InjectRepository } from '@nestjs/typeorm';
import { OrderRepository } from './order.repository';
import { from, Observable } from 'rxjs';
import { Order } from './entities/order.entity';
import { OrderItem } from './entities/order-item.entity';
import { Product } from 'product';

@Injectable()
export class OrdersGrpcService {
    constructor(
        private service: OrdersService,
        @InjectRepository(OrderRepository)
        private orderRepository: OrderRepository,
    ) { }

    @GrpcMethod('OrderService', 'GetOrder')
    async getOrder(data: OrderRequest): Promise<GrpcOrder> {
        const order = await this.orderRepository.findOneBy({ id: data.id });
        const result: GrpcOrder = {
            items: []
        };
        for (const item of order.items) {
            const entry: GrpcOrderItem = {};
            entry.product_id = item.productId;
            entry.name = (item.name);
            entry.description = (item.description);
            entry.quantity = (item.quantity);
            entry.price = (item.price);
            result.items.push(entry);
        }
        return result;
    }

    @GrpcMethod('OrderService', 'CreateOrder')
    async createOrder(data: CreateOrderRequest): Promise<GrpcOrder> {

        console.log("CREATE ORDER REQUEST: ", { ...data });

        const order = new Order();
        order.customerId = data.customerId;
        order.customerName = data.customerName?? '[N/A]';
        order.items = [];
        for (const itemDto of data.items) {
            const productId: number = ((itemDto as any).product_id) ?? ((itemDto as any)).productId;

            const response = await this.service.checkAvailability(productId, itemDto.quantity);

            if (!response.available) {
                throw new NotFoundException(`Product with ID ${productId} is not available in required quantity.`);
            }

            var product: Product | undefined = undefined;
            try {
                product = await this.service.getProduct(productId);
            } catch (err) { }

            const orderItem = new OrderItem();
            orderItem.productId = productId;
            orderItem.quantity = itemDto.quantity;
            orderItem.name = product?.name;
            orderItem.description = product?.description;
            order.items.push(orderItem);
        }

        const _order = await this.orderRepository.save(order);


        // 
        // const order = await this.service.findOne(data.id);
        const result: GrpcOrder = {
            items: []
        };
        for (const item of _order.items) {
            const entry: GrpcOrderItem = {};
            entry.product_id = (item.productId);
            (entry as any).productId = (item.productId);
            entry.name = (item.name);
            entry.description = (item.description);
            entry.quantity = (item.quantity);
            entry.price = (item.price);
            result.items.push(entry);
        }
        return result;
    }


    // @GrpcMethod('OrderService', 'UpdateOrder')
    // async updateOrder(data: UpdateOrderRequest): Promise<Order> {
    //     throw new UnprocessableEntityException("We can't update an order")
    //     return result;
    // }


    @GrpcMethod('OrderService', 'ListOrders')
    listOrder(data: Empty): Observable<OrderListResponse> {
        const promise = this.orderRepository.find().then(orders => {
            const result: OrderListResponse = {
                orders: []
            };
            for (const order of orders) {
                const _order: GrpcOrder = {
                    items: []
                };
                for (const item of order.items) {
                    const entry: GrpcOrderItem = {};
                    entry.product_id = (item.productId);
                    entry.name = (item.name);
                    entry.description = (item.description);
                    entry.quantity = (item.quantity);
                    entry.price = (item.price);
                    _order.items.push(entry);
                }
                result.orders.push(_order);
            }
            return result;
        });
        return from(promise);
    }


}
