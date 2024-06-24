package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	RouteDefinitions = append(RouteDefinitions, RouteDefinition{
		Method: "GET",
		Path:   "/list/news",
		Handler: func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"code": 200,
				"msg":  "success",
				"data": "List News",
			})
		},
	})
}
