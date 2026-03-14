package integration

import (
	"os"
	"sync"
	"testing"

	"rodrigoorlandini/vet-shifter/internal/_shared/database"
	"rodrigoorlandini/vet-shifter/internal/_shared/utils"
)

var prepareDBMu sync.Mutex

func PrepareDB(t *testing.T) {
	t.Helper()

	utils.LoadTestEnvironment()
	if os.Getenv("DATABASE_URL") == "" {
		t.Skip("DATABASE_URL not set; skipping integration test (add backend/.env.test or set env)")
	}

	prepareDBMu.Lock()
	defer prepareDBMu.Unlock()

	utils.LoadEnvironment()
	database.ResetConnectionForTest()

	conn := database.GetPostgresConnection()
	if err := database.RunMigrations(conn); err != nil {
		t.Fatalf("run migrations: %v", err)
	}

	if err := database.TruncateAllTables(conn); err != nil {
		t.Fatalf("truncate tables: %v", err)
	}
}
