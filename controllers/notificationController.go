package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/models"
	"gorm.io/gorm"
)

func NotificationCreate(c *gin.Context) {
	// Define the structure for receiving JSON input
	var NotificationInput struct {
		NotificationTypeID int       `json:"notification_type_id"`
		Message            string    `json:"message"`
		Date               time.Time `json:"date"`
		Status             bool      `json:"status"` // True = !Read, False = Read
	}

	// Parse JSON input into NotificationInput struct
	if err := c.BindJSON(&NotificationInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Map input data to the Notification model
	notification := models.Notification{
		NotificationTypeID: NotificationInput.NotificationTypeID,
		Message:            NotificationInput.Message,
		Date:               NotificationInput.Date,
		Status:             NotificationInput.Status,
		CreatedDate:        time.Now(),
	}

	// Insert the new notification record into the database
	result := initializers.DB.Create(&notification)

	// Handle any errors that may occur during database insertion
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond with the created notification data
	c.JSON(http.StatusOK, gin.H{
		"message":         "Notification created successfully",
		"notification_id": notification.NotificationID,
		"notification":    notification,
	})
}

func NotificationIndex(c *gin.Context) {
	var allNotifications []models.Notification
	initializers.DB.Find(&allNotifications)
	result := initializers.DB.Find(&allNotifications)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"post": allNotifications,
	})
}

func NotificationGet(c *gin.Context) {
	id := c.Param("notification_id")

	var getNotification models.Notification
	result := initializers.DB.Where("notification_id = ?", id).First(&getNotification)

	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"Notification": getNotification,
	})
}

func NotificationUpdate(c *gin.Context) {
	// Get the notification ID from the URL parameter
	id := c.Param("notification_id")

	// Define a temporary struct to hold the fields that can be updated
	var updateData struct {
		NotificationTypeID int       `json:"notification_type_id"`
		Message            string    `json:"message"`
		Date               time.Time `json:"date"`
		Status             bool      `json:"status"`
	}

	// Bind JSON input to updateData struct
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Find the notification by ID
	var notification models.Notification
	result := initializers.DB.First(&notification, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "Notification not found"})
			return
		}
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// Update the fields in the notification model
	initializers.DB.Model(&notification).Updates(updateData)

	// Respond with success and the updated notification
	c.JSON(200, gin.H{"message": "Notification updated successfully", "notification": notification})
}

func NotificationDelete(c *gin.Context) {
	// Get the notification ID from the URL parameter
	id := c.Param("notification_id")

	// Find the notification by ID
	var notification models.Notification
	result := initializers.DB.First(&notification, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "Notification not found"})
			return
		}
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// Delete the notification
	initializers.DB.Delete(&notification)

	// Respond with success message
	c.JSON(200, gin.H{"message": "Notification deleted successfully"})
}
