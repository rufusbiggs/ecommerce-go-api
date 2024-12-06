# Ecommerce API

This project is an **Ecommerce API** built with **Golang** and the **Gorilla Mux** router. The API is designed to handle typical e-commerce functionalities, such as managing products, customers, and orders. The project is containerized using **Docker** for seamless deployment and scalability.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Prerequisites](#prerequisites)
- [Setup and Running Locally](#setup-and-running-locally)
- [Dockerizing the API](#dockerizing-the-api)
- [Endpoints](#endpoints)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## Features

- RESTful API for managing products, customers, and orders.
- Built with **Gorilla Mux** for routing.
- Containerized using Docker for easy deployment.
- Lightweight and fast using Go’s concurrency model.

## Technologies Used

- **Golang**: Backend programming language.
- **Gorilla Mux**: Router for handling API routes.
- **Docker**: For containerizing the application.

## Prerequisites

Ensure the following tools are installed on your system:

- **Go** (version 1.20 or later)
- **Docker**
- A text editor or IDE (e.g., VS Code)

## Setup and Running Locally

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-username/ecommerce-api.git
   cd ecommerce-api

2. **Install Dependencies**:

The `go.mod` file contains the project dependencies. To download them:

  ```bash
  go mod download
```

3. **Run the Application**:

  ```bash
  go run main.go
```

4. **Test the API**:

Open a tool like Postman or use `curl` to test the API endpoints. By default, the server runs on `http://localhost:8080`.

## Dockerizing the API

1. **Write a Dockerfile**:

The `Dockerfile` is already included in the project. Here's what it does:

```bash
# Stage 1: Build the Go application
FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Stage 2: Create a minimal runtime container
FROM debian:buster-slim
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/main .
RUN chmod +x /app/main
EXPOSE 8080
ENTRYPOINT ["/app/main"]
```

2. **Build the Docker Image**:

Run the following command to build the Docker image:

```bash
docker build -t ecommerce-api .
```

3. **Run the Container**:

To run the containerized application:

```bash
docker run -p 8080:8080 ecommerce-api
```

4. **Access the API**:

The API is now accessible at `http://localhost:8080` or the host's IP address.

## Endpoints

Here are some example API endpoints:

- **GET /products**: Retrieve a list of products.
- **POST /products**: Add a new product.
- **GET /products/{id}**: Get a single product by ID.
- **PUT /products/{id}**: Update a product by ID.
- **DELETE /products/{id}**: Delete a product by ID.

You can add more endpoints for managing customers and orders as needed.

## Project Structure

```bash
ecommerce-api/
├── Dockerfile        # Instructions to build the Docker image
├── main.go           # Entry point of the application
├── go.mod            # Go module file
├── go.sum            # Dependencies checksum
├── README.md         # Project documentation
```

## Contributing

Contributions are welcome! Please follow these steps:

 1. Fork the repository.
 2. Create a new branch (`git checkout -b feature-branch`).
 3. Make your changes.
 4. Push to your branch (`git push origin feature-branch`).
 5. Create a pull request.

## License

This project is licensed under the MIT LIcense. See the LICENSE file for details.

```bash

This version keeps the formatting consistent and ready for a GitHub **README.md** file.
```

   



