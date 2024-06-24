package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type RouteDefinition struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
}

var RouteDefinitions []RouteDefinition

func RegisterRoutes(r *gin.Engine) {
	fmt.Println("Registering routes", len(RouteDefinitions))
	for _, route := range RouteDefinitions {
		fmt.Println("Path", route.Path)
		switch route.Method {
		case "GET":
			r.GET(route.Path, route.Handler)
		case "POST":
			r.POST(route.Path, route.Handler)
			// Add other methods as needed
		}
	}
}
