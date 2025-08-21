package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// getProducts handles GET /api/products
func getProducts(c *gin.Context) {
	// Prefer Postgres when available
	if db != nil {
		category := c.Query("category")
		search := c.Query("search")
		sortBy := c.Query("sort")

		query := `SELECT id, name, description, price, category, image, stock FROM products`
		var filters []string
		var args []interface{}
		arg := 1
		if category != "" && category != "all" {
			filters = append(filters, fmt.Sprintf("category = $%d", arg))
			args = append(args, category)
			arg++
		}
		if search != "" {
			filters = append(filters, fmt.Sprintf("(LOWER(name) LIKE $%d OR LOWER(description) LIKE $%d)", arg, arg+1))
			like := "%" + strings.ToLower(search) + "%"
			args = append(args, like, like)
			arg += 2
		}
		if len(filters) > 0 {
			query += " WHERE " + strings.Join(filters, " AND ")
		}
		switch sortBy {
		case "price-asc":
			query += " ORDER BY price ASC"
		case "price-desc":
			query += " ORDER BY price DESC"
		case "name-asc":
			query += " ORDER BY name ASC"
		case "name-desc":
			query += " ORDER BY name DESC"
		}

		rows, err := db.Query(query, args...)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
			return
		}
		defer rows.Close()

		products := make([]Product, 0)
		for rows.Next() {
			var p Product
			if err := rows.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Image, &p.Stock); err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
				return
			}
			products = append(products, p)
		}
		if err := rows.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
			return
		}

		c.JSON(http.StatusOK, ProductsResponse{Success: true, Data: products, Total: len(products)})
		return
	}

	// Fallback to mock data
	category := c.Query("category")
	search := c.Query("search")
	sortBy := c.Query("sort")

	filteredProducts := make([]Product, 0)
	for _, product := range mockProducts {
		// Filter by category
		if category != "" && category != "all" && product.Category != category {
			continue
		}

		// Filter by search term
		if search != "" {
			searchLower := strings.ToLower(search)
			nameLower := strings.ToLower(product.Name)
			descLower := strings.ToLower(product.Description)
			if !strings.Contains(nameLower, searchLower) && !strings.Contains(descLower, searchLower) {
				continue
			}
		}

		filteredProducts = append(filteredProducts, product)
	}

	// Sort products
	switch sortBy {
	case "price-asc":
		sort.Slice(filteredProducts, func(i, j int) bool {
			return filteredProducts[i].Price < filteredProducts[j].Price
		})
	case "price-desc":
		sort.Slice(filteredProducts, func(i, j int) bool {
			return filteredProducts[i].Price > filteredProducts[j].Price
		})
	case "name-asc":
		sort.Slice(filteredProducts, func(i, j int) bool {
			return filteredProducts[i].Name < filteredProducts[j].Name
		})
	case "name-desc":
		sort.Slice(filteredProducts, func(i, j int) bool {
			return filteredProducts[i].Name > filteredProducts[j].Name
		})
	}

	response := ProductsResponse{
		Success: true,
		Data:    filteredProducts,
		Total:   len(filteredProducts),
	}

	c.JSON(http.StatusOK, response)
}

// getProduct handles GET /api/products/:id
func getProduct(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid product ID",
		})
		return
	}

	if db != nil {
		var p Product
		row := db.QueryRow(`SELECT id, name, description, price, category, image, stock FROM products WHERE id = $1`, id)
		if err := row.Scan(&p.ID, &p.Name, &p.Description, &p.Price, &p.Category, &p.Image, &p.Stock); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusNotFound, ErrorResponse{Success: false, Error: "Product not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
			return
		}
		c.JSON(http.StatusOK, ProductResponse{Success: true, Data: p})
		return
	}

	// Fallback mock
	for _, product := range mockProducts {
		if product.ID == id {
			response := ProductResponse{
				Success: true,
				Data:    product,
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{
		Success: false,
		Error:   "Product not found",
	})
}

// getCategories handles GET /api/products/categories
func getCategories(c *gin.Context) {
	if db != nil {
		rows, err := db.Query(`SELECT DISTINCT category FROM products ORDER BY category ASC`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
			return
		}
		defer rows.Close()
		cats := make([]string, 0)
		for rows.Next() {
			var cat string
			if err := rows.Scan(&cat); err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
				return
			}
			cats = append(cats, cat)
		}
		c.JSON(http.StatusOK, CategoriesResponse{Success: true, Data: cats})
		return
	}

	// Fallback mock
	categoryMap := make(map[string]bool)
	for _, product := range mockProducts {
		categoryMap[product.Category] = true
	}

	categories := make([]string, 0, len(categoryMap))
	for category := range categoryMap {
		categories = append(categories, category)
	}

	sort.Strings(categories)

	response := CategoriesResponse{
		Success: true,
		Data:    categories,
	}

	c.JSON(http.StatusOK, response)
}

