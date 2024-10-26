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

	r.Run()
}
