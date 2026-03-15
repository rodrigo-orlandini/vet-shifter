# Testing

This document describes how to run and write tests in the Vet Shifter project.

## Overview

- **Unit tests** – Use cases, entities, value objects; no database. Use stubs from `backend/test/unit`.
- **Integration tests** – HTTP endpoints against a real Postgres DB; one shared DB, cleared before each test. Run with **one package at a time** (`-p 1`) and **no test-level parallelism** so tests do not truncate each other’s data or hit duplicate keys.
- **Architecture tests** – `arch-go` plus custom rules in `backend/test/architecture` (naming, layering).

All commands below assume you are in the **`backend/`** directory unless stated otherwise.

### Running tests on Windows

On Windows we use the **PowerShell scripts** in `backend/` to run each test suite. No Make or Bash required:

| Script                  | What it runs        |
|-------------------------|---------------------|
| `.\test-unit.ps1`       | Unit tests          |
| `.\test-integration.ps1`| Integration tests   |
| `.\test-architecture.ps1` | Architecture tests |

Open PowerShell, `cd backend`, then run for example `.\test-unit.ps1`. Integration tests load env from `.env.test` automatically; ensure the test Postgres is up (`docker compose -f docker-compose.test.yml up -d`) before running `.\test-integration.ps1`.

---

### Environment (`.env.test`)

Integration tests load **`backend/.env.test`** automatically when they run. The file is committed with default values for the local test DB:

- `DATABASE_URL` – Postgres on port 5433, database `vet_shifter_test`
- `JWT_SECRET` – at least 32 characters (required for auth tests)

You don’t need to set these in the shell. To use another DB or secret, edit `.env.test` or set the env vars before running (they override the file).
