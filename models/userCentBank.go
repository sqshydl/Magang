package models

import (
	"time"

	"github.com/google/uuid"
)

// User represents the structure for the Users table
type UserCentBank struct {
	UserID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"` // Unique user identifier
	Name           string    `gorm:"type:varchar(255)"`                               // User's name
	Username       string    `gorm:"unique;type:varchar(255)"`                        // User's unique username
	Password       string    `gorm:"type:varchar(255)"`                               // User's password (hashed)
	LastActivityIP string    `gorm:"type:varchar(255)"`                               // IP address of last activity
	LastActiveTime time.Time `gorm:"autoUpdateTime"`                                  // Timestamp of last activity
	CreatedDate    time.Time `gorm:"autoCreateTime"`                                  // Timestamp of when the user was created
}
