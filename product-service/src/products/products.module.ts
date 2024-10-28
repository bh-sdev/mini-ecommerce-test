import { Module } from '@nestjs/common';
import { ProductsService } from './products.service';
import { ProductsController } from './products.controller';
import { ProductRepository } from './product.repository';
import { Product } from './entities/product.entity';
import { TypeOrmModule } from '@nestjs/typeorm';
import { DataSource } from 'typeorm';
// import { ProductsGrpcService } from './products.grpc';
import { ClientsModule, Transport } from '@nestjs/microservices';
import * as path from 'path';
import { ProductsGrpcService } from './products.service.grpc';



@Module({
  imports: [
    // ClientsModule.register([
    //   {
    //     name: 'PRODUCT_PACKAGE',
    //     transport: Transport.GRPC,
    //     options: {
    //       package: 'product',
    //       protoPath: path.join(__dirname, '../../proto/product.proto'),
    //       url: `0.0.0.0:${process.env.PORT??'50051'}`,
    //     },
    //   },
    // ]),
    TypeOrmModule.forFeature([Product, ProductRepository])
  ],
  controllers: [ProductsController,ProductsGrpcService],
  // providers: [ProductsService],
  providers: [
    ProductsService,
    ProductsGrpcService,
    {
      provide: ProductRepository,
      useFactory: (dataSource: DataSource) => {
        return new ProductRepository(dataSource);
      },
      inject: [DataSource],
    },
  ],
  exports: [ProductsService,ProductsGrpcService],
})
export class ProductsModule {}
