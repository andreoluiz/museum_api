package database

import (
	"log"
	"os"
)

func RunMigrations() {
	if DB == nil {
		log.Fatalf("Database connection is nil")
	}
	migrationFiles := []string{
		"database/migrations/0001_create_managers_table.sql",
		"database/migrations/0002_create_museums_table.sql",
		"database/migrations/0003_create_artworks_table.sql",
	}

	for _, file := range migrationFiles {
		if err := applyMigration(file); err != nil {
			log.Fatalf("Failed to execute migration file %s: %v", file, err)
		}
	}
}

func applyMigration(file string) error {
	migration, err := os.ReadFile(file)
	if err != nil {
		return err
	}
	return DB.Exec(string(migration)).Error
}
