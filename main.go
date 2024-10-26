package main

import (
	"github.com/gin-gonic/gin"
	"github.com/squishydal/MAGANG/controllers"
	"github.com/squishydal/MAGANG/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	//User Intermediaris (UserInt) CRUD endpoints
	r.POST("/userint", controllers.UserIntCreate)
	r.GET("/userint", controllers.UserIntIndex)
	r.GET("/userint/:username", controllers.UserIntGet)
	r.DELETE("/userint/:username", controllers.UserIntDelete)
	r.PUT("/userint/:username", controllers.UserIntUpdate)

	//User Central Bank (UserCentBank) CRUD endpoints
	r.POST("/usercentbank", controllers.UserCentBankCreate)
	r.GET("/usercentbank", controllers.UserCentBankIndex)
	r.GET("/usercentbank/:username", controllers.UserCentBankGet)
	r.DELETE("/usercentbank/:username", controllers.UserCentBankDelete)
	r.PUT("/usercentbank/:username", controllers.UserCentBankUpdate)

	//notification CRUD endpoints
	r.POST("/notification", controllers.NotificationCreate)
	r.GET("/notification", controllers.NotificationIndex)
	r.GET("/notification/:notification_id", controllers.NotificationGet)
	r.DELETE("/notification/:notification_id", controllers.NotificationDelete)
	r.PUT("/notification/:notification_id", controllers.NotificationUpdate)

	//transaction CRUD endpoints
	r.POST("/transaction", controllers.TransactionCreate)
	r.GET("/transaction", controllers.TransactionIndex)
	r.GET("/transaction/:transaction_id", controllers.TransactionGet)
	r.DELETE("/transaction/:transaction_id", controllers.TransactionDelete)
	r.PUT("/transaction/:transaction_id", controllers.TransactionUpdate)

	//validator CRUD endpoints
	r.POST("/validator", controllers.ValidatorCreate)
	r.GET("/validator", controllers.ValidatorIndex)
	r.GET("/validator/:validators_id", controllers.ValidatorGet)
	r.DELETE("/validator/:validators_id", controllers.ValidatorDelete)
	r.PUT("/validator/:validators_id", controllers.ValidatorUpdate)

	//issuing intermediaries CRUD endpoints
	r.POST("/issuingintermediaries", controllers.IssuingIntermediaryCreate)
	r.GET("/issuingintermediaries", controllers.IssuingIntermediaryIndex)
	r.GET("/issuingintermediaries/:issuing_intermediaries_id", controllers.IssuingIntermediaryGet)
	r.DELETE("/issuingintermediaries/:issuing_intermediaries_id", controllers.IssuingIntermediaryDelete)
	r.PUT("/issuingintermediaries/:issuing_intermediaries_id", controllers.IssuingIntermediaryUpdate)

	//redeem CRUD endpoints
	r.POST("/redeem", controllers.RedeemCreate)
	r.GET("/redeem", controllers.RedeemIndex)
	r.GET("/redeem/:redeem_id", controllers.RedeemGet)
	r.DELETE("/redeem/:redeem_id", controllers.RedeemDelete)
	r.PUT("/redeem/:redeem_id", controllers.RedeemUpdate)

	r.Run()
}
