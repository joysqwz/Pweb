package database

import (
	"log"
	"pweb/backend/models/db_models"
)

func (db *Database) autoMigrate() {
	tx := db.Exec(`
			DO
			$$
			BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'mode_enum') THEN
    			create type mode_enum AS ENUM ('', 'cooperative', 'competitive', 'neutral');
  			END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'difficulty_enum') THEN
    			create type difficulty_enum AS ENUM ('', 'easy', 'medium', 'hard');
  			END IF;
			END
			$$;
	`)
	if tx.Error != nil {
		log.Fatalf("Couldn't create enum types: %v", tx.Error)
	}

	err := db.AutoMigrate(
		&db_models.User{},
		&db_models.Sheet{},
	)
	if err != nil {
		log.Fatalf("Couldn't auto migrate: %v", err)
	}

	if tx.Error != nil {
		log.Fatalf("Couldn't exec migrate: %v", tx.Error)
	}
}
