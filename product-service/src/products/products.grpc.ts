// src/products/products.grpc.ts
import { Injectable } from '@nestjs/common';
import { GrpcMethod } from '@nestjs/microservices';
import { ProductsService } from './products.service';

@Injectable()
export class ProductsGrpcService {
  constructor(private readonly productsService: ProductsService) {}

  @GrpcMethod('ProductService', 'CheckProductAvailability')
  async checkAvailability(data: { productId: number; quantity: number }): Promise<{ available: boolean }> {
    const available = await this.productsService.checkAvailability(data.productId, data.quantity);
    return { available };
  }
}
