package controller

import (
	"goapi/dataType"

	"github.com/gin-gonic/gin"
)

func init() {
	RouteDefinitions = append(RouteDefinitions, RouteDefinition{
		Method: "GET",
		Path:   "/list/news",
		Handler: func(c *gin.Context) {
			c.JSON(200, dataType.Response{
				Code: 200,
				Msg:  "success",
				Data: "List News",
			})
		},
	})
}
