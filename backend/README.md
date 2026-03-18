# Backend (Go)

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
