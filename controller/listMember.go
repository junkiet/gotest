package controller

import (
	"goapi/model"
	"goapi/mysql"
	"goapi/res"

	"github.com/gin-gonic/gin"
)

func init() {
	RouteDefinitions = append(RouteDefinitions, RouteDefinition{
		Method: "GET",
		Path:   "/list/member",
		Handler: func(c *gin.Context) {
			var members []model.Member

			mysql.DB.Where("id = ?", 1200).Find(&members)
			if len(members) == 0 {
				res.NotFound(c, "notfound")
				return
			}

			res.OKWithData(c, members)
		},
	})
}
