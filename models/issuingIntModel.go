package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IssuingIntermediaries struct {
	IssuingIntermediariesID   string    `gorm:"primaryKey;type:varchar(36)" json:"issuing_intermediaries_id"` // Unique identifier
	UserIntermediariesID      uuid.UUID `gorm:"type:uuid" json:"user_intermediaries_id"`                      // Foreign key to UserIntModel
	IssuingIntermediariesName string    `gorm:"type:varchar(255)" json:"issuing_intermediaries_name"`         // Name of the issuing intermediary
	Point                     string    `gorm:"type:varchar(255)" json:"point"`                               // Point value as string
	Saldo                     string    `gorm:"type:varchar(255)" json:"saldo"`                               // Saldo as string
	Redeem                    int       `gorm:"type:int" json:"redeem"`                                       // Redeem count
	CreatedDate               time.Time `gorm:"autoCreateTime" json:"created_date"`                           // Creation timestamp
	LogActivityIP             string    `gorm:"type:varchar(45)" json:"log_activity_ip"`                      // IP address of log activity
	LogActivityTimestamp      time.Time `json:"log_activity_timestamp"`                                       // Timestamp of log activity
}

func (issuing *IssuingIntermediaries) BeforeCreate(tx *gorm.DB) (err error) {
	issuing.IssuingIntermediariesID = uuid.NewString()
	return
}
