-- Create the categories table if it doesn't exist
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);

-- Insert initial data with idempotency
INSERT INTO categories (id, title, created_at, updated_at) VALUES
(1, 'Books', NOW(), NOW()),
(2, 'Health', NOW(), NOW()),
(3, 'Electronics', NOW(), NOW()),
(4, 'Clothing', NOW(), NOW()),
(5, 'Home & Garden', NOW(), NOW()),
(6, 'Sports', NOW(), NOW()),
(7, 'Toys & Games', NOW(), NOW()),
(8, 'Automotive', NOW(), NOW()),
(9, 'Beauty & Personal Care', NOW(), NOW()),
(10, 'Travel', NOW(), NOW())
ON CONFLICT (id) DO NOTHING;