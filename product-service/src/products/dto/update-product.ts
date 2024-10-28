// src/products/dto/update-product.dto.ts
import { PartialType } from '@nestjs/mapped-types';
import { CreateProductDto } from './create-product';

export class UpdateProductDto extends PartialType(CreateProductDto) {}
