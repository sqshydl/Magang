package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid" // Generates UUID for UserID	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/models"
)

// Register handles user registration by receiving user data, hashing the password, and storing it in the database.
func Register(c *gin.Context) {
	// Struct to capture input from JSON request
	var registerInput struct {
		Username       string `json:"username" binding:"required"`
		Password       string `json:"password" binding:"required"`
		Name           string `json:"name" binding:"required"`
		Phone          int    `json:"phone"`
		NIK            int    `json:"nik"`
		Role           string `json:"role" binding:"required"`
		BankAccount    int    `json:"bank_account"`
		BankAccName    string `json:"bank_account_name"`
		IsIntermediary int    `json:"is_intermediaries"` // 1 for yes, 0 for no
	}

	// Bind JSON to registerInput struct; return error if any required field is missing
	if err := c.ShouldBindJSON(&registerInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: missing or incorrect data fields"})
		return
	}

	// Hash the password
	hashedPassword, err := HashPassword(registerInput.Password) // HashPassword is in auth package
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Initialize new user based on received input
	user := models.UserIntModel{
		UserID:       uuid.New(), // Generate a new UUID for the user
		Username:     registerInput.Username,
		Password:     hashedPassword, // Store hashed password
		Name:         registerInput.Name,
		Phone:        registerInput.Phone,
		NIK:          registerInput.NIK,
		Role:         registerInput.Role == "validator", // true if role is validator
		BankAcc:      registerInput.BankAccount,
		BankAccName:  registerInput.BankAccName,
		Status:       true,         // default status as active
		CreatedAt:    time.Now(),   // sets created date
		LastActiveIP: c.ClientIP(), // captures current client IP
	}

	// Insert user into the database
	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond with success, excluding password for security
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user_id": user.UserID,
		"user": gin.H{
			"username":          user.Username,
			"name":              user.Name,
			"phone":             user.Phone,
			"role":              registerInput.Role,
			"bank_account":      user.BankAcc,
			"bank_account_name": user.BankAccName,
			"created_date":      user.CreatedAt,
		},
	})
}
