package database

import (
	"database/sql"
	"fmt"

	"rodrigoorlandini/vet-shifter/internal/_shared/database/queries"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var postgresConnection *sql.DB

func ResetConnectionForTest() {
	postgresConnection = nil
}

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
	databaseUrl := utils.GetDatabaseURL()
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

func TruncateAllTables(db *sql.DB) error {
	_, err := db.Exec(`
		TRUNCATE TABLE password_reset_tokens, company_owners, companies, shift_veterinaries
		RESTART IDENTITY CASCADE
	`)
	return err
}
