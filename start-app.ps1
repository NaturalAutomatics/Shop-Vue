# Vue Shop - Start Both Frontend and Backend
# This script starts both the Vue.js frontend and Go backend

Write-Host "🚀 Starting Vue Shop Application..." -ForegroundColor Green
Write-Host ""

# Function to check if a port is in use
function Test-Port {
    param([int]$Port)
    $connection = Get-NetTCPConnection -LocalPort $Port -ErrorAction SilentlyContinue
    return $connection -ne $null
}

# Check if ports are already in use
if (Test-Port -Port 3000) {
    Write-Host "⚠️  Port 3000 is already in use. Frontend might already be running." -ForegroundColor Yellow
}

if (Test-Port -Port 5000) {
    Write-Host "⚠️  Port 5000 is already in use. Backend might already be running." -ForegroundColor Yellow
}

Write-Host "📦 Starting Frontend (Vue.js) on http://localhost:3000..." -ForegroundColor Cyan
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PSScriptRoot'; npm run dev" -WindowStyle Normal

Write-Host "⏳ Waiting 3 seconds for frontend to initialize..." -ForegroundColor Yellow
Start-Sleep -Seconds 3

Write-Host "🔧 Starting Backend (Go) on http://localhost:5000..." -ForegroundColor Cyan
Start-Process powershell -ArgumentList "-NoExit", "-Command", "cd '$PSScriptRoot\backend-go'; go run ." -WindowStyle Normal

Write-Host ""
Write-Host "✅ Both applications are starting..." -ForegroundColor Green
Write-Host ""
Write-Host "🌐 Frontend: http://localhost:3000" -ForegroundColor White
Write-Host "🔌 Backend API: http://localhost:5000" -ForegroundColor White
Write-Host "🏥 Health Check: http://localhost:5000/health" -ForegroundColor White
Write-Host ""
Write-Host "Press any key to exit this script (applications will continue running)..."
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
