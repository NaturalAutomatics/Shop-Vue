package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define metrics
var (
	httpRequestsTotal = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Total number of HTTP requests",
	}, []string{"method", "path", "status_code"})

	httpRequestDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests in seconds",
		Buckets: []float64{0.1, 0.5, 1, 2, 5},
	}, []string{"method", "path"})

	httpRequestsInProgress = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "http_requests_in_progress",
		Help: "Current number of HTTP requests in progress",
	})
)

// PrometheusMiddleware tracks request metrics
func PrometheusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.FullPath()
		if path == "" {
			path = "unknown"
		}

		// Increment in-progress requests
		httpRequestsInProgress.Inc()

		// Process request
		c.Next()

		// Record metrics after request is processed
		duration := time.Since(start).Seconds()
		statusCode := c.Writer.Status()

		httpRequestDuration.WithLabelValues(c.Request.Method, path).Observe(duration)
		httpRequestsTotal.WithLabelValues(c.Request.Method, path, string(statusCode)).Inc()

		// Decrement in-progress requests
		httpRequestsInProgress.Dec()
	}
}

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

	// Add Prometheus middleware
	r.Use(PrometheusMiddleware())

	// CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{
		"http://localhost:3000",
		"http://localhost:3001",
		"http://localhost:3002",
		"http://127.0.0.1:3000",
		"http://127.0.0.1:3001",
		"http://127.0.0.1:3002",
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

	// Add Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

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
	log.Printf("📈 Metrics endpoint: http://localhost:%s/metrics", port)
	log.Printf("🛍️ API Base URL: http://localhost:%s/api", port)
	log.Printf("🌐 CORS enabled for: http://localhost:3000, 3001, 3002")

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
