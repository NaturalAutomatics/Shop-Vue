# Vue Shop Backend (Go)

A high-performance Go backend API for the Vue Shop e-commerce application using the Gin framework.

## 🚀 Features

- **High Performance**: Built with Go and Gin framework
- **Product Management**: Get products with filtering, search, and sorting
- **Order Management**: Create, retrieve, and update orders
- **Order Numbers**: Unique order number generation (VUE-1001, VUE-1002, etc.)
- **Input Validation**: Comprehensive request validation using Gin's binding
- **CORS Support**: Configured for frontend communication
- **Mock Data**: In-memory storage with sample products
- **MySQL Ready**: Complete database schema provided

## 🛠️ Technology Stack

- **Language**: Go 1.21+
- **Framework**: Gin (HTTP web framework)
- **Validation**: Gin binding with custom validators
- **Database**: MySQL (schema provided)
- **Storage**: In-memory (mock data)

## 📋 API Endpoints

### Products

- `GET /api/products` - Get all products
- `GET /api/products/:id` - Get single product
- `GET /api/products/categories` - Get all categories

**Query Parameters:**
- `category` - Filter by category (electronics, clothing, books)
- `search` - Search in name and description
- `sort` - Sort by price or name (price-asc, price-desc, name-asc, name-desc)

### Orders

- `POST /api/orders` - Create new order
- `GET /api/orders` - Get all orders (admin)
- `GET /api/orders/:orderNumber` - Get order by order number
- `PUT /api/orders/:orderId/status` - Update order status
- `DELETE /api/orders/:orderId` - Delete order (admin)

## 🛠️ Installation

### Prerequisites

- Go 1.21 or higher
- Git

### Setup

1. **Clone or navigate to the backend directory:**
   ```bash
   cd backend-go
   ```

2. **Install dependencies:**
   ```bash
   go mod tidy
   ```

3. **Run the server:**
   ```bash
   go run .
   ```

   Or build and run:
   ```bash
   go build -o vue-shop-backend .
   ./vue-shop-backend
   ```

4. **Access the API:**
   - Health check: http://localhost:5000/health
   - API Base: http://localhost:5000/api

## 🗄️ Database Setup

### MySQL Database

1. **Run the SQL script:**
   ```bash
   mysql -u your_username -p < database.sql
   ```

2. **Or copy and paste the contents of `database.sql` into your MySQL client**

### Database Features

- **4 Tables**: products, customers, orders, order_items
- **Sample Data**: 9 products pre-loaded
- **Stored Procedures**: For order creation and management
- **Triggers**: Automatic stock management
- **Views**: Order summary view
- **Indexes**: Optimized for performance

## 📊 API Examples

### Create Order
```bash
POST http://localhost:5000/api/orders
Content-Type: application/json

{
  "customer": {
    "name": "John Doe",
    "email": "john@example.com",
    "address": "123 Main St, City, Country"
  },
  "items": [
    {
      "id": 1,
      "name": "Vue.js T-Shirt",
      "price": 25.99,
      "quantity": 2
    }
  ],
  "subtotal": 51.98,
  "shipping": 5.99,
  "tax": 4.16,
  "total": 62.13
}
```

**Response:**
```json
{
  "success": true,
  "message": "Order created successfully",
  "data": {
    "orderId": "uuid-here",
    "orderNumber": "VUE-1001",
    "total": 62.13,
    "status": "pending"
  }
}
```

### Get Products with Filters
```bash
GET http://localhost:5000/api/products?category=electronics&sort=price-asc
```

## 🔧 Configuration

### Environment Variables

- `PORT` - Server port (default: 5000)
- `GIN_MODE` - Gin mode (debug/release)

### CORS Configuration

The server is configured to allow requests from:
- http://localhost:3000
- http://localhost:3001

## 🏗️ Project Structure

```
backend-go/
├── main.go           # Server entry point
├── models.go         # Data models and mock data
├── handlers.go       # HTTP handlers
├── database.sql      # MySQL database schema
├── go.mod           # Go module file
├── go.sum           # Go dependencies checksum
├── .gitignore       # Git ignore rules
└── README.md        # This file
```

## 🧪 Testing

```bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...
```

## 🚀 Deployment

### Build for Production

```bash
# Build for current platform
go build -o vue-shop-backend .

# Build for specific platform (Linux)
GOOS=linux GOARCH=amd64 go build -o vue-shop-backend .

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o vue-shop-backend.exe .
```

### Docker (Optional)

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o vue-shop-backend .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/vue-shop-backend .
EXPOSE 5000
CMD ["./vue-shop-backend"]
```

## 📝 Development

### Live Reload (Optional)

Install Air for live reloading:

```bash
# Install Air
go install github.com/cosmtrek/air@latest

# Run with Air
air
```

### Code Formatting

```bash
# Format code
go fmt ./...

# Run linter
golangci-lint run
```

## 🔒 Security Features

- **Input Validation**: All requests are validated
- **CORS**: Configured for frontend communication
- **Error Handling**: Proper HTTP status codes
- **No SQL Injection**: Using prepared statements (when using database)

## 📈 Performance

- **High Performance**: Go's efficient runtime
- **Low Memory Usage**: Minimal memory footprint
- **Fast Startup**: Quick application startup
- **Concurrent**: Handles multiple requests efficiently

## 🔄 Migration from Node.js

This Go backend provides the same API as the Node.js version:

- Same endpoints and response formats
- Same validation rules
- Same error handling
- Compatible with existing Vue.js frontend

## 📝 Future Enhancements

- Database integration (MySQL/PostgreSQL)
- User authentication (JWT)
- Payment gateway integration
- Email notifications
- Admin dashboard
- Inventory management
- Analytics and reporting
- Rate limiting
- Caching (Redis)
- Logging and monitoring

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## 📄 License

MIT License - see LICENSE file for details

---

Happy coding! 🚀
