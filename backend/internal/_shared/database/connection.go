package database

import (
	"database/sql"
	"fmt"
	"os"

	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var postgresConnection *sql.DB

func GetPostgresConnection() *sql.DB {
	if postgresConnection != nil {
		return postgresConnection
	}

	conn, err := newConnection()
	if err != nil {
		panic(fmt.Sprintf("[POSTGRES_CONNECTION_FAILED]: %v", err))
	}

	postgresConnection = conn
	return conn
}

func newConnection() (*sql.DB, error) {
	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == "" {
		panic("[DATABASE_URL_NOT_FOUND]: DATABASE_URL is not set on Environment Variables")
	}

	conn, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}

func NewQueries(conn *sql.DB) *queries.Queries {
	return queries.New(conn)
}
