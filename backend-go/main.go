package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin mode
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Connect to Postgres (if configured via env)
	if err := connectPostgres(); err != nil {
		log.Printf("Postgres not connected: %v (set PG* env vars to enable DB)", err)
	} else {
		if err := migratePostgres(); err != nil {
			log.Printf("Postgres migration failed: %v", err)
		} else if err := seedPostgresIfEmpty(); err != nil {
			log.Printf("Postgres seed failed: %v", err)
		} else {
			log.Printf("Postgres connected and ready")
		}
	}

	// Create router
	r := gin.Default()

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"http://localhost:3001",
		"http://127.0.0.1:3000",
		"http://127.0.0.1:3001",
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	// Health check endpoint
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "OK",
			"message":   "Vue Shop Backend is running",
			"language":  "Go",
			"framework": "Gin",
			"version":   "1.0.0",
		})
	})

	// API routes
	api := r.Group("/api")
	{
		// Auth routes
		auth := api.Group("/auth")
		{
			auth.POST("/login", login)
			auth.POST("/logout", logout)
			auth.GET("/me", getCurrentUser)
		}

		// Product routes
		api.GET("/products", getProducts)
		api.GET("/products/:id", getProduct)
		api.GET("/products/categories", getCategories)

		// Order routes
		api.POST("/orders", createOrder)
		api.GET("/orders", getAllOrders)
		api.GET("/orders/:orderNumber", getOrderByNumber)
		api.PUT("/orders/:orderId/status", updateOrderStatus)
		api.DELETE("/orders/:orderId", deleteOrder)

		// Admin routes
		admin := api.Group("/admin")
		{
			admin.POST("/test-connection", testConnection)
			admin.GET("/stats", getStats)
			admin.GET("/users", getAllUsers)
			admin.POST("/seed", seedDatabase)
			admin.POST("/clear", clearDatabase)
			admin.GET("/export", exportData)
			admin.DELETE("/products/:id", deleteProduct)
			admin.DELETE("/users/:id", deleteUser)
		}
	}

	// Get port from environment or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	// Start server
	log.Printf("🚀 Vue Shop Backend (Go) running on port %s", port)
	log.Printf("📊 Health check: http://localhost:%s/health", port)
	log.Printf("🛍️ API Base URL: http://localhost:%s/api", port)
	log.Printf("🌐 CORS enabled for: http://localhost:3000")

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
