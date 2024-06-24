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

func OKWithMsg(c *gin.Context, msg string) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  msg,
	})
}

func OKWithData(c *gin.Context, data any) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func BadRequest(c *gin.Context, msg string) {
	c.JSON(400, Response{
		Code: 400,
		Msg:  msg,
	})
}

func Unauthorized(c *gin.Context, msg string) {
	c.JSON(401, Response{
		Code: 401,
		Msg:  msg,
	})
}

func Forbidden(c *gin.Context, msg string) {
	c.JSON(403, Response{
		Code: 403,
		Msg:  msg,
	})
}

func NotFound(c *gin.Context, msg string) {
	c.JSON(404, Response{
		Code: 404,
		Msg:  msg,
	})
}

func InternalServerError(c *gin.Context, msg string) {
	c.JSON(500, Response{
		Code: 500,
		Msg:  msg,
	})
}

func ServiceUnavailable(c *gin.Context, msg string) {
	c.JSON(503, Response{
		Code: 503,
		Msg:  msg,
	})
}

func TooManyRequests(c *gin.Context, msg string) {
	c.JSON(429, Response{
		Code: 429,
		Msg:  msg,
	})
}

func Conflict(c *gin.Context, msg string) {
	c.JSON(409, Response{
		Code: 409,
		Msg:  msg,
	})
}

func Gone(c *gin.Context, msg string) {
	c.JSON(410, Response{
		Code: 410,
		Msg:  msg,
	})
}

func PreconditionFailed(c *gin.Context, msg string) {
	c.JSON(412, Response{
		Code: 412,
		Msg:  msg,
	})
}
