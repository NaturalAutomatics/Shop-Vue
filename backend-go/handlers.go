package main

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// getProducts handles GET /api/products
func getProducts(c *gin.Context) {
	// Get query parameters
	category := c.Query("category")
	search := c.Query("search")
	sortBy := c.Query("sort")

	// Filter products
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

	// Find product by ID
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
	// Get unique categories
	categoryMap := make(map[string]bool)
	for _, product := range mockProducts {
		categoryMap[product.Category] = true
	}

	categories := make([]string, 0, len(categoryMap))
	for category := range categoryMap {
		categories = append(categories, category)
	}

	// Sort categories alphabetically
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

	// Validate that all products exist and have sufficient stock
	for _, item := range order.Items {
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
	order.ID = uuid.New().String()
	orderCounter++
	order.OrderNumber = fmt.Sprintf("VUE-%d", orderCounter)
	order.Status = "pending"
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	// Store order
	orders[order.ID] = order

	// Prepare response
	response := OrderResponse{
		Success: true,
		Message: "Order created successfully",
	}
	response.Data.OrderID = order.ID
	response.Data.OrderNumber = order.OrderNumber
	response.Data.Total = order.Total
	response.Data.Status = order.Status

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
