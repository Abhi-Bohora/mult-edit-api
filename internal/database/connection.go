package database

import (
	"fmt"
	"log"

	"github.com/Abhi-Bohora/multi-edit-api/internal/config"
	"github.com/Abhi-Bohora/multi-edit-api/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(config *config.DBConfig) (*Database, error) {
	db, err := gorm.Open(postgres.Open(config.GetDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}

	return &Database{db},nil
}

func (db *Database) AutoMigrate() error {
	log.Println("Running database migrations...")

	db.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	
	err := db.DB.AutoMigrate(
		&models.User{},
		&models.Document{},
		&models.DocumentCollaborator{},
		&models.DocumentVersion{},
	)

	if err != nil {
		return fmt.Errorf("error running migrations: %w", err)
	}

	log.Println("Database migrations completed successfully")
	return nil

}