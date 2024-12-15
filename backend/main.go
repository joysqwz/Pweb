package main

import (
	"log"
	"os"
	"pweb/backend/external/database"
	"pweb/backend/external/httpserver"

	"github.com/joho/godotenv"
	"github.com/natefinch/lumberjack"
)

func main() {
	if godotenv.Load() != nil {
		log.Printf("Using environment variables from container environment")
	} else {
		log.Printf("Using environment variables from .env file")
	}

	if os.Getenv("ENVIRONMENT") != "server" {
		log.Printf("No dev environment was set so logging to out.log file")
		log.SetOutput(&lumberjack.Logger{
			Filename:   "out.log",
			MaxSize:    250, // megabytes
			MaxBackups: 2,
			MaxAge:     14,   //days
			Compress:   true, // disabled by default
		})
	}

	db := database.InitDB()
	httpserver.InitAndStart(db)

}
