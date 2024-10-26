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
	r.POST("/UserInt", controllers.UserIntCreate)
	r.GET("/UserInt", controllers.UserIntIndex)
	r.GET("/UserInt/:username", controllers.UserIntGet)
	r.DELETE("/UserInt/:username", controllers.UserIntDelete)
	r.PUT("/UserInt/:username", controllers.UserIntUpdate)

	//User Central Bank (UserCentBank) CRUD endpoints
	r.POST("/UserCentBank", controllers.UserCentBankCreate)
	r.GET("/UserCentBank", controllers.UserCentBankIndex)
	r.GET("/UserCentBank/:username", controllers.UserCentBankGet)
	r.DELETE("/UserCentBank/:username", controllers.UserCentBankDelete)
	r.PUT("/UserCentBank/:username", controllers.UserCentBankUpdate)

	r.Run()
}
