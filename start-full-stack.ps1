# Vue Shop - Full Stack Startup Script
# Starts PostgreSQL, Backend (Go), and Frontend (Vue.js) with one command

Write-Host "Starting Vue Shop Full Stack Application..." -ForegroundColor Green
Write-Host ""

# Function to check if a service is running
function Test-ServiceRunning {
    param([string]$ServiceName)
    $service = Get-Service -Name $ServiceName -ErrorAction SilentlyContinue
    return $service -and $service.Status -eq 'Running'
}

# Function to check if a port is in use
function Test-Port {
    param([int]$Port)
    $connection = Get-NetTCPConnection -LocalPort $Port -ErrorAction SilentlyContinue
    return $connection -ne $null
}

# 1. Start PostgreSQL Service
Write-Host "Checking PostgreSQL service..." -ForegroundColor Cyan
$pgService = Get-Service -Name "postgresql*" -ErrorAction SilentlyContinue | Select-Object -First 1

if ($pgService) {
    if ($pgService.Status -ne 'Running') {
        Write-Host "   Starting PostgreSQL service: $($pgService.Name)..." -ForegroundColor Yellow
        try {
            Start-Service $pgService.Name
            Write-Host "   PostgreSQL service started successfully" -ForegroundColor Green
        }
        catch {
            Write-Host "   Failed to start PostgreSQL service: $($_.Exception.Message)" -ForegroundColor Red
            Write-Host "   Please start PostgreSQL manually or run as Administrator" -ForegroundColor Yellow
        }
    } else {
        Write-Host "   PostgreSQL service is already running" -ForegroundColor Green
    }
} else {
    Write-Host "   PostgreSQL service not found. Please ensure PostgreSQL is installed." -ForegroundColor Yellow
}

# Wait for PostgreSQL to be ready
Write-Host "   Waiting for PostgreSQL to be ready..." -ForegroundColor Yellow
Start-Sleep -Seconds 3

# 2. Set Environment Variables
Write-Host "Setting environment variables..." -ForegroundColor Cyan
$env:PGHOST = "localhost"
$env:PGPORT = "5432"
$env:PGUSER = "vueshop"
$env:PGPASSWORD = "vueshop123"
$env:PGDATABASE = "vueshop"
Write-Host "   Environment variables set" -ForegroundColor Green

# 3. Check if ports are available
if (Test-Port -Port 3000) {
    Write-Host "Port 3000 is already in use. Frontend might already be running." -ForegroundColor Yellow
}

if (Test-Port -Port 5000) {
    Write-Host "Port 5000 is already in use. Backend might already be running." -ForegroundColor Yellow
}

# 4. Start Backend (Go)
Write-Host "Starting Backend (Go) on http://localhost:5000..." -ForegroundColor Cyan
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PSScriptRoot\backend-go'; `$env:PGHOST='localhost'; `$env:PGPORT='5432'; `$env:PGUSER='vueshop'; `$env:PGPASSWORD='vueshop123'; `$env:PGDATABASE='vueshop'; go run ." -WindowStyle Normal

Write-Host "Waiting 5 seconds for backend to initialize..." -ForegroundColor Yellow
Start-Sleep -Seconds 5

# 5. Start Frontend (Vue.js)
Write-Host "Starting Frontend (Vue.js) on http://localhost:3000..." -ForegroundColor Cyan
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PSScriptRoot\frontend-vue'; npm run dev" -WindowStyle Normal

Write-Host ""
Write-Host "Full stack application is starting..." -ForegroundColor Green
Write-Host ""
Write-Host "Access Points:" -ForegroundColor White
Write-Host "   Frontend:    http://localhost:3000" -ForegroundColor Cyan
Write-Host "   Backend API: http://localhost:5000" -ForegroundColor Cyan
Write-Host "   Health:      http://localhost:5000/health" -ForegroundColor Cyan
Write-Host "   pgAdmin:     http://localhost:5050" -ForegroundColor Cyan
Write-Host ""
Write-Host "Admin Login:" -ForegroundColor White
Write-Host "   Username: admin" -ForegroundColor Yellow
Write-Host "   Password: admin123" -ForegroundColor Yellow
Write-Host ""
Write-Host "Press any key to exit this script (applications will continue running)..."
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")