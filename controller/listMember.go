package controller

import (
	"goapi/res"

	"github.com/gin-gonic/gin"
)

func init() {
	RouteDefinitions = append(RouteDefinitions, RouteDefinition{
		Method: "GET",
		Path:   "/list/member",
		Handler: func(c *gin.Context) {
			res.OKWithData(c, "List Member")
		},
	})
}
