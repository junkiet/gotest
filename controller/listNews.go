package controller

import (
	"goapi/res"
	"time"

	"github.com/gin-gonic/gin"
)

type News struct {
	Id          int
	Title       string
	Description string
	CreatedAt   time.Time
}

func init() {
	RouteDefinitions = append(RouteDefinitions, RouteDefinition{
		Method: "GET",
		Path:   "/list/news",
		Handler: func(c *gin.Context) {

			jsonData := []News{
				{
					Id:          1,
					Title:       "News 1",
					Description: "This is news 1",
					CreatedAt:   time.Now(),
				},
				{
					Id:          2,
					Title:       "News 2",
					Description: "This is news 2",
					CreatedAt:   time.Now(),
				},
			}

			res.OKWithData(c, jsonData)
		},
	})
}
