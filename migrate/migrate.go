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
}
