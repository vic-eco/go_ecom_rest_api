# Simple Go E-Commerce API

A minimal e-commerce API built with **Go**, designed to manage products and orders. This project uses **[Chi](https://github.com/go-chi/chi)** for routing, **[sqlc](https://sqlc.dev/)** for type-safe SQL queries, and **[Goose](https://github.com/pressly/goose)** for database migrations.

---

## Features

### Products
- `GET /products` – List all products
- `GET /products/{id}` – Get a specific product by ID
- `POST /products` – Create a new product

### Orders
- `GET /orders/{id}` – Get a specific order by ID
- `POST /orders` – Place a new order

---

## Tech Stack

- **Go** 
- **Chi**
- **sqlc**
- **Goose**
- **PostgreSQL**

---
## Next Steps

The next step for this API is to implement authentication and authorization, ensuring that:
 - Users must log in to place orders or manage products.
 - Only authorized users (e.g., admins) can create or update products.

And also implement more endpoints.