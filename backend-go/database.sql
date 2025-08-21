-- Vue Shop Database Schema for MySQL
-- Run this SQL to create the database and tables

-- Create database
CREATE DATABASE IF NOT EXISTS vue_shop_db
CHARACTER SET utf8mb4
COLLATE utf8mb4_unicode_ci;

USE vue_shop_db;

-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(100) NOT NULL,
    image_url VARCHAR(500),
    stock INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_category (category),
    INDEX idx_price (price),
    INDEX idx_stock (stock)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create customers table
CREATE TABLE IF NOT EXISTS customers (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    address TEXT NOT NULL,
    phone VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create orders table
CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(36) PRIMARY KEY,
    order_number VARCHAR(20) NOT NULL UNIQUE,
    customer_id INT NOT NULL,
    subtotal DECIMAL(10, 2) NOT NULL,
    shipping DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    tax DECIMAL(10, 2) NOT NULL DEFAULT 0.00,
    total DECIMAL(10, 2) NOT NULL,
    status ENUM('pending', 'processing', 'shipped', 'delivered', 'cancelled') DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    INDEX idx_order_number (order_number),
    INDEX idx_status (status),
    INDEX idx_created_at (created_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create order_items table
CREATE TABLE IF NOT EXISTS order_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_id VARCHAR(36) NOT NULL,
    product_id INT NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    product_price DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE RESTRICT,
    INDEX idx_order_id (order_id),
    INDEX idx_product_id (product_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Insert sample products data
INSERT INTO products (name, description, price, category, image_url, stock) VALUES
('Vue.js T-Shirt', 'Comfortable cotton t-shirt with Vue.js logo', 25.99, 'clothing', 'https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400&h=300&fit=crop&bg=white', 50),
('JavaScript Book', 'Comprehensive guide to modern JavaScript', 39.99, 'books', 'https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=400&h=300&fit=crop&bg=white', 25),
('Wireless Headphones', 'High-quality wireless headphones with noise cancellation', 89.99, 'electronics', 'https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=400&h=300&fit=crop&bg=white', 15),
('Vue Hoodie', 'Warm and cozy hoodie perfect for Vue developers', 45.99, 'clothing', 'https://images.unsplash.com/photo-1556821840-3a63f95609a7?w=400&h=300&fit=crop&bg=white', 30),
('Laptop Stand', 'Ergonomic laptop stand for better posture', 29.99, 'electronics', 'https://images.unsplash.com/photo-1586953208448-b95a79798f07?w=400&h=300&fit=crop&bg=white', 20),
('Vue.js Guide', 'Complete Vue.js 3 tutorial and reference', 19.99, 'books', 'https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=400&h=300&fit=crop&bg=white', 40),
('Mechanical Keyboard', 'Premium mechanical keyboard with RGB backlighting', 129.99, 'electronics', 'https://images.unsplash.com/photo-1541140532154-b024d705b90a?w=400&h=300&fit=crop&bg=white', 10),
('Developer Mug', 'Ceramic coffee mug perfect for coding sessions', 12.99, 'clothing', 'https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400&h=300&fit=crop&bg=white', 100),
('React vs Vue Book', 'In-depth comparison of modern JavaScript frameworks', 34.99, 'books', 'https://images.unsplash.com/photo-1589829085413-56de8ae18c73?w=400&h=300&fit=crop&bg=white', 35);

-- Create indexes for better performance
CREATE INDEX idx_products_category_price ON products(category, price);
CREATE INDEX idx_orders_customer_status ON orders(customer_id, status);
CREATE INDEX idx_order_items_order_product ON order_items(order_id, product_id);

-- Create a view for order summary
CREATE VIEW order_summary AS
SELECT 
    o.id,
    o.order_number,
    c.name as customer_name,
    c.email as customer_email,
    o.subtotal,
    o.shipping,
    o.tax,
    o.total,
    o.status,
    o.created_at,
    COUNT(oi.id) as item_count
FROM orders o
JOIN customers c ON o.customer_id = c.id
LEFT JOIN order_items oi ON o.id = oi.order_id
GROUP BY o.id, o.order_number, c.name, c.email, o.subtotal, o.shipping, o.tax, o.total, o.status, o.created_at;

-- Create stored procedure for creating orders
DELIMITER //
CREATE PROCEDURE CreateOrder(
    IN p_customer_name VARCHAR(255),
    IN p_customer_email VARCHAR(255),
    IN p_customer_address TEXT,
    IN p_customer_phone VARCHAR(20),
    IN p_subtotal DECIMAL(10, 2),
    IN p_shipping DECIMAL(10, 2),
    IN p_tax DECIMAL(10, 2),
    IN p_total DECIMAL(10, 2),
    OUT p_order_id VARCHAR(36),
    OUT p_order_number VARCHAR(20)
)
BEGIN
    DECLARE v_customer_id INT;
    DECLARE v_order_counter INT;
    
    -- Insert or get customer
    INSERT INTO customers (name, email, address, phone) 
    VALUES (p_customer_name, p_customer_email, p_customer_address, p_customer_phone)
    ON DUPLICATE KEY UPDATE 
        name = VALUES(name),
        address = VALUES(address),
        phone = VALUES(phone),
        updated_at = CURRENT_TIMESTAMP;
    
    SET v_customer_id = LAST_INSERT_ID();
    IF v_customer_id = 0 THEN
        SELECT id INTO v_customer_id FROM customers WHERE email = p_customer_email;
    END IF;
    
    -- Generate order number
    SELECT COALESCE(MAX(CAST(SUBSTRING(order_number, 5) AS UNSIGNED)), 1000) + 1 
    INTO v_order_counter 
    FROM orders;
    
    SET p_order_number = CONCAT('VUE-', v_order_counter);
    SET p_order_id = UUID();
    
    -- Create order
    INSERT INTO orders (id, order_number, customer_id, subtotal, shipping, tax, total)
    VALUES (p_order_id, p_order_number, v_customer_id, p_subtotal, p_shipping, p_tax, p_total);
END //
DELIMITER ;

-- Create stored procedure for adding order items
DELIMITER //
CREATE PROCEDURE AddOrderItem(
    IN p_order_id VARCHAR(36),
    IN p_product_id INT,
    IN p_quantity INT
)
BEGIN
    DECLARE v_product_name VARCHAR(255);
    DECLARE v_product_price DECIMAL(10, 2);
    DECLARE v_total_price DECIMAL(10, 2);
    
    -- Get product details
    SELECT name, price INTO v_product_name, v_product_price
    FROM products WHERE id = p_product_id;
    
    SET v_total_price = v_product_price * p_quantity;
    
    -- Insert order item
    INSERT INTO order_items (order_id, product_id, product_name, product_price, quantity, total_price)
    VALUES (p_order_id, p_product_id, v_product_name, v_product_price, p_quantity, v_total_price);
    
    -- Update product stock
    UPDATE products SET stock = stock - p_quantity WHERE id = p_product_id;
END //
DELIMITER ;

-- Create trigger to update product stock when order is cancelled
DELIMITER //
CREATE TRIGGER after_order_cancelled
AFTER UPDATE ON orders
FOR EACH ROW
BEGIN
    IF NEW.status = 'cancelled' AND OLD.status != 'cancelled' THEN
        UPDATE products p
        JOIN order_items oi ON p.id = oi.product_id
        SET p.stock = p.stock + oi.quantity
        WHERE oi.order_id = NEW.id;
    END IF;
END //
DELIMITER ;

-- Show created tables
SHOW TABLES;

-- Show sample data
SELECT 'Products' as table_name, COUNT(*) as record_count FROM products
UNION ALL
SELECT 'Customers' as table_name, COUNT(*) as record_count FROM customers
UNION ALL
SELECT 'Orders' as table_name, COUNT(*) as record_count FROM orders
UNION ALL
SELECT 'Order Items' as table_name, COUNT(*) as record_count FROM order_items;
