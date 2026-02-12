package database

import (
	"database/sql"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const migrationsTable = "schema_migrations"
const migrationsDir = "internal/_shared/database/migrations"

type Migration struct {
	Version string
	File    string
	SQL     string
}

func RunMigrations(db *sql.DB) error {
	if err := createMigrationsTable(db); err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	migrations, err := getMigrationFiles()
	if err != nil {
		return fmt.Errorf("failed to read migration files: %w", err)
	}

	applied, err := getAppliedMigrations(db)
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	pending := getPendingMigrations(migrations, applied)

	if len(pending) == 0 {
		fmt.Println("No pending migrations")
		return nil
	}

	fmt.Printf("Found %d pending migration(s)\n", len(pending))

	for _, migration := range pending {
		if err := applyMigration(db, migration); err != nil {
			return fmt.Errorf("failed to apply migration %s: %w", migration.Version, err)
		}
		fmt.Printf("Applied migration: %s\n", migration.Version)
	}

	return nil
}

func createMigrationsTable(db *sql.DB) error {
	query := fmt.Sprintf(`
		CREATE TABLE IF NOT EXISTS %s (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		)
	`, migrationsTable)

	_, err := db.Exec(query)
	return err
}

func getMigrationFiles() ([]Migration, error) {
	var migrations []Migration

	paths := []string{
		migrationsDir,
		filepath.Join("..", migrationsDir),
		filepath.Join("..", "..", migrationsDir),
		filepath.Join("backend", migrationsDir),
	}

	var migrationsPath string
	for _, path := range paths {
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			migrationsPath = path
			break
		}
	}

	if migrationsPath == "" {
		return nil, fmt.Errorf("migrations directory not found. Tried paths: %v", paths)
	}

	err := filepath.WalkDir(migrationsPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(path, ".sql") {
			return nil
		}

		if strings.Contains(path, "atlas.sum") {
			return nil
		}

		filename := filepath.Base(path)
		parts := strings.Split(filename, "_")
		if len(parts) == 0 {
			return nil
		}
		version := parts[0]

		sql, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", path, err)
		}

		migrations = append(migrations, Migration{
			Version: version,
			File:    filename,
			SQL:     string(sql),
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

func getAppliedMigrations(db *sql.DB) (map[string]bool, error) {
	query := fmt.Sprintf("SELECT version FROM %s", migrationsTable)
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	applied := make(map[string]bool)
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, err
		}
		applied[version] = true
	}

	return applied, rows.Err()
}

func getPendingMigrations(migrations []Migration, applied map[string]bool) []Migration {
	var pending []Migration
	for _, migration := range migrations {
		if !applied[migration.Version] {
			pending = append(pending, migration)
		}
	}
	return pending
}

func applyMigration(db *sql.DB, migration Migration) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.Exec(migration.SQL); err != nil {
		return fmt.Errorf("failed to execute migration SQL: %w", err)
	}

	query := fmt.Sprintf("INSERT INTO %s (version) VALUES ($1)", migrationsTable)
	if _, err := tx.Exec(query, migration.Version); err != nil {
		return fmt.Errorf("failed to record migration: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
