package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/models"
	"gorm.io/gorm"
)

func RedeemCreate(c *gin.Context) {
	var input struct {
		UserID        int    `json:"user_id"`
		NominalRedeem int    `json:"nominal_redeem"`
		Status        string `json:"status"`
	}

	// Parse JSON input
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Create a new redeem entry
	redeem := models.Redeem{
		UserID:        input.UserID,
		NominalRedeem: input.NominalRedeem,
		Status:        input.Status,
		CreatedDate:   time.Now(),
	}

	if result := initializers.DB.Create(&redeem); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Redeem request created successfully",
		"redeem":  redeem,
	})
}
func RedeemIndex(c *gin.Context) {
	var redeems []models.Redeem

	if result := initializers.DB.Find(&redeems); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"redeems": redeems})
}
func RedeemGet(c *gin.Context) {
	id := c.Param("redeem_id")

	var redeem models.Redeem
	if result := initializers.DB.First(&redeem, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Redeem request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"redeem": redeem})
}
func RedeemUpdate(c *gin.Context) {
	id := c.Param("redeem_id")

	var updateData struct {
		NominalRedeem int    `json:"nominal_redeem"`
		Status        string `json:"status"`
	}

	// Parse JSON input
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	var redeem models.Redeem
	if result := initializers.DB.First(&redeem, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Redeem request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	initializers.DB.Model(&redeem).Updates(updateData)

	c.JSON(http.StatusOK, gin.H{"message": "Redeem request updated successfully", "redeem": redeem})
}
func RedeemDelete(c *gin.Context) {
	id := c.Param("redeem_id")

	var redeem models.Redeem
	if result := initializers.DB.First(&redeem, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Redeem request not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	initializers.DB.Delete(&redeem)

	c.JSON(http.StatusOK, gin.H{"message": "Redeem request deleted successfully"})
}
