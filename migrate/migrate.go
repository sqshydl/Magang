package main

import (
	"github.com/squishydal/MAGANG/initializers"
	"github.com/squishydal/MAGANG/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.UserIntModel{})
	initializers.DB.AutoMigrate(&models.UserCentBankModel{})
	initializers.DB.AutoMigrate(&models.Notification{})
	initializers.DB.AutoMigrate(&models.Transaction{})
	initializers.DB.AutoMigrate(&models.Validator{})
	initializers.DB.AutoMigrate(&models.IssuingIntermediaries{})
	initializers.DB.AutoMigrate(&models.Redeem{})
	initializers.DB.AutoMigrate(&models.User{})
}
