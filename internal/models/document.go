package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Email     string    `gorm:"uniqueIndex;not null" json:"email"`
    Name      string    `gorm:"not null" json:"name"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Document struct {
    ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    Title     string    `gorm:"not null" json:"title"`
    Content   string    `gorm:"type:text" json:"content"`
    OwnerID   uuid.UUID `gorm:"type:uuid;not null" json:"owner_id"`
    Version   int       `gorm:"not null;default:1" json:"version"`
    IsPublic  bool      `gorm:"default:false" json:"is_public"`
    CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

	Owner         User                  `gorm:"foreignKey:OwnerID" json:"owner"`
    Collaborators []DocumentCollaborator `gorm:"foreignKey:DocumentID" json:"collaborators"`
    Versions      []DocumentVersion      `gorm:"foreignKey:DocumentID" json:"versions"`
}

type DocumentCollaborator struct {
    ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    DocumentID uuid.UUID `gorm:"type:uuid;not null" json:"document_id"`
    UserID     uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
    Permission string    `gorm:"not null" json:"permission"` // read, write, admin
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
    UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`

    User     User     `gorm:"foreignKey:UserID" json:"user"`
    Document Document `gorm:"foreignKey:DocumentID" json:"document"`
}

type DocumentVersion struct {
    ID         uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
    DocumentID uuid.UUID `gorm:"type:uuid;not null" json:"document_id"`
    Content    string    `gorm:"type:text;not null" json:"content"`
    Version    int       `gorm:"not null" json:"version"`
    CreatedBy  uuid.UUID `gorm:"type:uuid;not null" json:"created_by"`
    CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`

    Document Document `gorm:"foreignKey:DocumentID" json:"document"`
    User     User     `gorm:"foreignKey:CreatedBy" json:"user"`
}