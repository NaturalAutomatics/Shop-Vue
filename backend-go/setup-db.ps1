# PostgreSQL Database Setup Script for Vue Shop
# Run this script after PostgreSQL installation

Write-Host "Setting up PostgreSQL database for Vue Shop..." -ForegroundColor Green

# Default PostgreSQL installation path
$pgBinPath = "C:\Program Files\PostgreSQL\17\bin"
$pgDataPath = "C:\Program Files\PostgreSQL\17\data"

# Check if PostgreSQL is installed
if (-not (Test-Path $pgBinPath)) {
    Write-Host "PostgreSQL not found at $pgBinPath" -ForegroundColor Red
    Write-Host "Please install PostgreSQL first or update the path in this script" -ForegroundColor Yellow
    exit 1
}

Write-Host "PostgreSQL found at $pgBinPath" -ForegroundColor Green

# Set environment variables for the current session
$env:PGHOST = "localhost"
$env:PGPORT = "5432"
$env:PGUSER = "postgres"
$env:PGDATABASE = "postgres"

Write-Host "Setting up database..." -ForegroundColor Yellow

# Create database and user
try {
    # Create vueshop database
    Write-Host "Creating 'vueshop' database..." -ForegroundColor Cyan
    & "$pgBinPath\createdb.exe" -U postgres -h localhost vueshop
    if ($LASTEXITCODE -eq 0) {
        Write-Host "Database 'vueshop' created successfully" -ForegroundColor Green
    } else {
        Write-Host "Database 'vueshop' might already exist" -ForegroundColor Yellow
    }

    # Create vueshop user
    Write-Host "Creating 'vueshop' user..." -ForegroundColor Cyan
    & "$pgBinPath\psql.exe" -U postgres -h localhost -d vueshop -c "CREATE USER vueshop WITH PASSWORD 'vueshop123';"
    if ($LASTEXITCODE -eq 0) {
        Write-Host "User 'vueshop' created successfully" -ForegroundColor Green
    } else {
        Write-Host "User 'vueshop' might already exist" -ForegroundColor Yellow
    }

    # Grant privileges
    Write-Host "Granting privileges to 'vueshop' user..." -ForegroundColor Cyan
    & "$pgBinPath\psql.exe" -U postgres -h localhost -d vueshop -c "GRANT ALL PRIVILEGES ON DATABASE vueshop TO vueshop;"
    & "$pgBinPath\psql.exe" -U postgres -h localhost -d vueshop -c "GRANT ALL ON SCHEMA public TO vueshop;"
    Write-Host "Privileges granted successfully" -ForegroundColor Green

} catch {
    Write-Host "Error setting up database: $($_.Exception.Message)" -ForegroundColor Red
    exit 1
}

Write-Host "Database setup completed!" -ForegroundColor Green
Write-Host "Connection Details:" -ForegroundColor Cyan
Write-Host "   Host: localhost" -ForegroundColor White
Write-Host "   Port: 5432" -ForegroundColor White
Write-Host "   Database: vueshop" -ForegroundColor White
Write-Host "   Username: vueshop" -ForegroundColor White
Write-Host "   Password: vueshop123" -ForegroundColor White

Write-Host "Environment Variables to set:" -ForegroundColor Cyan
Write-Host "   PGHOST=localhost" -ForegroundColor White
Write-Host "   PGPORT=5432" -ForegroundColor White
Write-Host "   PGUSER=vueshop" -ForegroundColor White
Write-Host "   PGPASSWORD=vueshop123" -ForegroundColor White
Write-Host "   PGDATABASE=vueshop" -ForegroundColor White

Write-Host "Next steps:" -ForegroundColor Yellow
Write-Host "   1. Set the environment variables above" -ForegroundColor White
Write-Host "   2. Start the backend with: go run ." -ForegroundColor White
Write-Host "   3. Open pgAdmin to manage the database" -ForegroundColor White
Write-Host "   4. The backend will automatically create tables and seed data" -ForegroundColor White

Write-Host "pgAdmin Access:" -ForegroundColor Cyan
Write-Host "   URL: http://localhost:5050" -ForegroundColor White
Write-Host "   Email: admin@admin.com" -ForegroundColor White
Write-Host "   Password: root" -ForegroundColor White
