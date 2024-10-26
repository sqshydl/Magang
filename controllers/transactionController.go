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

// TransactionCreate creates a new transaction
func TransactionCreate(c *gin.Context) {
	// Define structure for receiving JSON input
	var transactionInput struct {
		UserID          uuid.UUID `json:"user_id"`
		TransactionDate time.Time `json:"transaction_date"`
		Price           float64   `json:"price"`
		TypePaymentID   int       `json:"type_payment_id"`
		PaymentStatus   string    `json:"payment_status"`
	}

	// Parse JSON input
	if err := c.BindJSON(&transactionInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Map input data to Transaction model
	transaction := models.Transaction{
		UserID:          transactionInput.UserID,
		TransactionDate: transactionInput.TransactionDate,
		Price:           transactionInput.Price,
		TypePaymentID:   transactionInput.TypePaymentID,
		PaymentStatus:   transactionInput.PaymentStatus,
		CreatedDate:     time.Now(),
	}

	// Insert new transaction into the database
	result := initializers.DB.Create(&transaction)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Respond with created transaction data
	c.JSON(http.StatusOK, gin.H{
		"message":     "Transaction created successfully",
		"transaction": transaction,
	})
}

// TransactionIndex retrieves all transactions
func TransactionIndex(c *gin.Context) {
	var allTransactions []models.Transaction
	result := initializers.DB.Find(&allTransactions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transactions": allTransactions,
	})
}

// TransactionGet retrieves a single transaction by ID
func TransactionGet(c *gin.Context) {
	id := c.Param("transaction_id")

	var transaction models.Transaction
	result := initializers.DB.First(&transaction, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction": transaction,
	})
}

// TransactionUpdate updates a transaction's details
func TransactionUpdate(c *gin.Context) {
	// Get the transaction ID from the URL parameter
	id := c.Param("transaction_id")

	// Define a temporary struct for fields that can be updated
	var updateData struct {
		UserID          int       `json:"user_id"`
		TransactionDate time.Time `json:"transaction_date"`
		Price           float64   `json:"price"`
		TypePaymentID   int       `json:"type_payment_id"`
		PaymentStatus   string    `json:"payment_status"`
	}

	// Bind JSON input to updateData struct
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Find the transaction by ID
	var transaction models.Transaction
	result := initializers.DB.First(&transaction, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Update the fields in the transaction model
	initializers.DB.Model(&transaction).Updates(updateData)

	// Respond with success and updated transaction
	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully", "transaction": transaction})
}

// TransactionDelete deletes a transaction by ID
func TransactionDelete(c *gin.Context) {
	// Get the transaction ID from the URL parameter
	id := c.Param("transaction_id")

	// Find the transaction by ID
	var transaction models.Transaction
	result := initializers.DB.First(&transaction, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// Delete the transaction
	initializers.DB.Delete(&transaction)

	// Respond with success message
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}
