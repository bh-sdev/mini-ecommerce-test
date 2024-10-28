import { Module } from '@nestjs/common';
import { AppController } from './app.controller';
import { AppService } from './app.service';
import { Product } from './products/entities/product.entity';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ProductsModule } from './products/products.module';

@Module({
  imports: [
    TypeOrmModule.forRoot({
      type: 'postgres',
      url: process.env.DATABASE_URL,
      // host: '127.0.0.1',
      // port: 5433,
      // username: 'postgres',
      // password: 'postgres', // replace with actual password
      // database: 'super_dev_ecommerce_product_service',
      entities: [Product],
      synchronize: true, // for development only
      retryAttempts: 10, // Number of retries
      retryDelay: 5000,  // Delay between retries in ms
    }),
    ProductsModule,
  ],
  controllers: [AppController],
  providers: [AppService],
},)
export class AppModule { }
