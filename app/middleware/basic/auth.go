package basic

import (
	"github.com/gin-gonic/gin"
	"kwd/app/constant"
	"kwd/app/service/helper"
	"kwd/kernel/authorize"
	"kwd/kernel/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		if !authorize.Check(ctx) {
			ctx.Abort()
			response.Unauthorized(ctx)
			return
		}

		claims := authorize.Jwt(ctx)

		if !claims.VerifyIssuer(constant.ContextAdmin, true) {
			ctx.Abort()
			response.Unauthorized(ctx)
			return
		}

		if !helper.CheckJwt(ctx, constant.ContextAdmin, *claims) {
			ctx.Abort()
			response.Unauthorized(ctx)
			return
		}

		ctx.Next()
	}
}
