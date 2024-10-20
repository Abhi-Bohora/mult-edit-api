package main

import (
	"log"

	"github.com/Abhi-Bohora/multi-edit-api/internal/config"
	"github.com/Abhi-Bohora/multi-edit-api/internal/database"
)

func main(){
	dbConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load database config: %w", err)
	}

	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		log.Fatal("Failed to connect to a database: %w", err)
	}

	//if connection is successfull than we will automigrate the models
	log.Println("Database connection successfull")
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
}