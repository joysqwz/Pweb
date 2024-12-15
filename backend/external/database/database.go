package database

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database struct {
	*gorm.DB
	prefix string
}

func InitDB() *Database {
	postgresDialector, err := newPostgresDialector()
	if err != nil {
		log.Fatalf("Couldn't init dialector: %v", err)
	}

	prefix := os.Getenv("POSTGRES_PREFIX")

	var gormDB *gorm.DB
	for i := 0; i < 12; i++ {
		// NB: POSTGRES_PREFIX не работает как надо, т.к есть JOIN запросы с явными именами таблиц
		// 		приходится там явно соблюдать нейминг
		gormDB, err = gorm.Open(postgresDialector, &gorm.Config{
			Logger: logger.Default.LogMode(logger.Silent),
			NamingStrategy: schema.NamingStrategy{
				TablePrefix: prefix,
			},
		})
		if err == nil {
			break
		}
		log.Printf("Init database error: %v. Will try again in 10 seconds...", err)
		time.Sleep(10 * time.Second)
	}

	if err != nil {
		log.Fatalf("Init database error: %v", err)
	}

	db := &Database{gormDB, prefix}

	db.autoMigrate()

	return db
}
