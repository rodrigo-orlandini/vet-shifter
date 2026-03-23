# Backend (Go)

## Initial setup

1. Install **Go** (this module targets **Go 1.24**; see `go.mod`).
2. From the `backend` directory, download and tidy module dependencies:

```powershell
cd backend
go mod tidy
```

3. Copy environment template and adjust values (database, JWT, optional email):

```powershell
copy .env.example .env
```

See `.env.example` for `DATABASE_URL`, `JWT_SECRET`, `API_PORT`, and email-related variables.

## Running in development

1. Start **PostgreSQL** (recommended: Docker Compose in this folder):

```powershell
cd backend
docker compose up -d
```

Wait until the `postgres` service is healthy (`pg_isready`).

2. Ensure `.env` exists and `DATABASE_URL` matches your Compose settings (default in `.env.example` uses `localhost:5432` and database `vet_shifter`).

3. Run the API (migrations run automatically on startup):

```powershell
cd backend
go run ./cmd/api
```

The server listens on `API_PORT` (default **8000**). Swagger is generated under `cmd/api/docs/` when you run `swag init` (see root `gen-api.ps1` if you use it).

---

## SQLC (queries → generated Go code)

This project uses **SQLC** to generate type-safe Go code from SQL.

### Where SQLC is configured

- **Config**: `internal/_shared/database/sqlc/sqlc.yaml`
- **Schema input**: `internal/_shared/database/sqlc/schemas.sql`
- **Queries input**: `internal/_shared/database/sqlc/queries.sql`
- **Generated output**: `internal/_shared/database/queries/`

### Install SQLC (Windows / PowerShell)

```powershell
cd backend
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Make sure your Go bin is in PATH (common default is `%USERPROFILE%\go\bin`).

### Regenerate SQLC files

```powershell
cd backend
sqlc generate -f internal/_shared/database/sqlc/sqlc.yaml
```

### Add / change database queries

1) Edit `internal/_shared/database/sqlc/queries.sql`.
2) Add a query with an SQLC name header:

- `-- name: FindSomething :one`
- `-- name: ListSomething :many`
- `-- name: CreateSomething :exec`

3) Regenerate with `sqlc generate ...`.

## Migrations (schema changes)

Migrations are plain `.sql` files in:
- `internal/_shared/database/migrations/`

They are applied automatically on API startup via `database.RunMigrations(...)` (called from `cmd/api/main.go`).

### Add a new migration

1) Create a new file under `internal/_shared/database/migrations/` named like:

- `002_add_some_table.sql`

The **numeric prefix** is used to order migrations.

2) Put your SQL changes in the file.
3) Start the API (or restart it) to apply pending migrations.

### When adding/changing tables, enums, columns

SQLC does **not** read migrations directly here; it reads `schemas.sql`.

So after writing a migration, also update:
- `internal/_shared/database/sqlc/schemas.sql`

Then regenerate SQLC:

```powershell
cd backend
sqlc generate -f internal/_shared/database/sqlc/sqlc.yaml
```
