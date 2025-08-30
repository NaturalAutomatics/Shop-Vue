package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// AdminStats represents database statistics
type AdminStats struct {
	Products   int     `json:"products"`
	Users      int     `json:"users"`
	Orders     int     `json:"orders"`
	TotalValue float64 `json:"totalValue"`
}

// DatabaseConnection represents connection test request
type DatabaseConnection struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Database string `json:"database"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// testConnection handles POST /api/admin/test-connection
func testConnection(c *gin.Context) {
	var conn DatabaseConnection
	if err := c.ShouldBindJSON(&conn); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid request",
			Message: err.Error(),
		})
		return
	}

	// Test connection with provided credentials
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conn.Host, conn.Port, conn.Username, conn.Password, conn.Database)

	testDB, err := sql.Open("postgres", connStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Connection failed",
			Message: err.Error(),
		})
		return
	}
	defer testDB.Close()

	if err := testDB.Ping(); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Connection test failed",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Database connection successful",
	})
}

// getStats handles GET /api/admin/stats
func getStats(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{
			Success: false,
			Error:   "Database not connected",
		})
		return
	}

	stats := AdminStats{}

	// Get products count
	if err := db.QueryRow(`SELECT COUNT(*) FROM products`).Scan(&stats.Products); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to get products count",
			Message: err.Error(),
		})
		return
	}

	// Get users count
	if err := db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&stats.Users); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to get users count",
			Message: err.Error(),
		})
		return
	}

	// Get orders count and total value
	if err := db.QueryRow(`SELECT COUNT(*), COALESCE(SUM(total), 0) FROM orders`).Scan(&stats.Orders, &stats.TotalValue); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to get orders stats",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    stats,
	})
}

// getAllUsers handles GET /api/admin/users
func getAllUsers(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{
			Success: false,
			Error:   "Database not connected",
		})
		return
	}

	rows, err := db.Query(`SELECT id, username, email, name, role FROM users ORDER BY id`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to fetch users",
			Message: err.Error(),
		})
		return
	}
	defer rows.Close()

	users := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Name, &user.Role); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Success: false,
				Error:   "Failed to scan user data",
				Message: err.Error(),
			})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    users,
	})
}

// seedDatabase handles POST /api/admin/seed
func seedDatabase(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{
			Success: false,
			Error:   "Database not connected",
		})
		return
	}

	// Seed products
	for _, p := range mockProducts {
		_, err := db.Exec(`INSERT INTO products (name, description, price, category, image, stock) VALUES ($1,$2,$3,$4,$5,$6) ON CONFLICT DO NOTHING`,
			p.Name, p.Description, p.Price, p.Category, p.Image, p.Stock,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Success: false,
				Error:   "Failed to seed products",
				Message: err.Error(),
			})
			return
		}
	}

	// Seed users
	seedUsers := []struct {
		Username string
		Email    string
		Name     string
		Role     string
		Password string
	}{
		{"admin", "admin@vueshop.com", "Admin User", "admin", "admin123"},
		{"john", "john@example.com", "John Doe", "customer", "password123"},
		{"jane", "jane@example.com", "Jane Smith", "customer", "password456"},
	}

	for _, u := range seedUsers {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Success: false,
				Error:   "Failed to hash password",
				Message: err.Error(),
			})
			return
		}

		_, err = db.Exec(`INSERT INTO users (username, email, name, role, password_hash) VALUES ($1,$2,$3,$4,$5) ON CONFLICT DO NOTHING`,
			u.Username, u.Email, u.Name, u.Role, string(hash),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Success: false,
				Error:   "Failed to seed users",
				Message: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Database seeded successfully",
	})
}

// clearDatabase handles POST /api/admin/clear
func clearDatabase(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{
			Success: false,
			Error:   "Database not connected",
		})
		return
	}

	// Clear all data (in reverse dependency order)
	tables := []string{"order_items", "orders", "products", "users"}
	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf(`DELETE FROM %s`, table))
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Success: false,
				Error:   fmt.Sprintf("Failed to clear %s", table),
				Message: err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Database cleared successfully",
	})
}

// exportData handles GET /api/admin/export
func exportData(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{
			Success: false,
			Error:   "Database not connected",
		})
		return
	}

	export := make(map[string]interface{})

	// Export products
	productRows, err := db.Query(`SELECT id, name, description, price, category, image, stock FROM products`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to export products",
			Message: err.Error(),
		})
		return
	}
	defer productRows.Close()

	products := make([]Product, 0)
	for productRows.Next() {
		var p Product
		if err := productRows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Image, &p.Stock); err != nil {
			continue
		}
		products = append(products, p)
	}
	export["products"] = products

	// Export users (without passwords)
	userRows, err := db.Query(`SELECT id, username, email, name, role FROM users`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to export users",
			Message: err.Error(),
		})
		return
	}
	defer userRows.Close()

	users := make([]User, 0)
	for userRows.Next() {
		var u User
		if err := userRows.Scan(&u.ID, &u.Username, &u.Email, &u.Name, &u.Role); err != nil {
			continue
		}
		users = append(users, u)
	}
	export["users"] = users

	// Export orders
	orderRows, err := db.Query(`SELECT id, order_number, customer_name, customer_email, customer_address, subtotal, shipping, tax, total, status, created_at, updated_at FROM orders`)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to export orders",
			Message: err.Error(),
		})
		return
	}
	defer orderRows.Close()

	orders := make([]Order, 0)
	for orderRows.Next() {
		var o Order
		if err := orderRows.Scan(&o.ID, &o.OrderNumber, &o.Customer.Name, &o.Customer.Email, &o.Customer.Address, &o.Subtotal, &o.Shipping, &o.Tax, &o.Total, &o.Status, &o.CreatedAt, &o.UpdatedAt); err != nil {
			continue
		}
		orders = append(orders, o)
	}
	export["orders"] = orders

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    export,
	})
}

// deleteProduct handles DELETE /api/admin/products/:id
func deleteProduct(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{
			Success: false,
			Error:   "Database not connected",
		})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid product ID",
		})
		return
	}

	result, err := db.Exec(`DELETE FROM products WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to delete product",
			Message: err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Success: false,
			Error:   "Product not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Product deleted successfully",
	})
}

