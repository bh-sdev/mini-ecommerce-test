// src/products/product.repository.ts
import { DataSource, Repository } from 'typeorm';
import { Product } from './entities/product.entity';

// @EntityRepository(Product)
export class ProductRepository extends Repository<Product> {
    constructor(private dataSource: DataSource) {
        super(Product, dataSource.createEntityManager());
      }
}
