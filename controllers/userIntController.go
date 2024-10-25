package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/models"
	"gorm.io/gorm"
)

func UserIntCreate(c *gin.Context) {

	var createUserInt struct {
		Username    string `json:"username"`
		Password    string `json:"password"`
		Name        string `json:"name"`
		Phone       int    `json:"phone"`
		NIK         int    `json:"NIK"`
		Role        bool   `json:"role"`
		BankAcc     int    `json:"bank_account"`
		BankAccName string `json:"bank_account_name"`
		Status      bool   `json:"status"`
	}

	if err := c.BindJSON(&createUserInt); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	userInt := models.UserIntModel{
		UserID:         uuid.New().String(),
		Username:       createUserInt.Username,    // Use input data
		Password:       createUserInt.Password,    // Use input data
		Name:           createUserInt.Name,        // Use input data
		Phone:          createUserInt.Phone,       // Use input data
		NIK:            createUserInt.NIK,         // Use input data
		Role:           createUserInt.Role,        // Use input data
		BankAcc:        createUserInt.BankAcc,     // Use input data
		BankAccName:    createUserInt.BankAccName, // Use input data
		Status:         createUserInt.Status,      // Use input data
		CreatedAt:      time.Now(),                // Set created date to now
		LastActiveIP:   c.ClientIP(),              // Capture the client IP
		LastActiveTime: time.Now(),
	}
	result := initializers.DB.Create(&userInt)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()}) // Handle database error
		return

	}
	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user_id": userInt.UserID,
		"post":    userInt,
	})
}

func UserIntIndex(c *gin.Context) {
	var getAllUserInt []models.UserIntModel
	initializers.DB.Find(&getAllUserInt)
	result := initializers.DB.Find(&getAllUserInt)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"post": getAllUserInt,
	})
}

func UserIntGet(c *gin.Context) {
	username := c.Param("username")

	var getUserInt models.UserIntModel
	result := initializers.DB.Where("username = ?", username).First(&getUserInt)

	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"user": getUserInt,
	})
}
