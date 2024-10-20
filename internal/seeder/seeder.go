package seeder

import (
	"fmt"
	"log"

	"github.com/Abhi-Bohora/multi-edit-api/internal/models"
	"gorm.io/gorm"
)

type Seeder struct {
	db *gorm.DB
}

func NewSeeder(db *gorm.DB) *Seeder {
	return &Seeder{db: db}
}

func (s *Seeder) SeedAll() error {
	log.Println("Starting database seeding...")

	if err := s.ClearData(); err != nil {
		return fmt.Errorf("error clearing data: %w", err)
	}

	users, err := s.SeedUsers()
	if err != nil {
		return fmt.Errorf("error seeding users: %w", err)
	}

	documents, err := s.SeedDocuments(users)
	if err != nil {
		return fmt.Errorf("error seeding documents: %w", err)
	}

	if err := s.SeedCollaborators(users, documents); err != nil {
		return fmt.Errorf("error seeding collaborators: %w", err)
	}

	if err := s.SeedDocumentVersions(documents, users); err != nil {
		return fmt.Errorf("error seeding document versions: %w", err)
	}

	log.Println("Database seeding completed successfully")
	return nil
}

func (s *Seeder) ClearData() error {
	log.Println("Clearing existing data...")

	// Disable foreign key checks during deletion
	if err := s.db.Exec("SET CONSTRAINTS ALL DEFERRED").Error; err != nil {
		return fmt.Errorf("error deferring constraints: %w", err)
	}

	if err := s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.DocumentVersion{}).Error; err != nil {
		return fmt.Errorf("error clearing document versions: %w", err)
	}

	if err := s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.DocumentCollaborator{}).Error; err != nil {
		return fmt.Errorf("error clearing document collaborators: %w", err)
	}

	if err := s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.Document{}).Error; err != nil {
		return fmt.Errorf("error clearing documents: %w", err)
	}

	if err := s.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&models.User{}).Error; err != nil {
		return fmt.Errorf("error clearing users: %w", err)
	}

	// Re-enable foreign key checks
	if err := s.db.Exec("SET CONSTRAINTS ALL IMMEDIATE").Error; err != nil {
		return fmt.Errorf("error re-enabling constraints: %w", err)
	}

	log.Println("Successfully cleared all existing data")
	return nil
}

func (s *Seeder) SeedUsers() ([]models.User, error) {
	log.Println("Seeding users...")
	users := []models.User{
		{
			Email: "john@example.com",
			Name:  "John Doe",
		},
		{
			Email: "jane@example.com",
			Name:  "Jane Smith",
		},
		{
			Email: "bob@example.com",
			Name:  "Bob Wilson",
		},
	}

	for i := range users {
		if err := s.db.Create(&users[i]).Error; err != nil {
			return nil, fmt.Errorf("error creating user %s: %w", users[i].Email, err)
		}
	}

	log.Printf("Successfully seeded %d users", len(users))
	return users, nil
}

func (s *Seeder) SeedDocuments(users []models.User) ([]models.Document, error) {
	log.Println("Seeding documents...")
	documents := []models.Document{
		{
			Title:    "Getting Started Guide",
			Content:  "# Welcome to our platform\n\nThis is a guide to help you get started...",
			OwnerID:  users[0].ID,
			Version:  1,
			IsPublic: true,
		},
		{
			Title:    "Project Proposal",
			Content:  "## Project Overview\n\nThis project aims to...",
			OwnerID:  users[1].ID,
			Version:  1,
			IsPublic: false,
		},
		{
			Title:    "Meeting Notes",
			Content:  "### Team Meeting - 2024\n\n1. Discussion points...",
			OwnerID:  users[0].ID,
			Version:  1,
			IsPublic: false,
		},
	}

	for i := range documents {
		if err := s.db.Create(&documents[i]).Error; err != nil {
			return nil, fmt.Errorf("error creating document %s: %w", documents[i].Title, err)
		}
	}

	log.Printf("Successfully seeded %d documents", len(documents))
	return documents, nil
}

func (s *Seeder) SeedCollaborators(users []models.User, documents []models.Document) error {
	log.Println("Seeding document collaborators...")
	collaborators := []models.DocumentCollaborator{
		{
			DocumentID: documents[0].ID,
			UserID:    users[1].ID,
			Permission: "write",
		},
		{
			DocumentID: documents[0].ID,
			UserID:    users[2].ID,
			Permission: "read",
		},
		{
			DocumentID: documents[1].ID,
			UserID:    users[2].ID,
			Permission: "admin",
		},
	}

	for _, collaborator := range collaborators {
		if err := s.db.Create(&collaborator).Error; err != nil {
			return fmt.Errorf("error creating collaborator: %w", err)
		}
	}

	log.Printf("Successfully seeded %d collaborators", len(collaborators))
	return nil
}

func (s *Seeder) SeedDocumentVersions(documents []models.Document, users []models.User) error {
	log.Println("Seeding document versions...")
	var totalVersions int

	for _, doc := range documents {
		versions := []models.DocumentVersion{
			{
				DocumentID: doc.ID,
				Content:   doc.Content + "\nFirst revision",
				Version:   1,
				CreatedBy: doc.OwnerID,
			},
			{
				DocumentID: doc.ID,
				Content:   doc.Content + "\nSecond revision with more content...",
				Version:   2,
				CreatedBy: doc.OwnerID,
			},
		}

		for _, version := range versions {
			if err := s.db.Create(&version).Error; err != nil {
				return fmt.Errorf("error creating version for document %s: %w", doc.Title, err)
			}
			totalVersions++
		}
	}

	log.Printf("Successfully seeded %d document versions", totalVersions)
	return nil
}