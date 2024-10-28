// src/orders/product-grpc.interface.ts
import { CheckProductRequest, CheckProductResponse, Product, ProductRequest } from 'product';
import { Observable } from 'rxjs';

// export interface CheckProductAvailabilityRequest {
//   productId: number;
//   quantity: number;
// }

// export interface CheckProductAvailabilityResponse {
//   available: boolean;
// }

export interface ProductService {
  CheckProductAvailability(data: CheckProductRequest): Observable<CheckProductResponse>;
  GetProduct(data: ProductRequest): Observable<Product>;
}
