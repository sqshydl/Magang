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

	r.POST("/UserInt", controllers.UserIntCreate)
	r.GET("/UserInt", controllers.UserIntIndex)
	r.GET("/UserInt/:username", controllers.UserIntGet)

	r.Run() // listen and serve on 0.0.0.0:8080
}
