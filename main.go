package main

import (
	"crud-golang/config"
	"crud-golang/models"
	"crud-golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	config.ConnectDB()

	config.DB.AutoMigrate(&models.User{})

	routes.UserRoutes(r)

	r.Run(":3000")
}
