# E-commerce API

An e-commerce API built with Go, using Gorilla Mux for routing, PostgreSQL for data persistence, and containerized with Docker.

---

## Features

- **CRUD operations** for products:
  - Create, Read, Update, and Delete product information.
- **Gorilla Mux** for efficient and flexible routing.
- **PostgreSQL** as the database for reliable and scalable data persistence.
- **Docker** support for seamless deployment and local development.

---

## Project Structure

```bash
.
├── main.go               # Entry point for the API
├── Dockerfile            # Docker configuration file
├── docker-compose.yml    # Docker Compose file for multi-container setup
├── README.md             # Project documentation
├── db/                   # Database-related files
│   ├── create_db.sql     # SQL script for database setup
│   └── migrate.sql       # SQL script for schema migrations
└── handlers/             # API handler logic
    ├── product.go        # Product-related business logic
    └── utils.go          # Utility functions
```

---

## Prerequisites

Ensure you have the following installed:

1. [Go](https://golang.org/dl/) (1.18 or later)
2. [Docker](https://www.docker.com/)
3. [PostgreSQL](https://www.postgresql.org/download/)

---

## Getting Started

### **Clone the Repository**

```bash
git clone https://github.com/yourusername/ecommerce-api.git
cd ecommerce-api
```
---

## Database Setup

1. Create a PostgreSQL database. Update the connection string in main.go to match your database credentials:

```bash
connStr := "user=<your_username> password=<your_password> dbname=ecommerce sslmode=disable"
```

2. Run the SQL script to set up the database:

```bash
psql -U <your_username> -d ecommerce -f db/create_db.sql
```

---

## Run Locally

1. Install dependencies:
```bash
go mod tidy
```

2.Start the API server:
```bash
go run main.go
```

3. Access the API at http://localhost:8080.

---

## Using Docker

1. Build and run the containers using Docker Compose:
```bash
docker-compose up --build
```
2. The API will be available at http://localhost:8080.

---

## API Endpoints

### Products

| Method | Endpoint          | Description               |
|--------|-------------------|---------------------------|
| GET    | `/products`       | Get all products          |
| GET    | `/products/{id}`  | Get a product by ID       |
| POST   | `/products`       | Create a new product      |
| PUT    | `/products/{id}`  | Update an existing product|
| DELETE | `/products/{id}`  | Delete a product by ID    |

---

## Example Product JSON
```bash
{
  "id": "1",
  "name": "Example Product",
  "description": "This is a sample product",
  "price": 19.99,
  "stock": 100
}
```
---

## Environment Variables

- POSTGRES_USER: PostgreSQL username
- POSTGRES_PASSWORD: PostgreSQL password
- POSTGRES_DB: PostgreSQL database name
- PORT: Port number for the API (default is 8080)

These can be configured in your `docker-compose.yml` file or a `.env` file for local development.

---

## Testing the API

Use a tool like Postman or cURL to interact with the API. For example:

```bash
curl -X GET http://localhost:8080/products
```
---

## Future Improvements

- Add user authentication (e.g., JWT-based).
- Implement unit and integration testing.
- Add Swagger documentation for the API.
- Deploy the API to a cloud provider like AWS or GCP.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## Contributing

Contributions are welcome! Feel free to submit a pull request or open an issue for any suggestions or bug reports.

---

## Author

**Your Name**  
GitHub: [rufusbiggs](https://github.com/rufusbiggs)


Feel free to update any placeholders like `rufusbiggs` or add your personal touches! Let me know if you need adjustments.
