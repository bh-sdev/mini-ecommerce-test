// src/products/products.service.ts
import { Controller, Injectable, NotFoundException } from '@nestjs/common';
import { GrpcMethod, GrpcService } from '@nestjs/microservices';
import { ProductsService } from './products.service';
import { CreateProductRequest, Product as GrpcProduct, ProductListResponse, ProductRequest, UpdateProductRequest, } from '../generated/product';
import { Empty } from '../generated/common';
import { ProductRepository } from './product.repository';
import { InjectRepository } from '@nestjs/typeorm';
import { from, Observable } from 'rxjs';
import { Product } from './entities/product.entity';

// @Controller("proto")
@Injectable()
@GrpcService("ProductService")
export class ProductsGrpcService {
    constructor(
        private service: ProductsService,
        @InjectRepository(ProductRepository)
        private productRepository: ProductRepository,
    ) { }

    //  -------- GRPC

    @GrpcMethod('ProductService', 'GetProduct')
    async getProduct(data: ProductRequest): Promise<GrpcProduct> {
        const product = await this.service.findOne(data.id);
        const result: GrpcProduct = {};
        result.id = (product.id);
        result.name = (product.name);
        result.description = (product.description);
        result.available_quantity = (product.availableQuantity);
        (result as any).availableQuantity = (product.availableQuantity);
        result.price = (product.price);
        return result;
    }

    @GrpcMethod('ProductService', 'DeleteProduct')
    async deleteProduct(data: ProductRequest): Promise<GrpcProduct> {
        const product = await this.service.findOne(data.id);
        const deleteResult = await this.productRepository.delete({ id: data.id });
        const result: GrpcProduct = {};
        result.id = (product.id);
        result.name = (product.name);
        result.description = (product.description);
        result.available_quantity = (product.availableQuantity);
        result.price = (product.price);
        return result;
    }


    @GrpcMethod('ProductService', 'UpdateProduct')
    updateProduct(data: UpdateProductRequest): Observable<GrpcProduct> {

        console.log("UPDATE RESQUEST: ", { ...data })

        const promise = this.service.findOne(data.id).then((product) => {

            if (data.name !== data.description)
                product.name = (data.name);
            if (undefined !== data.description)
                product.description = (data.description);
            if (undefined !== (data as any).availableQuantity)
                product.availableQuantity = ((data as any).availableQuantity);
            if (data.price !== undefined)
                product.price = (data.price);

            const updateResult = this.productRepository.update({ id: data.id }, {
                ...data,
                id: undefined,
            });
            return updateResult.then(() => product);

        }).then(product => {
            const result: GrpcProduct = {};
            result.id = (product.id);
            result.name = (product.name);
            result.description = (product.description);
            result.available_quantity = (product.availableQuantity);
            (result as any).availableQuantity = (product.availableQuantity);
            result.price = (product.price);
            return result;
        });

        return from(promise);
    }


    @GrpcMethod('ProductService', 'CreateProduct')
    createProduct(data: CreateProductRequest): Observable<GrpcProduct> {

        console.log("CREATE RESQUEST: ", { ...data })

        const product = new Product();

        product.name = (data.name);
        product.description = (data.description);
        product.availableQuantity = ((data as any).availableQuantity ?? 0);
        product.price = (data.price);

        const promise = this.service.create(product).then((product) => {
            const result: GrpcProduct = {};
            result.id = (product.id);
            result.name = (product.name);
            result.description = (product.description);
            result.available_quantity = (product.availableQuantity ?? 0);
            (result as any).availableQuantity = (product.availableQuantity);
            result.price = (product.price);
            return result;

        })
        return from(promise);
    }


    @GrpcMethod('ProductService', 'ListProducts')
    async listProduct(data: Empty): Promise<ProductListResponse> {
        const products = await this.service.findAll();
        const result: ProductListResponse = { products: [] };
        for (const product of products) {
            const _product: GrpcProduct = {};
            _product.id = (product.id);
            _product.name = (product.name);
            _product.description = (product.description);
            _product.available_quantity = (product.availableQuantity);
            (_product as any).availableQuantity = (product.availableQuantity);
            _product.price = (product.price);
            result.products.push(_product);
        }
        return result;
    }


    // @GrpcMethod('ProductService', 'ListProducts')
    // listProduct(data: Empty): Observable<ProductListResponse> {
    //     const promise = this.service.findAll().then(products => {
    //         const result: ProductListResponse = { products: [] };
    //         for (const product of products) {
    //             const _product: Product = {};
    //             _product.id = (product.id);
    //             _product.name = (product.name);
    //             _product.description = (product.description);
    //             _product.available_quantity = (product.availableQuantity);
    //             _product.price = (product.price);
    //             result.products.push(_product);
    //         }
    //         return result;
    //     });
    //     return from(promise);
    // }


    @GrpcMethod('ProductService', 'CheckProductAvailability')
    async checkAvailability(data: { productId: number; quantity: number }): Promise<{ available: boolean }> {
        const available = await this.service.checkAvailability(data.productId, data.quantity);
        return { available };
    }

}
