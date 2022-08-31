package migrations

import "testing"

func TestMigrate(t *testing.T) {
	if testing.Short() {
		t.Skip("debugger only")
	}

	Migrate()
}

func TestRollback(t *testing.T) {
	if testing.Short() {
		t.Skip("debugger only")
	}

	Rollback()
}
