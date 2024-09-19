package main

import (
	"goapi/controller"
	"goapi/mysql"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Initialize the database
	mysql.Init()

	// Load routes from controllers
	controller.RegisterRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
