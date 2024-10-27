package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID           uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"user_id"` // Unique user identifier
	Username         string    `gorm:"unique;type:varchar(255)" json:"username"`                       // Unique username
	Password         string    `gorm:"type:varchar(255)" json:"password"`                              // Password (hashed)
	Name             string    `gorm:"type:varchar(255)" json:"name"`                                  // Full name
	Phone            int       `gorm:"type:bigint" json:"phone"`                                       // Phone number
	NIK              int       `gorm:"type:bigint" json:"nik"`                                         // National ID number
	Role             string    `gorm:"type:varchar(50)" json:"role"`                                   // Role (merchant/user)
	BankAccount      int       `gorm:"type:bigint" json:"bank_account"`                                // Bank account number
	BankAccountName  string    `gorm:"type:varchar(255)" json:"bank_account_name"`                     // Bank account name
	IsIntermediaries int       `gorm:"type:int" json:"is_intermediaries"`                              // Is intermediaries (1 = yes, 0 = no)
	CreatedDate      time.Time `gorm:"autoCreateTime" json:"created_date"`                             // Timestamp when the user was created
	LogActivityIP    string    `gorm:"type:varchar(45)" json:"log_activity_ip"`                        // IP address for log activity
	LogActivityTime  time.Time `gorm:"type:timestamp" json:"log_activity_time"`                        // Timestamp of log activity
}