// createOrder handles POST /api/orders
func createOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Validation failed",
			Message: err.Error(),
		})
		return
	}

	if db != nil {
		// Validate products exist and stock
		for _, item := range order.Items {
			var stock int
			row := db.QueryRow(`SELECT stock FROM products WHERE id = $1`, item.ID)
			if err := row.Scan(&stock); err != nil {
				if err == sql.ErrNoRows {
					c.JSON(http.StatusBadRequest, ErrorResponse{Success: false, Error: "Product not found", Message: fmt.Sprintf("Product with ID %d does not exist", item.ID)})
					return
				}
				c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
				return
			}
			if stock < item.Quantity {
				c.JSON(http.StatusBadRequest, ErrorResponse{Success: false, Error: "Insufficient stock", Message: fmt.Sprintf("Product %d only has %d items in stock", item.ID, stock)})
				return
			}
		}

		// Create order
		order.ID = uuid.New().String()
		orderCounter++
		order.OrderNumber = fmt.Sprintf("VUE-%d", orderCounter)
		order.Status = "pending"
		order.CreatedAt = time.Now()
		order.UpdatedAt = time.Now()

		tx, err := db.Begin()
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
			return
		}
		defer tx.Rollback()

		_, err = tx.Exec(`INSERT INTO orders (id, order_number, customer_name, customer_email, customer_address, subtotal, shipping, tax, total, status, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
			order.ID, order.OrderNumber, order.Customer.Name, order.Customer.Email, order.Customer.Address, order.Subtotal, order.Shipping, order.Tax, order.Total, order.Status, order.CreatedAt, order.UpdatedAt,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
			return
		}
		for _, item := range order.Items {
			_, err := tx.Exec(`INSERT INTO order_items (order_id, product_id, name, price, quantity) VALUES ($1,$2,$3,$4,$5)`, order.ID, item.ID, item.Name, item.Price, item.Quantity)
			if err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
				return
			}
			_, err = tx.Exec(`UPDATE products SET stock = stock - $1 WHERE id = $2`, item.Quantity, item.ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
				return
			}
		}
		if err := tx.Commit(); err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
			return
		}

		response := OrderResponse{Success: true, Message: "Order created successfully"}
		response.Data.OrderID = order.ID
		response.Data.OrderNumber = order.OrderNumber
		response.Data.Total = order.Total
		response.Data.Status = order.Status
		c.JSON(http.StatusCreated, response)
		return
	}

	// Fallback to mock logic
	var fallbackOrder Order
	if err := c.ShouldBindJSON(&fallbackOrder); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Validation failed",
			Message: err.Error(),
		})
		return
	}

	// Validate that all products exist and have sufficient stock
	for _, item := range fallbackOrder.Items {
		found := false
		for _, product := range mockProducts {
			if product.ID == item.ID {
				if product.Stock < item.Quantity {
					c.JSON(http.StatusBadRequest, ErrorResponse{
						Success: false,
						Error:   "Insufficient stock",
						Message: fmt.Sprintf("Product %s only has %d items in stock", product.Name, product.Stock),
					})
					return
				}
				found = true
				break
			}
		}
		if !found {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Success: false,
				Error:   "Product not found",
				Message: fmt.Sprintf("Product with ID %d does not exist", item.ID),
			})
			return
		}
	}

	// Generate order ID and number
	fallbackOrder.ID = uuid.New().String()
	orderCounter++
	fallbackOrder.OrderNumber = fmt.Sprintf("VUE-%d", orderCounter)
	fallbackOrder.Status = "pending"
	fallbackOrder.CreatedAt = time.Now()
	fallbackOrder.UpdatedAt = time.Now()

	// Store order
	orders[fallbackOrder.ID] = fallbackOrder

	// Prepare response
	response := OrderResponse{
		Success: true,
		Message: "Order created successfully",
	}
	response.Data.OrderID = fallbackOrder.ID
	response.Data.OrderNumber = fallbackOrder.OrderNumber
	response.Data.Total = fallbackOrder.Total
	response.Data.Status = fallbackOrder.Status

	c.JSON(http.StatusCreated, response)
}

// getAllOrders handles GET /api/orders
func getAllOrders(c *gin.Context) {
	orderList := make([]Order, 0, len(orders))
	for _, order := range orders {
		orderList = append(orderList, order)
	}

	// Sort orders by creation date (newest first)
	sort.Slice(orderList, func(i, j int) bool {
		return orderList[i].CreatedAt.After(orderList[j].CreatedAt)
	})

	response := OrdersResponse{
		Success: true,
		Data:    orderList,
		Total:   len(orderList),
	}

	c.JSON(http.StatusOK, response)
}

// getOrderByNumber handles GET /api/orders/:orderNumber
func getOrderByNumber(c *gin.Context) {
	orderNumber := c.Param("orderNumber")

	// Find order by order number
	for _, order := range orders {
		if order.OrderNumber == orderNumber {
			response := OrderDetailResponse{
				Success: true,
				Data:    order,
			}
			c.JSON(http.StatusOK, response)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{
		Success: false,
		Error:   "Order not found",
	})
}

// updateOrderStatus handles PUT /api/orders/:orderId/status
func updateOrderStatus(c *gin.Context) {
	orderID := c.Param("orderId")
	var statusUpdate OrderStatus

	if err := c.ShouldBindJSON(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid status",
			Message: err.Error(),
		})
		return
	}

	// Find and update order
	if order, exists := orders[orderID]; exists {
		order.Status = statusUpdate.Status
		order.UpdatedAt = time.Now()
		orders[orderID] = order

		response := OrderDetailResponse{
			Success: true,
			Data:    order,
		}
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusNotFound, ErrorResponse{
		Success: false,
		Error:   "Order not found",
	})
}

// deleteOrder handles DELETE /api/orders/:orderId
func deleteOrder(c *gin.Context) {
	orderID := c.Param("orderId")

	if _, exists := orders[orderID]; exists {
		delete(orders, orderID)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Order deleted successfully",
		})
		return
	}

	c.JSON(http.StatusNotFound, ErrorResponse{
		Success: false,
		Error:   "Order not found",
	})
}

// generateToken generates a random token
func generateToken() string {
	bytes := make([]byte, 16)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

// login handles POST /api/auth/login
func login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Success: false,
			Error:   "Invalid request",
			Message: err.Error(),
		})
		return
	}

	if db != nil {
		var (
			id                                        int
			username, email, name, role, passwordHash string
		)
		row := db.QueryRow(`SELECT id, username, email, name, role, password_hash FROM users WHERE username = $1`, loginReq.Username)
		if err := row.Scan(&id, &username, &email, &name, &role, &passwordHash); err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, ErrorResponse{Success: false, Error: "Invalid credentials"})
				return
			}
			c.JSON(http.StatusInternalServerError, ErrorResponse{Success: false, Error: "DB error", Message: err.Error()})
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(loginReq.Password)) != nil {
			c.JSON(http.StatusUnauthorized, ErrorResponse{Success: false, Error: "Invalid credentials"})
			return
		}
		user := User{ID: id, Username: username, Email: email, Name: name, Role: role}
		token := generateToken()
		activeSessions[token] = user
		resp := AuthResponse{Success: true, Message: "Login successful"}
		resp.Data.User = user
		resp.Data.Token = token
		c.JSON(http.StatusOK, resp)
		return
	}

	// Fallback to mock credentials
	expectedPassword, exists := userCredentials[loginReq.Username]
	if !exists || expectedPassword != loginReq.Password {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Success: false,
			Error:   "Invalid credentials",
			Message: "Username or password is incorrect",
		})
		return
	}
	var user User
	for _, u := range mockUsers {
		if u.Username == loginReq.Username {
			user = u
			break
		}
	}
	token := generateToken()
	activeSessions[token] = user
	response := AuthResponse{Success: true, Message: "Login successful"}
	response.Data.User = user
	response.Data.Token = token
	c.JSON(http.StatusOK, response)
}

// logout handles POST /api/auth/logout
func logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token != "" {
		// Remove "Bearer " prefix if present
		if strings.HasPrefix(token, "Bearer ") {
			token = token[7:]
		}
		delete(activeSessions, token)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Logout successful",
	})
}

// getCurrentUser handles GET /api/auth/me
func getCurrentUser(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Success: false,
			Error:   "No token provided",
		})
		return
	}

	// Remove "Bearer " prefix if present
	if strings.HasPrefix(token, "Bearer ") {
		token = token[7:]
	}

	user, exists := activeSessions[token]
	if !exists {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Success: false,
			Error:   "Invalid token",
		})
		return
	}

	response := UserResponse{
		Success: true,
		Data:    user,
	}

	c.JSON(http.StatusOK, response)
}
