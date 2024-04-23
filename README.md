# Product Order Microservices Project

=====================================

## Overview

This project demonstrates a real-world implementation of a microservices architecture using Golang, Fiber, SQLC, gRPC, PostgreSQL, Docker, and Docker Compose. The project consists of three main services: Order Service, Product Service, and API Gateway (Product_Order).

### Goal

The primary goal of this project is to showcase how to design and implement a scalable and  maintainable microservices architecture using the aforementioned technologies.

## Services

### 1. Order Service

The Order Service is responsible for managing orders. It exposes a gRPC API that allows clients to create, read, update, and delete orders.

### 2. Product Service

The Product Service is responsible for managing products. It exposes a gRPC API that allows clients to create, read, update, and delete products.

### 3. API Gateway (Product_Order)

The API Gateway acts as an entry point for clients. It exposes a RESTful API that allows clients to interact with the Order Service and Product Service.

## Technologies Used

* **Golang**: The programming language used for developing the microservices.
* **Fiber**: A fast and flexible web framework for building web applications.
* **SQLC**: A SQL client library for interacting with PostgreSQL databases.
* **Mockery**: A tool for generating mock implementations of interfaces.
* **Migration**: A tool for managing database schema migrations.
* **gRPC**: A high-performance RPC framework for building scalable APIs.
* **PostgreSQL**: A relational database management system used for storing data.
* **Docker**: A containerization platform for deploying and managing applications.
* **Docker Compose**: A tool for defining and running multi-container Docker applications.
* **hexagonal architecture**: A software architecture pattern that separates an application's business logic (the hexagon) from its infrastructure and presentation layers. It uses **ports** (interfaces) to define interactions and **adapters** to connect the hexagon to specific infrastructure or presentation layers.

## Running the Application

To run this application, you need to have `Makefile` and `Docker` installed on your laptop. Follow these steps to get started:

### Step 1: Navigate to the Project Directory

Open a terminal and navigate to the project directory.

### Step 2: Run the Application

Run in the terminal `docker network create store` then
Run the command `make run` to start the application. This will spin up the necessary containers and start the services.
That's it! You should now have the application up and running.

### Accessing the Application

Once the application is running, you can access it at:

* `http://localhost:4002/order/{order_id}`
You can create orders and products by using the gRPC connections:
* Order Service: `grpc://localhost:4000`
* Product Service: `grpc://localhost:4001`
Refer to the `order.proto` and `product.proto` files for more information on how to interact with the services.

## Testing

Testing is an essential part of this project. We employ a comprehensive testing strategy to ensure the reliability and integrity of our microservices.

### Unit Tests

We write unit tests to verify that each service functions correctly in isolation. These tests focus on individual components, ensuring they behave as expected.

### Running Tests

To run the tests, follow these steps:

1. Open any microservice directory.
2. Run the command `make test`.
This will spin up a Postgres container specifically for testing purposes, allowing you to execute the tests in a controlled environment.

## GitHub Action Setup

This repository's workflow employs two main components:

### 1. GitHub Action for Each Microservice

Each microservice has its dedicated GitHub Action workflow. This workflow ensures that the code pushed to the repository meets several quality checks before integration.

#### Quality Checks

* **Dependency Verification:** The action verifies the dependencies listed in `go.sum` and ensures that the required dependencies are correctly vendored.
* **Code Formatting:** It ensures code consistency and readability by running a code vetting process.
* **Static Code Analysis:** Utilizes staticcheck to perform static code analysis, enhancing code quality and identifying potential issues early in the development cycle.

These checks collectively ensure that each microservice maintains high standards of code quality and reliability.

### 2. GitHub Action for the Main Project

The main project's GitHub Action workflow orchestrates the integration of all microservices and ensures the smooth functioning of the entire application.

#### Integration Process

* **Docker Compose Execution:** The workflow orchestrates the execution of the application via Docker Compose, ensuring that all microservices interact seamlessly within the defined environment.
* **Testing:** It executes both unit tests and end-to-end tests to validate the functionality and integrity of the codebase.
* **Inter-Microservice Communication:** Verifies that each microservice communicates effectively with others, maintaining the expected behavior of the entire system.

By employing these actions, the repository ensures that any code pushed to it undergoes rigorous testing and integration processes before being integrated into the main project.
