package main

import (
	"time"
)

// Product represents a product in the shop
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Stock       int     `json:"stock"`
}

// Customer represents customer information
type Customer struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Address string `json:"address" binding:"required"`
}

// OrderItem represents an item in an order
type OrderItem struct {
	ID       int     `json:"id" binding:"required"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity" binding:"required,min=1"`
}

// Order represents a complete order
type Order struct {
	ID          string      `json:"id"`
	OrderNumber string      `json:"orderNumber"`
	Customer    Customer    `json:"customer" binding:"required"`
	Items       []OrderItem `json:"items" binding:"required,min=1"`
	Subtotal    float64     `json:"subtotal" binding:"required,min=0"`
	Shipping    float64     `json:"shipping" binding:"required,min=0"`
	Tax         float64     `json:"tax" binding:"required,min=0"`
	Total       float64     `json:"total" binding:"required,min=0"`
	Status      string      `json:"status"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

// OrderStatus represents order status update request
type OrderStatus struct {
	Status string `json:"status" binding:"required,oneof=pending processing shipped delivered cancelled"`
}

// API Response structures
type ProductsResponse struct {
	Success bool      `json:"success"`
	Data    []Product `json:"data"`
	Total   int       `json:"total"`
}

type ProductResponse struct {
	Success bool    `json:"success"`
	Data    Product `json:"data"`
}

type CategoriesResponse struct {
	Success bool     `json:"success"`
	Data    []string `json:"data"`
}

type OrderResponse struct {
	Success bool  `json:"success"`
	Message string `json:"message"`
	Data    struct {
		OrderID      string  `json:"orderId"`
		OrderNumber  string  `json:"orderNumber"`
		Total        float64 `json:"total"`
		Status       string  `json:"status"`
	} `json:"data"`
}

type OrdersResponse struct {
	Success bool    `json:"success"`
	Data    []Order `json:"data"`
	Total   int     `json:"total"`
}

type OrderDetailResponse struct {
	Success bool  `json:"success"`
	Data    Order `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

// Mock data
var mockProducts = []Product{
	{
		ID:          1,
		Name:        "Vue.js T-Shirt",
		Description: "Comfortable cotton t-shirt with Vue.js logo",
		Price:       25.99,
		Category:    "clothing",
		Image:       "https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?w=400&h=300&fit=crop&bg=white",
		Stock:       50,
	},
	{
		ID:          2,
		Name:        "JavaScript Book",
		Description: "Comprehensive guide to modern JavaScript",
		Price:       39.99,
		Category:    "books",
		Image:       "https://images.unsplash.com/photo-1544947950-fa07a98d237f?w=400&h=300&fit=crop&bg=white",
		Stock:       25,
	},
	{
		ID:          3,
		Name:        "Wireless Headphones",
		Description: "High-quality wireless headphones with noise cancellation",
		Price:       89.99,
		Category:    "electronics",
		Image:       "https://images.unsplash.com/photo-1505740420928-5e560c06d30e?w=400&h=300&fit=crop&bg=white",
		Stock:       15,
	},
	{
		ID:          4,
		Name:        "Vue Hoodie",
		Description: "Warm and cozy hoodie perfect for Vue developers",
		Price:       45.99,
		Category:    "clothing",
		Image:       "https://images.unsplash.com/photo-1556821840-3a63f95609a7?w=400&h=300&fit=crop&bg=white",
		Stock:       30,
	},
	{
		ID:          5,
		Name:        "Laptop Stand",
		Description: "Ergonomic laptop stand for better posture",
		Price:       29.99,
		Category:    "electronics",
		Image:       "https://images.unsplash.com/photo-1586953208448-b95a79798f07?w=400&h=300&fit=crop&bg=white",
		Stock:       20,
	},
	{
		ID:          6,
		Name:        "Vue.js Guide",
		Description: "Complete Vue.js 3 tutorial and reference",
		Price:       19.99,
		Category:    "books",
		Image:       "https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=400&h=300&fit=crop&bg=white",
		Stock:       40,
	},
	{
		ID:          7,
		Name:        "Mechanical Keyboard",
		Description: "Premium mechanical keyboard with RGB backlighting",
		Price:       129.99,
		Category:    "electronics",
		Image:       "https://images.unsplash.com/photo-1541140532154-b024d705b90a?w=400&h=300&fit=crop&bg=white",
		Stock:       10,
	},
	{
		ID:          8,
		Name:        "Developer Mug",
		Description: "Ceramic coffee mug perfect for coding sessions",
		Price:       12.99,
		Category:    "clothing",
		Image:       "https://images.unsplash.com/photo-1578662996442-48f60103fc96?w=400&h=300&fit=crop&bg=white",
		Stock:       100,
	},
	{
		ID:          9,
		Name:        "React vs Vue Book",
		Description: "In-depth comparison of modern JavaScript frameworks",
		Price:       34.99,
		Category:    "books",
		Image:       "https://images.unsplash.com/photo-1589829085413-56de8ae18c73?w=400&h=300&fit=crop&bg=white",
		Stock:       35,
	},
}

// In-memory storage for orders
var orders = make(map[string]Order)
var orderCounter = 1000
