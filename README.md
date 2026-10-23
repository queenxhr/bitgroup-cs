# Bitgroup-CS: Product & Category CRUD Service

Golang REST API for Product & Category CRUD with Clean Architecture following the `logistic-service` design pattern:
```
entity -> repository -> model -> MySQL/PostgreSQL -> usecase -> http handler -> main -> router
```

## 🛠 Tech Stack & Architecture
- **Language**: Go 1.26+
- **Framework**: [Gin Gonic](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io)
- **Database**: PostgreSQL (`db_bitgroup_cs`)
- **Architecture**: Domain-Driven / Clean Layered Architecture:
  - `internal/entity`: Domain structs (`Category`, `Product`)
  - `internal/repository`: Repository interfaces
  - `infrastructure/repository`: Database models & GORM queries (`model.go`, `mysql.go`, `postgres.go`)
  - `internal/usecase`: Business logic layer (`CategoryUseCase`, `ProductUseCase`)
  - `handler/v1`: HTTP transport layer (`dto_request.go`, `dto_response.go`, `handler.go`, `router.go`)
  - `infrastructure/database`: PostgreSQL connection & auto-migration
  - `config`: Environment configuration loader
  - `utils`: HTTP standard response and error handlers

---

## 🗄 Database Schema (PostgreSQL: `db_bitgroup_cs`)

### 1. Categories
```sql
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### 2. Products
```sql
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    price NUMERIC(12, 2) NOT NULL DEFAULT 0.00,
    category_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_products_categories 
        FOREIGN KEY (category_id) 
        REFERENCES categories(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE
);
```

---

## 🚀 Getting Started

### 1. Configure Environment
Copy `.env.example` to `.env` and adjust the PostgreSQL credentials if needed:
```env
PORT=8080
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=db_bitgroup_cs
DB_SSLMODE=disable
DB_TIMEZONE=Asia/Jakarta
```

### 2. Create PostgreSQL Database
Run in your PostgreSQL terminal:
```sql
CREATE DATABASE db_bitgroup_cs;
```
*(The tables and foreign keys will be auto-migrated on application startup, or you can run `db_bitgroup_cs.sql` manually).*

### 3. Run the Service
```bash
go run main.go
```

---

## 📡 API Endpoints

### Categories (`/api/v1/categories`)
| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/v1/categories` | Get all categories |
| `GET` | `/api/v1/categories/:id` | Get category by ID |
| `POST` | `/api/v1/categories` | Create new category |
| `PUT` | `/api/v1/categories/:id` | Update category by ID |
| `DELETE` | `/api/v1/categories/:id` | Delete category by ID |

#### Category Payload Examples:
**Create/Update Category:**
```json
{
  "name": "Electronics"
}
```

---

### Products (`/api/v1/products`)
| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/api/v1/products` | Get all products (with preloaded category) |
| `GET` | `/api/v1/products/:id` | Get product by ID (with preloaded category) |
| `POST` | `/api/v1/products` | Create new product |
| `PUT` | `/api/v1/products/:id` | Update product by ID |
| `DELETE` | `/api/v1/products/:id` | Delete product by ID |

#### Product Payload Examples:
**Create Product:**
```json
{
  "name": "Logitech MX Master 3S",
  "price": 1650000.00,
  "category_id": 1
}
```

**Update Product:**
```json
{
  "name": "Logitech MX Master 3S Black",
  "price": 1699000.00,
  "category_id": 1
}
```
