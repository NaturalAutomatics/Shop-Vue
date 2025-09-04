-- Export existing data from local PostgreSQL
-- Run these commands to export your current data

-- Export users
\copy (SELECT username, email, name, role, password_hash FROM users) TO 'users_export.csv' WITH CSV HEADER;

-- Export products
\copy (SELECT name, description, price, category, image, stock FROM products) TO 'products_export.csv' WITH CSV HEADER;

-- Export orders
\copy (SELECT id, order_number, customer_name, customer_email, customer_address, subtotal, shipping, tax, total, status, created_at, updated_at FROM orders) TO 'orders_export.csv' WITH CSV HEADER;

-- Export order items
\copy (SELECT order_id, product_id, name, price, quantity FROM order_items) TO 'order_items_export.csv' WITH CSV HEADER;