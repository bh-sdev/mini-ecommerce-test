import { ClientsModule, Transport } from "@nestjs/microservices";
import * as path from "path";

export const ProductRemote = ClientsModule.register([
    {
        name: 'PRODUCT_PACKAGE',
        transport: Transport.GRPC,
        options: {
            // url: 'localhost:50051',
            url: process.env.PRODUCT_SERVICE_ADDRESS ?? `0.0.0.0:${process.env.PORT ?? '50051'}`,
            package: 'proto',
            protoPath: [
                path.join(__dirname, '../proto/product.proto'),
                path.join(__dirname, '../proto/common.proto'),
            ],
        },
    },
]);