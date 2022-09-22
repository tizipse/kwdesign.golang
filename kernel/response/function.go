package response

import (
	"github.com/gin-gonic/gin"
	"kwd/kernel/validator"
	"net/http"
)

func Unauthorized(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response[any]{
		Code:    40100,
		Message: "Unauthorized",
	})
}

func Forbidden(ctx *gin.Context) {
	ctx.JSON(http.StatusForbidden, Response[any]{
		Code:    40400,
		Message: "Forbidden",
	})
}

func NotFound(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, Response[any]{
		Code:    40400,
		Message: message,
	})
}

func FailByRequest(ctx *gin.Context, err error) {

	ctx.JSON(http.StatusOK, Response[any]{
		Code:    40000,
		Message: validator.Translate(err),
	})
}

func FailByLogin(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, Response[any]{
		Code:    40100,
		Message: "登陆失败",
	})
}

func Success[T any](ctx *gin.Context, data ...T) {

	response := Response[T]{
		Code:    20000,
		Message: "Success",
	}

	if len(data) > 0 {
		response.Data = data[0]
	}

	ctx.JSON(http.StatusOK, response)
}

func Fail(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, Response[any]{
		Code:    60000,
		Message: message,
	})
}
