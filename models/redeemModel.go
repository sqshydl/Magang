package models

import (
	"time"
)

type Redeem struct {
	RedeemID      int       `gorm:"primaryKey;autoIncrement" json:"redeem_id"` // Unique identifier
	UserID        int       `json:"user_id"`                                   // Foreign key to UserIntModel (or related user model)
	NominalRedeem int       `json:"nominal_redeem"`                            // Amount to redeem
	Status        string    `gorm:"type:varchar(10)" json:"status"`            // Status, can be "approve" or "reject"
	CreatedDate   time.Time `gorm:"autoCreateTime" json:"created_date"`        // Creation timestamp
}