// deleteUser handles DELETE /api/admin/users/:id
func deleteUser(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{
			Success: false,
			Error:   "Database not connected",
		})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid user ID",
		})
		return
	}

	result, err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to delete user",
			Message: err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Success: false,
			Error:   "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User deleted successfully",
	})
}

// createProduct handles POST /api/admin/products
func createProduct(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{
			Success: false,
			Error:   "Database not connected",
		})
		return
	}

	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid product data",
			Message: err.Error(),
		})
		return
	}

	var id int
	err := db.QueryRow(`INSERT INTO products (name, description, price, category, image, stock) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`,
		product.Name, product.Description, product.Price, product.Category, product.Image, product.Stock).Scan(&id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to create product",
			Message: err.Error(),
		})
		return
	}

	product.ID = id
	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    product,
		"message": "Product created successfully",
	})
}

// updateProduct handles PUT /api/admin/products/:id
func updateProduct(c *gin.Context) {
	if db == nil {
		c.JSON(http.StatusServiceUnavailable, ErrorResponse{
			Success: false,
			Error:   "Database not connected",
		})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid product ID",
		})
		return
	}

	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid product data",
			Message: err.Error(),
		})
		return
	}

	result, err := db.Exec(`UPDATE products SET name=$1, description=$2, price=$3, category=$4, image=$5, stock=$6 WHERE id=$7`,
		product.Name, product.Description, product.Price, product.Category, product.Image, product.Stock, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Success: false,
			Error:   "Failed to update product",
			Message: err.Error(),
		})
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, ErrorResponse{
			Success: false,
			Error:   "Product not found",
		})
		return
	}

	product.ID = id
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
		"message": "Product updated successfully",
	})
}
