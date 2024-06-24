package main

import (
	"goapi/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Load routes from controllers
	controller.RegisterRoutes(r)

	// Start the server on port 8080
	r.Run(":8080")
}
