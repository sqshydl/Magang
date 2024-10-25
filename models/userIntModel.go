package models

import (
	"time"
)

type UserIntModel struct {
	UserID         string    `json:"user_id" gorm:"primaryKey"` // Use string as primary key
	Username       string    `json:"username"`
	Password       string    `json:"password"`
	Name           string    `json:"name"`
	Phone          int       `json:"phone"`
	NIK            int       `json:"NIK"`
	Role           bool      `json:"role"`
	BankAcc        int       `json:"bank_account"`
	BankAccName    string    `json:"bank_account_name"`
	Status         bool      `json:"status"`
	CreatedAt      time.Time `json:"created_date"`
	LastActiveTime time.Time `json:"last_active"`
	LastActiveIP   string    `json:"last_active_ip"`
}
