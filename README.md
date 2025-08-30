# 🛍️ Vue Shop - Full-Stack E-commerce Application

A modern e-commerce application built with Vue.js 3 frontend and Go backend, featuring PostgreSQL database integration and comprehensive admin panel.

## 🚀 Features

- **Frontend**: Vue.js 3 with Composition API, Pinia state management, Vue Router 4
- **Backend**: Go with Gin framework, PostgreSQL database
- **Authentication**: JWT-based auth with bcrypt password hashing
- **Database**: PostgreSQL with automatic migrations and seeding
- **Admin Panel**: Web-based database administration with real-time statistics
- **Responsive Design**: Mobile-first design with adaptive product grid
- **Real-time Data**: Live database statistics and management

## 🗄️ Database Setup

### Prerequisites
- PostgreSQL 17+ installed
- pgAdmin 4 for database management (optional but recommended)

### Installation Steps

1. **Install PostgreSQL**
   ```bash
   winget install PostgreSQL.PostgreSQL.17
   ```

2. **Install pgAdmin (Optional)**
   ```bash
   winget install PostgreSQL.pgAdmin
   ```

3. **Run Database Setup Script**
   ```bash
   cd backend-go
   powershell -ExecutionPolicy Bypass -File setup-db.ps1
   ```
   
   **Note**: When prompted for PostgreSQL password, use: `mypostgres`

4. **Set Environment Variables**
   ```bash
   $env:PGHOST="localhost"
   $env:PGPORT="5432"
   $env:PGUSER="vueshop"
   $env:PGPASSWORD="vueshop123"
   $env:PGDATABASE="vueshop"
   ```

### Database Connection Details
- **Host**: localhost
- **Port**: 5432
- **Database**: vueshop
- **Username**: vueshop
- **Password**: vueshop123

## 🏃‍♂️ Running the Application

### Option 1: Unified Startup Script
```bash
npm start
```

### Option 2: Manual Startup
1. **Start Backend**
   ```bash
   cd backend-go
   go run .
   ```

2. **Start Frontend**
   ```bash
   cd frontend-vue
   npm run dev
   ```

### Option 3: PowerShell Script
```bash
powershell -ExecutionPolicy Bypass -File start-app.ps1
```

## 🌐 Access Points

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:5000
- **pgAdmin**: http://localhost:5050
  - Email: admin@admin.com
  - Password: root

## 🗄️ Database Admin Panel

Access the built-in database administration panel at `/admin/database` (requires admin login).

### Features
- **Real-time Statistics**: Products, users, orders, and total value
- **Connection Testing**: Test database connectivity
- **Data Management**: View, delete products and users
- **Database Operations**: Seed, clear, and export data
- **pgAdmin Integration**: Direct access to pgAdmin web interface

### Admin Login
- **Username**: admin
- **Password**: admin123

## 📊 Database Schema

### Tables
- **users**: User accounts with bcrypt-hashed passwords
- **products**: Product catalog with stock management
- **orders**: Order tracking with customer details
- **order_items**: Order line items with quantities

### Automatic Features
- **Migrations**: Tables created automatically on startup
- **Seeding**: Sample data loaded if tables are empty
- **Indexing**: Optimized queries for performance

## 🔧 Development Commands

```bash
# Root Level
npm run dev              # Start both frontend and backend
npm run build            # Build both parts
npm run install:all      # Install frontend dependencies

# Frontend (in frontend-vue/)
npm run dev              # Start development server
npm run build            # Build for production
npm run type-check       # TypeScript checking
npm run preview          # Preview production build

# Backend (in backend-go/)
go run .                 # Run backend
go build -o vue-shop-backend.exe  # Build executable
go mod tidy              # Update dependencies

# Database
cd backend-go
powershell -ExecutionPolicy Bypass -File setup-db.ps1  # Setup database
```

## 🛠️ Technology Stack

### Frontend
- **Vue.js 3**: Progressive JavaScript framework
- **Pinia**: State management for Vue
- **Vue Router 4**: Client-side routing
- **Vite**: Fast build tool and dev server

### Backend
- **Go**: High-performance programming language
- **Gin**: HTTP web framework
- **PostgreSQL**: Advanced open-source database
- **bcrypt**: Secure password hashing

### Database Tools
- **pgAdmin**: Web-based PostgreSQL administration
- **lib/pq**: Go PostgreSQL driver
- **Automatic Migrations**: Schema management

## 📁 Project Structure

```
Shop/
├── frontend-vue/          # Vue.js frontend
│   ├── src/              # Vue source code
│   │   ├── views/        # Page components
│   │   ├── stores/       # Pinia state management
│   │   ├── services/     # API services
│   │   ├── types/        # TypeScript interfaces
│   │   └── router/       # Vue Router configuration
│   ├── package.json      # Frontend dependencies
│   ├── tsconfig.json     # TypeScript configuration
│   └── vite.config.js    # Vite build configuration
├── backend-go/           # Go backend
│   ├── db.go            # Database connection & migrations
│   ├── admin.go         # Admin API endpoints
│   ├── handlers.go      # Main API handlers
│   ├── models.go        # Data structures
│   └── main.go          # Server entry point
├── package.json         # Root project management
├── start-full-stack.ps1 # Full stack startup script
└── README.md            # This file
```

## 🔐 Security Features

- **Password Hashing**: bcrypt with configurable cost
- **JWT Tokens**: Secure authentication
- **CORS Protection**: Cross-origin request security
- **Input Validation**: Request data sanitization
- **SQL Injection Protection**: Parameterized queries

## 📈 Performance Features

- **Connection Pooling**: Optimized database connections
- **Query Optimization**: Efficient SQL queries
- **Caching**: In-memory session management
- **Async Operations**: Non-blocking I/O operations

## 🚨 Troubleshooting

### Database Connection Issues
1. Ensure PostgreSQL service is running
2. Verify environment variables are set correctly
3. Check firewall settings for port 5432
4. Verify database and user exist

### Backend Issues
1. Check environment variables
2. Verify PostgreSQL is accessible
3. Check logs for specific error messages
4. Ensure all dependencies are installed

### Frontend Issues
1. Verify backend is running on port 5000
2. Check browser console for errors
3. Ensure CORS is configured correctly
4. Verify API endpoints are accessible

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## 📄 License

This project is licensed under the MIT License.

## 🆘 Support

For issues and questions:
1. Check the troubleshooting section
2. Review the logs for error messages
3. Verify all prerequisites are met
4. Check the database connection status

---

**Note**: Keep your PostgreSQL password secure and never share it in public repositories or discussions.
