// import { NestFactory } from '@nestjs/core';
// import { AppModule } from './app.module';

// async function bootstrap() {
//   const app = await NestFactory.create(AppModule);
//   await app.listen(process.env.PORT ?? 3000);
// }
// bootstrap();

// src/main.ts
import { ValidationPipe } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import * as path from 'path';

async function bootstrap() {
  // const app = await NestFactory.create(AppModule);
  // app.useGlobalPipes(new ValidationPipe({ whitelist: true, forbidNonWhitelisted: true }));
  // await app.listen(3000);
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(AppModule, {
    transport: Transport.GRPC,
    options: {
      package: 'proto',
      protoPath: [
        path.join(__dirname, '../proto/product.proto'),
        path.join(__dirname, '../proto/common.proto'),
      ],
      url: `0.0.0.0:${process.env.PORT??'50051'}`, // Specify your gRPC port
    },
  });

  await app.listen();
  console.log('Product gRPC Service is running on port 50051');
}
bootstrap();
