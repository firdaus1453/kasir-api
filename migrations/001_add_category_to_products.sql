-- Create categories table if not exists (already implied by repo, but good to ensure)
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- Add category_id to products
ALTER TABLE products ADD COLUMN IF NOT EXISTS category_id INTEGER;

-- Add foreign key constraint
ALTER TABLE products ADD CONSTRAINT fk_category FOREIGN KEY (category_id) REFERENCES categories(id);

-- Insert some dummy categories
INSERT INTO categories (name) VALUES ('Makanan'), ('Minuman'), ('Elektronik');

-- Update existing products to have a default category (e.g., 1)
UPDATE products SET category_id = 1 WHERE category_id IS NULL;
