package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/models"
	"gorm.io/gorm"
)

func ValidatorCreate(c *gin.Context) {
	var validatorInput struct {
		UserIntermediariesID uuid.UUID `json:"user_intermediaries_id"`
	}

	// Parse JSON input
	if err := c.BindJSON(&validatorInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Map input data to Validator model
	validator := models.Validator{
		ValidatorsID:         uuid.NewString(),
		UserIntermediariesID: validatorInput.UserIntermediariesID,
		CreatedDate:          time.Now(),
	}

	// Insert new validator into the database
	if result := initializers.DB.Create(&validator); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond with created validator data
	c.JSON(http.StatusOK, gin.H{
		"message":   "Validator created successfully",
		"validator": validator,
	})
}

func ValidatorIndex(c *gin.Context) {
	var validators []models.Validator

	// Retrieve all validators from the database
	if result := initializers.DB.Preload("UserIntermediary").Find(&validators); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"validators": validators})
}

func ValidatorGet(c *gin.Context) {
	id := c.Param("validators_id")

	var validator models.Validator
	// Find the validator by ID
	if result := initializers.DB.Preload("UserIntermediary").First(&validator, "validators_id = ?", id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Validator not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"validator": validator})
}

func ValidatorUpdate(c *gin.Context) {
	id := c.Param("validators_id")

	var updateData struct {
		UserIntermediariesID uuid.UUID `json:"user_intermediaries_id"`
	}

	// Bind JSON input to updateData struct
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var validator models.Validator
	// Find the validator by ID
	if result := initializers.DB.First(&validator, "validators_id = ?", id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Validator not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Update the validator with the new UserIntermediariesID
	initializers.DB.Model(&validator).Update("UserIntermediariesID", updateData.UserIntermediariesID)

	c.JSON(http.StatusOK, gin.H{"message": "Validator updated successfully", "validator": validator})
}

func ValidatorDelete(c *gin.Context) {
	id := c.Param("validators_id")

	var validator models.Validator
	// Find the validator by ID
	if result := initializers.DB.First(&validator, "validators_id = ?", id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Validator not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Delete the validator
	initializers.DB.Delete(&validator)

	c.JSON(http.StatusOK, gin.H{"message": "Validator deleted successfully"})
}
