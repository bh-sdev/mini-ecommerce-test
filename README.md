
# Project Documentation

## Description

The project consists of three core components:

1. **Products Service** - Built with **NestJS (TypeScript)**, this service manages product information, including name, price, available quantity, and descriptions.
2. **Order Service** - Built with **NestJS (TypeScript)**, this service is responsible for placing and listing orders.
3. **API Gateway** - A **Golang-based** API gateway that provides a unified API for accessing both the Products and Order services.

Communication between the services and the API Gateway uses **gRPC** with data encoded in the **protobuf** format.

---

## Usage Instructions

The project provides three Docker files, one for each service, and a `docker-compose.yml` file that acts as a single entry point to build, boot, and orchestrate the services.

### Getting Started
To build and start all services, run:
```bash
docker compose up --build
```

- The **API Gateway** will be exposed on port `8080`.
- The **Product Service** is accessible internally on port `50051`.
- The **Order Service** is accessible internally on port `50052`.

### API Usage

Once all services are booted and running, you can interact with them using the following endpoints:

1. **Create an Account**
   ```http
   POST http://localhost:8080/register
   ```
   ```json
    {
        "username": "admin",
        "password": "password"
    }
   ```
   
2. **Login**
   ```http
   POST http://localhost:8080/login
   ```
   ```json
    {
        "username": "admin",
        "password": "password"
    }
   ```

3. **List Orders**
   ```http
   GET http://localhost:8080/orders
   ```

4. **Create an Order**
   ```http
   POST http://localhost:8080/orders
   ```
   ```json
    {
        "items":[
            {
                "productId": 1,
                "quantity": 3
            },
            {
                "productId": 24,
                "quantity": 2
            }
        ],
    }
   ```

5. **View an Order**
   ```http
   GET http://localhost:8080/orders/:id
   ```

### Admin Zone

Admin-specific actions require creating an admin user through the API Gateway command line.

1. Access the command line of the `ecommerce-api-gateway` container:
   ```bash
   docker exec -it <container-id> /bin/ash
   ```
   The above command should open the terminal for the container's filesystem
   By default, container id is likely to be `ecommerce-api-gateway-1`
   
2. Run the command to create a user with admin privileges. [Please selct admin as the user type here]:
   ```bash
   cmd user create
   ```
   OR 
    ```bash
   ./cmd user create
   ```

3. **Create a Product** (Admin Access Required)
    List Products
   ```http
   POST http://localhost:8080/products
   ```

---

## Further Documentation

More detailed documentation on advanced features and configurations is coming soon. Stay tuned!
