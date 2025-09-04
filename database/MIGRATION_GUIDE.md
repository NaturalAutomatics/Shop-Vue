# PostgreSQL Database Migration Guide

## Requirements

### PostgreSQL Version
- **Minimum**: PostgreSQL 12+
- **Recommended**: PostgreSQL 15+ or 16+
- **Extensions**: uuid-ossp (for UUID generation)

### Hosting Providers (Free/Paid Options)

#### Free Options:
1. **Supabase** (Recommended)
   - 500MB database
   - 2 projects
   - 50,000 monthly active users
   - Built-in auth and real-time features

2. **Railway**
   - $5/month after trial
   - PostgreSQL 13+
   - 1GB storage

3. **Render**
   - Free tier: 90 days, then $7/month
   - PostgreSQL 14+
   - 1GB storage

#### Paid Options:
1. **AWS RDS**
2. **Google Cloud SQL**
3. **Azure Database**
4. **DigitalOcean Managed Databases**

## Migration Steps

### 1. Export Current Data (Optional)
```bash
# Connect to your local database
psql -h localhost -U vueshop -d vueshop

# Run export script
\i database/export_data.sql
```

### 2. Set Up Remote Database

#### For Supabase:
1. Go to [supabase.com](https://supabase.com)
2. Create new project
3. Get connection details from Settings > Database

#### For Railway:
1. Go to [railway.app](https://railway.app)
2. Create new project
3. Add PostgreSQL service
4. Get connection string from Variables tab

### 3. Run Migration
```bash
# Connect to remote database
psql "postgresql://username:password@host:port/database"

# Run migration
\i database/migration.sql

# Run seed data
\i database/seed_data.sql
```

### 4. Update Backend Configuration

#### Environment Variables Needed:
```bash
DATABASE_URL=postgresql://username:password@host:port/database
PGHOST=your-host
PGPORT=5432
PGUSER=your-username
PGPASSWORD=your-password
PGDATABASE=your-database
PGSSLMODE=require
```

#### Update Go Backend:
Replace hardcoded connection string in `backend-go/db.go`:
```go
func getPostgresConnString() string {
    if dbURL := os.Getenv("DATABASE_URL"); dbURL != "" {
        return dbURL
    }
    
    host := getenv("PGHOST", "localhost")
    port := getenv("PGPORT", "5432")
    user := getenv("PGUSER", "vueshop")
    password := getenv("PGPASSWORD", "vueshop123")
    dbname := getenv("PGDATABASE", "vueshop")
    sslmode := getenv("PGSSLMODE", "disable")
    
    return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        host, port, user, password, dbname, sslmode)
}
```

## Database Schema

### Tables:
- **users**: User accounts with bcrypt passwords
- **products**: Product catalog
- **orders**: Customer orders with UUID IDs
- **order_items**: Order line items

### Features:
- UUID primary keys for orders
- Automatic timestamps (created_at, updated_at)
- Data validation constraints
- Performance indexes
- Foreign key relationships

## Connection Details Template

```
Host: your-remote-host
Port: 5432
Database: your-database-name
Username: your-username
Password: your-password
SSL Mode: require (for production)
```

## Security Notes

1. **Always use SSL** in production (`sslmode=require`)
2. **Use environment variables** for credentials
3. **Create dedicated database user** with limited permissions
4. **Enable connection pooling** for better performance
5. **Regular backups** are essential

## Testing Connection

```bash
# Test connection
psql "postgresql://username:password@host:port/database?sslmode=require"

# Test from Go application
go run . --test-db
```

## Troubleshooting

### Common Issues:
1. **SSL Certificate errors**: Use `sslmode=require` or `sslmode=verify-full`
2. **Connection timeouts**: Check firewall settings
3. **Authentication failed**: Verify username/password
4. **Database not found**: Ensure database exists

### Performance Tips:
1. Use connection pooling (already configured)
2. Add indexes for frequently queried columns
3. Monitor query performance
4. Use prepared statements (already implemented)