package models

import (
	"time"

	"github.com/google/uuid"
)

// UUID: Universally Unique Identifier.
// Email: unique to ensure that no two users end up with the same email addresses.
// “: specify tags with backtick annotation on the fields that need modification.
type User struct {
	ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	// ID     uint   `gorm:"primary_key"`
	Name      string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
