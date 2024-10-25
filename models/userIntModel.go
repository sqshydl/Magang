package models

import (
	"time"

	"github.com/google/uuid"
)

type UserIntModel struct {
	UserID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Username       string    `gorm:"unique"`
	Password       string
	Name           string
	Phone          int
	NIK            int
	Role           bool
	BankAcc        int
	BankAccName    string
	Status         bool
	CreatedAt      time.Time
	LastActiveIP   string
	LastActiveTime time.Time
}
