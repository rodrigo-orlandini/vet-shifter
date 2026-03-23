$ErrorActionPreference = "Stop"

$root = Split-Path -Parent $MyInvocation.MyCommand.Path

Write-Host "==> Generating Swagger (backend)"
Set-Location (Join-Path $root "backend")

$gopath = & go env GOPATH
$goBin = Join-Path $gopath "bin"
$env:Path = "$goBin;$env:Path"

& go install github.com/swaggo/swag/cmd/swag@latest

& swag init -g main.go -d cmd/api,internal/auth/infrastructure/controllers,internal/companies/infrastructure/controllers,internal/veterinaries/infrastructure/controllers --parseDependency -o cmd/api/docs

Write-Host "==> Updating Orval client (frontend)"
$frontend = Join-Path $root "frontend"
Set-Location $frontend

if (-not (Test-Path (Join-Path $frontend "node_modules"))) {
  Write-Host "node_modules not found, running npm install..."
  cmd /c "npm.cmd install"
}

cmd /c "npm.cmd run gen:api"

Write-Host "==> Done. Returning to repo root"
Set-Location $root

