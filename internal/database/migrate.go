package database

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MigrationService struct {
	databaseURL string
}

func NewMigrationService(databaseURL string) *MigrationService {
	return &MigrationService{
		databaseURL: databaseURL,
	}
}

// RunMigrations runs all pending migrations
func (ms *MigrationService) RunMigrations() error {
	m, err := migrate.New(
		"file://migrations",
		ms.databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil {
		if err == migrate.ErrNoChange {
			log.Println("✓ No new migrations to apply")
			return nil
		}
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	version, dirty, err := m.Version()
	if err != nil {
		log.Println("✓ Migrations applied successfully")
		return nil
	}

	log.Printf("✓ Migrations applied successfully (current version: %d, dirty: %v)", version, dirty)
	return nil
}

// RollbackMigration rolls back the last N migrations
func (ms *MigrationService) RollbackMigration(steps int) error {
	m, err := migrate.New(
		"file://migrations",
		ms.databaseURL,
	)
	if err != nil {
		return fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	if err := m.Steps(-steps); err != nil {
		return fmt.Errorf("failed to rollback migrations: %w", err)
	}

	log.Printf("✓ Rolled back %d migration(s)", steps)
	return nil
}

// GetVersion returns the current migration version
func (ms *MigrationService) GetVersion() (uint, bool, error) {
	m, err := migrate.New(
		"file://migrations",
		ms.databaseURL,
	)
	if err != nil {
		return 0, false, fmt.Errorf("failed to create migrate instance: %w", err)
	}
	defer m.Close()

	return m.Version()
}