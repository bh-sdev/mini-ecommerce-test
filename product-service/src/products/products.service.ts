// src/products/products.service.ts
import { Injectable, NotFoundException } from '@nestjs/common';
import { InjectRepository } from '@nestjs/typeorm';
import { ProductRepository } from './product.repository';
import { Product } from './entities/product.entity';
import { CreateProductDto } from './dto/create-product';
import { UpdateProductDto } from './dto/update-product';

@Injectable()
export class ProductsService {
    constructor(
        @InjectRepository(ProductRepository)
        private productRepository: ProductRepository,
    ) { }

    create(createProductDto: CreateProductDto): Promise<Product> {
        const product = this.productRepository.create(createProductDto);
        return this.productRepository.save(product);
    }

    findAll(): Promise<Product[]> {
        return this.productRepository.find();
    }

    async paginate(
        search?: string,
        minPrice?: number,
        maxPrice?: number,
        page = 1,
        limit = 10
    ): Promise<{ data: Product[]; total: number; page: number; limit: number }> {
        const query = this.productRepository.createQueryBuilder('product');

        // Search by name or description
        if (search) {
            query.andWhere('(product.name ILIKE :search OR product.description ILIKE :search)', {
                search: `%${search}%`,
            });
        }

        // Filter by price range
        if (minPrice) {
            query.andWhere('product.price >= :minPrice', { minPrice });
        }
        if (maxPrice) {
            query.andWhere('product.price <= :maxPrice', { maxPrice });
        }

        // Pagination
        query.skip((page - 1) * limit).take(limit);

        // Get total count for pagination metadata
        const [data, total] = await query.getManyAndCount();

        return {
            data,
            total,
            page,
            limit,
        };
    }

    async findOne(id: number): Promise<Product> {
        const product = await this.productRepository.findOne({ where: { id } })//findOne(id);
        if (!product) throw new NotFoundException(`Product with ID ${id} not found`);
        return product;
    }

    async update(id: number, updateProductDto: UpdateProductDto): Promise<Product> {
        await this.productRepository.update(id, updateProductDto);
        return this.findOne(id);
    }

    async remove(id: number): Promise<void> {
        const product = await this.findOne(id);
        await this.productRepository.remove(product);
    }


    async checkAvailability(productId: number, quantity: number): Promise<boolean> {
        const product = await this.productRepository.findOneBy({ id: productId });
        if (!product) throw new NotFoundException(`Product with ID ${productId} not found`);
        return product.availableQuantity >= quantity;
    }

}
