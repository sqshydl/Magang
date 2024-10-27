package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/squishydal/MAGANG/auth"
	"github.com/squishydal/MAGANG/controllers"
	"github.com/squishydal/MAGANG/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

// AuthMiddleware checks for valid JWT token
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix if present
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}

		claims, err := auth.ValidateToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set user information in context
		c.Set("userID", claims.UserID)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func main() {
	r := gin.Default()

	// Public routes (no authentication required)
	public := r.Group("/")
	{
		// Auth endpoints
		public.POST("/register", controllers.UserCreate) // Using existing UserCreate as register
		public.POST("/login", controllers.Login)         // New login endpoint
	}

	// Protected routes (require authentication)
	protected := r.Group("/")
	protected.Use(AuthMiddleware())
	{
		// User Intermediaris (UserInt) CRUD endpoints
		protected.POST("/userint", controllers.UserIntCreate)
		protected.GET("/userint", controllers.UserIntIndex)
		protected.GET("/userint/:username", controllers.UserIntGet)
		protected.DELETE("/userint/:username", controllers.UserIntDelete)
		protected.PUT("/userint/:username", controllers.UserIntUpdate)

		// User Central Bank (UserCentBank) CRUD endpoints
		protected.POST("/usercentbank", controllers.UserCentBankCreate)
		protected.GET("/usercentbank", controllers.UserCentBankIndex)
		protected.GET("/usercentbank/:username", controllers.UserCentBankGet)
		protected.DELETE("/usercentbank/:username", controllers.UserCentBankDelete)
		protected.PUT("/usercentbank/:username", controllers.UserCentBankUpdate)

		// Notification CRUD endpoints
		protected.POST("/notification", controllers.NotificationCreate)
		protected.GET("/notification", controllers.NotificationIndex)
		protected.GET("/notification/:notification_id", controllers.NotificationGet)
		protected.DELETE("/notification/:notification_id", controllers.NotificationDelete)
		protected.PUT("/notification/:notification_id", controllers.NotificationUpdate)

		// Transaction CRUD endpoints
		protected.POST("/transaction", controllers.TransactionCreate)
		protected.GET("/transaction", controllers.TransactionIndex)
		protected.GET("/transaction/:transaction_id", controllers.TransactionGet)
		protected.DELETE("/transaction/:transaction_id", controllers.TransactionDelete)
		protected.PUT("/transaction/:transaction_id", controllers.TransactionUpdate)

		// Validator CRUD endpoints
		protected.POST("/validator", controllers.ValidatorCreate)
		protected.GET("/validator", controllers.ValidatorIndex)
		protected.GET("/validator/:validators_id", controllers.ValidatorGet)
		protected.DELETE("/validator/:validators_id", controllers.ValidatorDelete)
		protected.PUT("/validator/:validators_id", controllers.ValidatorUpdate)

		// Issuing intermediaries CRUD endpoints
		protected.POST("/issuingintermediaries", controllers.IssuingIntermediaryCreate)
		protected.GET("/issuingintermediaries", controllers.IssuingIntermediaryIndex)
		protected.GET("/issuingintermediaries/:issuing_intermediaries_id", controllers.IssuingIntermediaryGet)
		protected.DELETE("/issuingintermediaries/:issuing_intermediaries_id", controllers.IssuingIntermediaryDelete)
		protected.PUT("/issuingintermediaries/:issuing_intermediaries_id", controllers.IssuingIntermediaryUpdate)

		// Redeem CRUD endpoints
		protected.POST("/redeem", controllers.RedeemCreate)
		protected.GET("/redeem", controllers.RedeemIndex)
		protected.GET("/redeem/:redeem_id", controllers.RedeemGet)
		protected.DELETE("/redeem/:redeem_id", controllers.RedeemDelete)
		protected.PUT("/redeem/:redeem_id", controllers.RedeemUpdate)

		// User CRUD endpoints
		protected.GET("/user", controllers.UserIndex)
		protected.GET("/user/:user_id", controllers.UserGet)
		protected.DELETE("/user/:user_id", controllers.UserDelete)
		protected.PUT("/user/:user_id", controllers.UserUpdate)
	}

	r.Run()
}
