package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	TransactionID   int          `gorm:"primaryKey;autoIncrement" json:"transaction_id"` // Unique transaction ID
	UserID          uuid.UUID    `gorm:"type:uuid;not null" json:"user_id"`              // Foreign key referencing UserIntModel
	User            UserIntModel `gorm:"foreignKey:UserID;references:UserID"`            // User relationship
	TransactionDate time.Time    `json:"transaction_date"`                               // Date of the transaction
	Price           float64      `json:"price"`                                          // Transaction amount
	TypePaymentID   int          `gorm:"column:type_payment_id" json:"type_payment_id"`  // Foreign key for payment type
	PaymentStatus   string       `gorm:"type:varchar(50)" json:"payment_status"`         // Payment status
	CreatedDate     time.Time    `gorm:"autoCreateTime" json:"created_date"`             // Creation timestamp
}
