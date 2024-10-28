import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import * as path from 'path';

async function bootstrap() {
  // const app = await NestFactory.create(AppModule);
  // await app.listen(process.env.PORT ?? 50052);
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(AppModule, 
    {
    transport: Transport.GRPC,
    options: {
      package: 'proto',
      protoPath: [
        path.join(__dirname, '../proto/order.proto'),
        path.join(__dirname, '../proto/common.proto'),
      ],
      url: `0.0.0.0:${process.env.PORT??'50052'}`, // Specify your gRPC port
    },
  }
);

  await app.listen();
  console.log('Order gRPC Service is running on port 50052');
}
bootstrap();
