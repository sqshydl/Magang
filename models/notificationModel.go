package models

import "time"

type Notification struct {
	NotificationID     int       `gorm:"primaryKey;autoIncrement" json:"notification_id"`
	NotificationTypeID int       `json:"notification_type_id"`
	Message            string    `gorm:"type:text" json:"message"`
	Date               time.Time `json:"date"`
	Status             bool      `gorm:"default:true" json:"status"` // True = !Read, False = Read
	CreatedDate        time.Time `gorm:"autoCreateTime" json:"created_date"`
}
