package res

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func OK(c *gin.Context, data any) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  "success",
	})
}

func OKWithData(c *gin.Context, data any) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func BadRequest(c *gin.Context, data any) {
	c.JSON(400, Response{
		Code: 400,
		Msg:  "failed",
		Data: data,
	})
}

func Unauthorized(c *gin.Context, data any) {
	c.JSON(401, Response{
		Code: 401,
		Msg:  "failed",
		Data: data,
	})
}

func Forbidden(c *gin.Context, data any) {
	c.JSON(403, Response{
		Code: 403,
		Msg:  "failed",
		Data: data,
	})
}

func NotFound(c *gin.Context, data any) {
	c.JSON(404, Response{
		Code: 404,
		Msg:  "failed",
		Data: data,
	})
}

func InternalServerError(c *gin.Context, data any) {
	c.JSON(500, Response{
		Code: 500,
		Msg:  "failed",
		Data: data,
	})
}

func ServiceUnavailable(c *gin.Context, data any) {
	c.JSON(503, Response{
		Code: 503,
		Msg:  "failed",
		Data: data,
	})
}

func TooManyRequests(c *gin.Context) {
	c.JSON(429, Response{
		Code: 429,
		Msg:  "failed",
	})
}
