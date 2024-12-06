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

3. **Run the Application**:

  ```bash
  go run main.go

4. **Test the API**:

Open a tool like Postman or use curl to test the API endpoints. By default, the server runs on http://localhost:8080.



