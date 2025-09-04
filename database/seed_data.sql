-- Seed Data for Vue Shop Database
-- Run this after migration.sql

-- Insert users with bcrypt hashed passwords
INSERT INTO users (username, email, name, role, password_hash) VALUES
('admin', 'admin@vueshop.com', 'Admin User', 'admin', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi'),
('john', 'john@example.com', 'John Doe', 'customer', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi'),
('jane', 'jane@example.com', 'Jane Smith', 'customer', '$2a$10$92IXUNpkjO0rOQ5byMi.Ye4oKoEa3Ro9llC/.og/at2.uheWG/igi')
ON CONFLICT (username) DO NOTHING;

-- Insert products
INSERT INTO products (name, description, price, category, image, stock) VALUES
('Vue.js T-Shirt', 'Comfortable cotton t-shirt with Vue.js logo', 25.99, 'clothing', 'https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400&h=300&fit=crop&bg=white', 50),
('JavaScript Book', 'Comprehensive guide to modern JavaScript', 39.99, 'books', 'https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=400&h=300&fit=crop&bg=white', 25),
('Wireless Headphones', 'High-quality wireless headphones with noise cancellation', 89.99, 'electronics', 'https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=400&h=300&fit=crop&bg=white', 15),
('Vue Hoodie', 'Warm and cozy hoodie perfect for Vue developers', 45.99, 'clothing', 'https://images.unsplash.com/photo-1556821840-3a63f95609a7?w=400&h=300&fit=crop&bg=white', 30),
('Laptop Stand', 'Ergonomic laptop stand for better posture', 29.99, 'electronics', 'https://images.unsplash.com/photo-1586953208448-b95a79798f07?w=400&h=300&fit=crop&bg=white', 20),
('Vue.js Guide', 'Complete Vue.js 3 tutorial and reference', 19.99, 'books', 'https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=400&h=300&fit=crop&bg=white', 40),
('Mechanical Keyboard', 'Premium mechanical keyboard with RGB backlighting', 129.99, 'electronics', 'https://images.unsplash.com/photo-1541140532154-b024d705b90a?w=400&h=300&fit=crop&bg=white', 10),
('Developer Mug', 'Ceramic coffee mug perfect for coding sessions', 12.99, 'clothing', 'https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400&h=300&fit=crop&bg=white', 100),
('React vs Vue Book', 'In-depth comparison of modern JavaScript frameworks', 34.99, 'books', 'https://images.unsplash.com/photo-1589829085413-56de8ae18c73?w=400&h=300&fit=crop&bg=white', 35)
ON CONFLICT DO NOTHING;