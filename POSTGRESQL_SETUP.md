# 🗄️ PostgreSQL Database Setup Guide

This guide provides step-by-step instructions for setting up PostgreSQL with the Vue Shop application.

## 🎯 What We've Accomplished

✅ **PostgreSQL 17** installed and configured  
✅ **Database 'vueshop'** created with proper user permissions  
✅ **Go backend** integrated with PostgreSQL using lib/pq driver  
✅ **Automatic migrations** - tables created on startup  
✅ **Data seeding** - sample products and users loaded automatically  
✅ **Admin panel** - web-based database management interface  
✅ **pgAdmin 4** - professional PostgreSQL administration tool  

## 🔐 Database Credentials

### PostgreSQL Superuser
- **Username**: postgres
- **Password**: mypostgres *(keep this secure!)*

### Application User
- **Username**: vueshop
- **Password**: vueshop123
- **Database**: vueshop
- **Permissions**: Full access to vueshop database

## 🌐 Access Points

| Service | URL | Credentials |
|---------|-----|-------------|
| **Frontend** | http://localhost:3000 | - |
| **Backend API** | http://localhost:5000 | - |
| **pgAdmin** | http://localhost:5050 | admin@admin.com / root |
| **Database Admin Panel** | http://localhost:3000/admin/database | admin / admin123 |

## 🚀 Quick Start

### 1. Start the Application
```bash
npm start
```

### 2. Access the Admin Panel
1. Navigate to http://localhost:3000
2. Login with admin/admin123
3. Click "🗄️ Database Admin" in the navigation
4. Use the admin panel to manage your database

### 3. Use pgAdmin (Optional)
1. Open http://localhost:5050
2. Login with admin@admin.com / root
3. Connect to your local PostgreSQL server
4. Browse and manage the vueshop database

## 📊 Database Schema

### Tables Created Automatically

```sql
-- Users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT NOT NULL,
    name TEXT NOT NULL,
    role TEXT NOT NULL,
    password_hash TEXT NOT NULL
);

-- Products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    category TEXT NOT NULL,
    image TEXT NOT NULL,
    stock INT NOT NULL
);

-- Orders table
CREATE TABLE orders (
    id UUID PRIMARY KEY,
    order_number TEXT UNIQUE NOT NULL,
    customer_name TEXT NOT NULL,
    customer_email TEXT NOT NULL,
    customer_address TEXT NOT NULL,
    subtotal NUMERIC(10,2) NOT NULL,
    shipping NUMERIC(10,2) NOT NULL,
    tax NUMERIC(10,2) NOT NULL,
    total NUMERIC(10,2) NOT NULL,
    status TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Order items table
CREATE TABLE order_items (
    id SERIAL PRIMARY KEY,
    order_id UUID NOT NULL REFERENCES orders(id) ON DELETE CASCADE,
    product_id INT NOT NULL REFERENCES products(id),
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    quantity INT NOT NULL
);
```

## 🔧 Admin Panel Features

### Real-time Statistics
- **Products Count**: Total products in database
- **Users Count**: Registered user accounts
- **Orders Count**: Total orders placed
- **Total Value**: Sum of all order totals

### Database Operations
- **🌱 Seed Database**: Load sample data
- **🗑️ Clear All Data**: Remove all data (with confirmation)
- **📤 Export Data**: Download database as JSON
- **💾 Backup DB**: Information about pgAdmin backup

### Data Management
- **Products Table**: View, edit, delete products
- **Users Table**: View, edit, delete users
- **Connection Testing**: Test database connectivity
- **pgAdmin Integration**: Direct access to pgAdmin

## 🛠️ Manual Database Operations

### Connect via psql
```bash
# Connect as vueshop user
psql -h localhost -U vueshop -d vueshop

# Connect as postgres superuser
psql -h localhost -U postgres -d postgres
```

### Backup Database
```bash
# Create backup
pg_dump -h localhost -U vueshop -d vueshop > vueshop_backup.sql

# Restore backup
psql -h localhost -U vueshop -d vueshop < vueshop_backup.sql
```

### Reset Database
```bash
# Drop and recreate database
dropdb -h localhost -U postgres vueshop
createdb -h localhost -U postgres vueshop

# Re-run setup script
powershell -ExecutionPolicy Bypass -File setup-db.ps1
```

## 🔍 Troubleshooting

### Common Issues

#### 1. Connection Refused
```bash
# Check if PostgreSQL is running
Get-Service postgresql*

# Start service if stopped
Start-Service postgresql-x64-17
```

#### 2. Authentication Failed
```bash
# Verify environment variables
echo $env:PGHOST
echo $env:PGUSER
echo $env:PGPASSWORD

# Reset environment variables
$env:PGHOST="localhost"
$env:PGPORT="5432"
$env:PGUSER="vueshop"
$env:PGPASSWORD="vueshop123"
$env:PGDATABASE="vueshop"
```

#### 3. Database Not Found
```bash
# Recreate database
cd backend-go
powershell -ExecutionPolicy Bypass -File setup-db.ps1
```

#### 4. Permission Denied
```bash
# Grant permissions as postgres user
psql -h localhost -U postgres -d vueshop -c "GRANT ALL ON SCHEMA public TO vueshop;"
psql -h localhost -U postgres -d vueshop -c "GRANT ALL PRIVILEGES ON DATABASE vueshop TO vueshop;"
```

## 📈 Performance Tips

### Connection Pooling
- Backend automatically manages connection pool
- Max connections: 10
- Connection lifetime: 2 hours
- Idle connections: 5

### Query Optimization
- All queries use parameterized statements
- Proper indexing on primary keys
- Efficient JOIN operations for orders

### Monitoring
- Use pgAdmin to monitor query performance
- Check connection pool status
- Monitor database size and growth

## 🔐 Security Best Practices

### Password Security
- ✅ Use strong passwords for production
- ✅ Never commit passwords to version control
- ✅ Use environment variables for sensitive data
- ✅ Regularly rotate database passwords

### Access Control
- ✅ Application user has minimal required permissions
- ✅ Superuser access limited to administration
- ✅ Network access restricted to localhost
- ✅ SSL enabled for production deployments

### Data Protection
- ✅ Passwords hashed with bcrypt
- ✅ Input validation on all endpoints
- ✅ SQL injection protection via parameterized queries
- ✅ CORS properly configured

## 🎉 Next Steps

### Immediate Actions
1. **Test the admin panel** - Navigate to /admin/database
2. **Seed the database** - Click "🌱 Seed Database" button
3. **Explore pgAdmin** - Open http://localhost:5050
4. **Test authentication** - Login with admin/admin123

### Future Enhancements
1. **Add more admin features** - User management, order tracking
2. **Implement data analytics** - Sales reports, user statistics
3. **Add backup automation** - Scheduled database backups
4. **Performance monitoring** - Query performance tracking

---

**🎯 Your Vue Shop now has a fully functional PostgreSQL database with professional administration tools!**

**🔐 Remember**: Keep your postgres password secure and never share it publicly.
