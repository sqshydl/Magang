package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/squishydal/MAGANG/auth"
	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/models"
	"gorm.io/gorm"
)

func UserCentBankCreate(c *gin.Context) {
	var createUserCentBank struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
	}

	if err := c.BindJSON(&createUserCentBank); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	// Check if the username already exists
	var existingUser models.UserCentBankModel
	if err := initializers.DB.Where("username = ?", createUserCentBank.Username).First(&existingUser).Error; err == nil {
		c.JSON(409, gin.H{"error": "Username already exists"})
		return
	}

	// Hash the password before storing it
	hashedPassword, err := auth.HashPassword(createUserCentBank.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new user with the defined structure
	userCentBank := models.UserCentBankModel{
		Name:           createUserCentBank.Name,
		Username:       createUserCentBank.Username,
		Password:       hashedPassword,
		LastActivityIP: c.ClientIP(),
		LastActiveTime: time.Now(),
		CreatedDate:    time.Now(),
	}

	result := initializers.DB.Create(&userCentBank)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully",
		"user":    userCentBank,
	})
}

func UserCentBankIndex(c *gin.Context) {
	var getAllUserCentBank []models.UserCentBankModel
	initializers.DB.Find(&getAllUserCentBank)
	result := initializers.DB.Find(&getAllUserCentBank)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"post": getAllUserCentBank,
	})
}

func UserCentBankGet(c *gin.Context) {
	username := c.Param("username")

	var getUserCentBank models.UserCentBankModel
	result := initializers.DB.Where("username = ?", username).First(&getUserCentBank)

	if result.Error != nil {

		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{"error": "User not found"})
			return
		}

		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(200, gin.H{
		"user": getUserCentBank,
	})
}

func UserCentBankUpdate(c *gin.Context) {
	username := c.Param("username")

	// Bind the request body to a temporary struct for updating specific fields
	var updateData struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name"`
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

func UserCentBankDelete(c *gin.Context) {
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
