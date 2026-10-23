-- ==========================================================
-- Database: db_bitgroup_cs
-- PostgreSQL Schema for Categories and Products
-- ==========================================================

-- 1. Create Database if not exists
-- CREATE DATABASE db_bitgroup_cs;

-- 2. Connect to db_bitgroup_cs
-- \c db_bitgroup_cs;

-- 3. Create categories table
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- 4. Create products table
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

-- ==========================================================
-- Sample Dummy Data (Optional)
-- ==========================================================
INSERT INTO categories (name) VALUES 
('Electronics'),
('Clothing'),
('Books')
ON CONFLICT DO NOTHING;

INSERT INTO products (name, price, category_id) VALUES 
('Mechanical Keyboard', 750000.00, 1),
('Cotton T-Shirt', 125000.00, 2),
('Golang Clean Architecture Book', 250000.00, 3)
ON CONFLICT DO NOTHING;
