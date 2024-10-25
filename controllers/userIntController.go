package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/squishydal/MAGANG/auth"
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
		Role        bool   `json:"role"` // True = Validator, False = Non-Validator
		BankAcc     int    `json:"bank_account"`
		BankAccName string `json:"bank_account_name"`
		Status      bool   `json:"status"` // True = Active, False = Inactive
	}

	if err := c.BindJSON(&createUserInt); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the username already exists
	var existingUser models.UserIntModel
	if err := initializers.DB.Where("username = ?", createUserInt.Username).First(&existingUser).Error; err == nil {
		c.JSON(409, gin.H{"error": "Username already exists"})
		return
	}

	hashedPassword, err := auth.HashPassword(createUserInt.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// Proceed with creating the user if the username is unique
	userInt := models.UserIntModel{
		Username:       createUserInt.Username,
		Password:       hashedPassword,
		Name:           createUserInt.Name,
		Phone:          createUserInt.Phone,
		NIK:            createUserInt.NIK,
		Role:           createUserInt.Role,
		BankAcc:        createUserInt.BankAcc,
		BankAccName:    createUserInt.BankAccName,
		Status:         createUserInt.Status,
		CreatedAt:      time.Now(),
		LastActiveIP:   c.ClientIP(),
		LastActiveTime: time.Now(),
	}
	result := initializers.DB.Create(&userInt)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user_id": userInt.UserID, // The UUID will be generated automatically by PostgreSQL
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

func UserIntUpdate(c *gin.Context) {
	username := c.Param("username")

	// Bind the request body to a temporary struct for updating specific fields
	var updateData struct {
		Username    string `json:"username"`
		Password    string `json:"password"` // Plain text password if updating
		Name        string `json:"name"`
		Phone       string `json:"phone"`
		NIK         string `json:"nik"`
		Role        bool   `json:"role"`
		BankAcc     string `json:"bank_acc"`
		BankAccName string `json:"bank_acc_name"`
		Status      bool   `json:"status"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	// Find the user by username
	var user models.UserIntModel
	result := initializers.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	// Check if password is provided and hash it if necessary
	if updateData.Password != "" {
		hashedPassword, err := auth.HashPassword(updateData.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash password"})
			return
		}
		updateData.Password = hashedPassword
	}

	// Update user fields (only non-empty fields will be updated)
	initializers.DB.Model(&user).Updates(updateData)

	c.JSON(200, gin.H{"message": "User updated successfully", "user": user})
}

func UserIntDelete(c *gin.Context) {
	username := c.Param("username")

	var user models.UserIntModel
	result := initializers.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	initializers.DB.Delete(&user)

	c.JSON(200, gin.H{"message": "User deleted successfully"})
}
