package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const POSTGRES_PORT = 5432

func newPostgresDialector() (gorm.Dialector, error) {
	connectStr := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v",
		os.Getenv("POSTGRES_HOST"),
		POSTGRES_PORT,
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PASSWORD"),
	)

	sslmode := os.Getenv("DB_SSLMODE")
	if len(sslmode) > 0 {
		connectStr += fmt.Sprintf(" sslmode=%v", sslmode)
	}

	return postgres.Open(connectStr), nil
}
