package database

import (
	"database/sql"
	"embed"
	"fmt"
	"io/fs"
	"log/slog"
	"os"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var migrationsFS embed.FS

const migrationsTable = "schema_migrations"

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

	slog.Info("Found %d pending migration(s)\n", "length", len(pending))

	for _, migration := range pending {
		if err := applyMigration(db, migration); err != nil {
			return fmt.Errorf("failed to apply migration %s: %w", migration.Version, err)
		}

		slog.Info("Applied migration: \n", migration.Version)
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
	migrations, err := getMigrationFilesFromEmbed()
	if err == nil && len(migrations) > 0 {
		return migrations, nil
	}

	return getMigrationFilesFromFS()
}

func getMigrationFilesFromEmbed() ([]Migration, error) {
	var migrations []Migration

	entries, err := fs.ReadDir(migrationsFS, "migrations")
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		filename := entry.Name()
		parts := strings.Split(filename, "_")
		if len(parts) == 0 {
			continue
		}

		version := parts[0]
		sqlContent, err := fs.ReadFile(migrationsFS, "migrations/"+filename)
		if err != nil {
			return nil, err
		}

		migrations = append(migrations, Migration{
			Version: version,
			File:    filename,
			SQL:     string(sqlContent),
		})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

func getMigrationFilesFromFS() ([]Migration, error) {
	dir := findMigrationsDir()
	if dir == "" {
		return nil, fmt.Errorf("migrations directory not found (tried embed and filesystem)")
	}

	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("read migrations dir: %w", err)
	}

	var migrations []Migration
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".sql") {
			continue
		}

		filename := entry.Name()
		parts := strings.Split(filename, "_")
		if len(parts) == 0 {
			continue
		}

		version := parts[0]
		sqlContent, err := os.ReadFile(filepath.Join(dir, filename))
		if err != nil {
			return nil, fmt.Errorf("read migration %s: %w", filename, err)
		}

		migrations = append(migrations, Migration{
			Version: version,
			File:    filename,
			SQL:     string(sqlContent),
		})
	}

	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

func findMigrationsDir() string {
	const relPath = "internal/_shared/database/migrations"

	wd, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		path := filepath.Join(wd, relPath)
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			return path
		}

		parent := filepath.Dir(wd)
		if parent == wd {
			break
		}
		wd = parent
	}

	for _, wd = range []string{".", "..", "../..", "../../..", "../../../..", "../../../../.."} {
		path := filepath.Join(wd, relPath)
		if info, err := os.Stat(path); err == nil && info.IsDir() {
			return path
		}
	}

	return ""
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
