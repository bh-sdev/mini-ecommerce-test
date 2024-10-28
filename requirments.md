Senior Back-end Developer - Technical Test
Overview
This test is designed to assess your proficiency in designing and implementing a microservices architecture using gRPC, Nest.JS, and Golang as an API gateway. You are expected to develop scalable, efficient, and secure back-end systems, with database management handled using PostgreSQL and TypeORM.
Requirements
1. Product Service (Nest.JS):
Implement CRUD operations (Create, Read, Update, Delete) for products.
Each product should have at least the following fields:
name
description
price
available quantity
Design the database entities using TypeORM and utilize PostgreSQL for the database.
2. Order Service (Nest.JS):
Implement functionality to place an order. An order should reference one or more products and specify the quantities.
Implement an order listing functionality that returns a list of all orders, including product details.
Use gRPC to communicate with the Product Service to verify that the product exists and has the
required quantity available. Database design for the order entities should be done using TypeORM with PostgreSQL as the database.
3. API Gateway (Golang):
Develop an API Gateway that exposes REST endpoints for both the Product and Order Services.
Implement JWT authentication to secure these endpoints.
Bonus: Add a rate-limiting feature to protect the services from overuse.
4. gRPC:
Define the necessary gRPC protocols in a separate project.
Compile the protocols and include them in the respective services for inter-service communication.
5. gRPC Communication:
Establish gRPC communication between the microservices (Product and Order) and the API Gateway.
Define the appropriate Protobuf messages and services to handle the operations described above.
6. Bonus - Unit Tests (Optional):
Write additional unit tests for the Product and Order services to verify the core business logic.

Deployment and Documentation
1. Docker:
Containerize the Product Service, Order Service, and API Gateway.
Provide a Docker Compose file to orchestrate the startup of the entire system.
2. Documentation:
Document the system architecture and explain the role of each component.
Provide a Postman collection or Swagger documentation for the REST API exposed by the API Gateway.
Include clear setup and run instructions for the test environment.
Evaluation Criteria
Architecture Design: Efficiency, clarity, and scalability of the microservices architecture.
Code Quality: Adherence to best practices, readability, and proper error handling.
Functionality: Accurate implementation of the specified features and requirements.
Security: Implementation of authentication, authorization, and secure communication.
Database Design: Use of TypeORM and PostgreSQL in a well-structured, optimized manner.
Documentation and Deployment: Completeness and clarity of documentation, ease of deployment and setup.
Submission Guidelines
Provide the source code via a GitHub repository.
Include a README file with instructions on how to build and run the application, as well as any other relevant documentation.