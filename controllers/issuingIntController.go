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

func IssuingIntermediaryCreate(c *gin.Context) {
	var input struct {
		UserIntermediariesID      uuid.UUID `json:"user_intermediaries_id"`
		IssuingIntermediariesName string    `json:"issuing_intermediaries_name"`
		Point                     string    `json:"point"`
		Saldo                     string    `json:"saldo"`
		Redeem                    int       `json:"redeem"`
		LogActivityIP             string    `json:"log_activity_ip"`
	}

	// Parse JSON input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create a new issuing intermediary entry
	issuingIntermediary := models.IssuingIntermediaries{
		UserIntermediariesID:      input.UserIntermediariesID,
		IssuingIntermediariesName: input.IssuingIntermediariesName,
		Point:                     input.Point,
		Saldo:                     input.Saldo,
		Redeem:                    input.Redeem,
		CreatedDate:               time.Now(),
		LogActivityIP:             input.LogActivityIP,
		LogActivityTimestamp:      time.Now(),
	}

	if result := initializers.DB.Create(&issuingIntermediary); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":              "Issuing intermediary created successfully",
		"issuing_intermediary": issuingIntermediary,
	})
}
func IssuingIntermediaryIndex(c *gin.Context) {
	var issuingIntermediaries []models.IssuingIntermediaries

	if result := initializers.DB.Find(&issuingIntermediaries); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"issuing_intermediaries": issuingIntermediaries})
}
func IssuingIntermediaryGet(c *gin.Context) {
	id := c.Param("issuing_intermediaries_id")

	var issuingIntermediary models.IssuingIntermediaries
	if result := initializers.DB.First(&issuingIntermediary, "issuing_intermediaries_id = ?", id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Issuing intermediary not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"issuing_intermediary": issuingIntermediary})
}
func IssuingIntermediaryUpdate(c *gin.Context) {
	id := c.Param("issuing_intermediaries_id")

	var updateData struct {
		IssuingIntermediariesName string `json:"issuing_intermediaries_name"`
		Point                     string `json:"point"`
		Saldo                     string `json:"saldo"`
		Redeem                    int    `json:"redeem"`
		LogActivityIP             string `json:"log_activity_ip"`
	}

	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var issuingIntermediary models.IssuingIntermediaries
	if result := initializers.DB.First(&issuingIntermediary, "issuing_intermediaries_id = ?", id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Issuing intermediary not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	initializers.DB.Model(&issuingIntermediary).Updates(updateData)

	c.JSON(http.StatusOK, gin.H{"message": "Issuing intermediary updated successfully", "issuing_intermediary": issuingIntermediary})
}
func IssuingIntermediaryDelete(c *gin.Context) {
	id := c.Param("issuing_intermediaries_id")

	var issuingIntermediary models.IssuingIntermediaries
	if result := initializers.DB.First(&issuingIntermediary, "issuing_intermediaries_id = ?", id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Issuing intermediary not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	initializers.DB.Delete(&issuingIntermediary)

	c.JSON(http.StatusOK, gin.H{"message": "Issuing intermediary deleted successfully"})
}
