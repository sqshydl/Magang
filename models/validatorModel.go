package models

import (
	"time"

	"github.com/google/uuid"
)

type Validator struct {
	ValidatorsID         string    `gorm:"primaryKey;type:varchar(255);default:gen_random_uuid()" json:"validators_id"` // Randomly generated ID
	UserIntermediariesID uuid.UUID `gorm:"type:uuid;not null" json:"user_intermediaries_id"`                            // Foreign key to UserIntModel
	CreatedDate          time.Time `gorm:"autoCreateTime" json:"created_date"`                                          // Timestamp when the validator record is created

	// Foreign key relationship
	UserIntermediary UserIntModel `gorm:"foreignKey:UserIntermediariesID;references:UserID" json:"user_intermediary"` // Establishes relation to UserIntModel
}
