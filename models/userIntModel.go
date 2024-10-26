package models

import (
	"time"

	"github.com/google/uuid"
)

type UserIntModel struct {
	UserID         uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"user_id"` // Unique user identifier
	Username       string    `gorm:"unique;type:varchar(255)" json:"username"`                       // User's unique username
	Password       string    `gorm:"type:varchar(255)" json:"password"`                              // Hashed password, excluded from JSON
	Name           string    `gorm:"type:varchar(255)" json:"name"`                                  // User's full name
	Phone          int       `gorm:"type:bigint" json:"phone"`                                       // User's phone number
	NIK            int       `gorm:"type:bigint" json:"nik"`                                         // National ID number
	Role           bool      `gorm:"default:true" json:"role"`                                       // True = Validator, False = Non-Validator
	BankAcc        int       `gorm:"type:bigint" json:"bank_acc"`                                    // Bank account number
	BankAccName    string    `gorm:"type:varchar(255)" json:"bank_acc_name"`                         // Name on the bank account
	Status         bool      `gorm:"default:true" json:"status"`                                     // True = Active, False = Inactive
	CreatedAt      time.Time `gorm:"autoCreateTime" json:"created_at"`                               // Timestamp when the user was created
	LastActiveIP   string    `gorm:"type:varchar(45)" json:"last_active_ip"`                         // IP address of the user's last activity
	LastActiveTime time.Time `gorm:"autoUpdateTime" json:"last_active_time"`                         // Timestamp of the user's last activity
}
