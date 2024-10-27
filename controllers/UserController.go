package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/squishydal/MAGANG/auth"
	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/models"
	"gorm.io/gorm"
)

// Create User
func UserCreate(c *gin.Context) {
	var userInput struct {
		Username         string `json:"username"`
		Password         string `json:"password"`
		Name             string `json:"name"`
		Phone            int    `json:"phone"`
		NIK              int    `json:"nik"`
		Role             string `json:"role"`
		BankAccount      int    `json:"bank_account"`
		BankAccountName  string `json:"bank_account_name"`
		IsIntermediaries int    `json:"is_intermediaries"`
	}

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	hashedPassword, err := auth.HashPassword(userInput.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to hash password"})
		return
	}

	user := models.User{
		UserID:           uuid.New(),
		Username:         userInput.Username,
		Password:         hashedPassword, // Assumes password is hashed before saving
		Name:             userInput.Name,
		Phone:            userInput.Phone,
		NIK:              userInput.NIK,
		Role:             userInput.Role,
		BankAccount:      userInput.BankAccount,
		BankAccountName:  userInput.BankAccountName,
		IsIntermediaries: userInput.IsIntermediaries,
		CreatedDate:      time.Now(),
		LogActivityIP:    c.ClientIP(),
		LogActivityTime:  time.Now(),
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

// Retrieve All Users
func UserIndex(c *gin.Context) {
	var users []models.User
	result := initializers.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// Retrieve Single User by ID
func UserGet(c *gin.Context) {
	id := c.Param("user_id")

	var user models.User
	result := initializers.DB.Where("user_id = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// Update User
func UserUpdate(c *gin.Context) {
	id := c.Param("user_id")

	var updateData struct {
		Username         string `json:"username"`
		Password         string `json:"password"`
		Name             string `json:"name"`
		Phone            int    `json:"phone"`
		NIK              int    `json:"nik"`
		Role             string `json:"role"`
		BankAccount      int    `json:"bank_account"`
		BankAccountName  string `json:"bank_account_name"`
		IsIntermediaries int    `json:"is_intermediaries"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var user models.User
	result := initializers.DB.First(&user, "user_id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	if updateData.Password != "" {
		hashedPassword, err := auth.HashPassword(updateData.Password)
		if err != nil {
			c.JSON(500, gin.H{"error": "Failed to hash password"})
			return
		}
		updateData.Password = hashedPassword
	}

	initializers.DB.Model(&user).Updates(updateData)

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

// Delete User
func UserDelete(c *gin.Context) {
	id := c.Param("user_id")

	var user models.User
	result := initializers.DB.First(&user, "user_id = ?", id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	initializers.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
